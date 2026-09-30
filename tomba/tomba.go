package tomba

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/tomba-io/go/tomba/models"
)

// TombaCall makes an HTTP request to the Tomba API with support for file uploads.
func (conf *Tomba) TombaCall(path string, params Params, method *string, fileUpload *models.BulkFileUpload) ([]byte, error) {
	var req *http.Request
	apiUrl := conf.baseURL() + path
	methodStr := "GET"
	if method != nil {
		methodStr = *method
	}

	// Handle file upload requests
	if fileUpload != nil && fileUpload.FilePath != "" {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// Add form fields from params
		for key, value := range params {
			switch v := value.(type) {
			case string:
				_ = writer.WriteField(key, v)
			case int:
				_ = writer.WriteField(key, strconv.Itoa(v))
			case bool:
				_ = writer.WriteField(key, strconv.FormatBool(v))
			case float64:
				_ = writer.WriteField(key, strconv.FormatFloat(v, 'f', -1, 64))
			}
		}

		// Add file
		file, openErr := os.Open(fileUpload.FilePath)
		if openErr != nil {
			return nil, openErr
		}
		defer func() { _ = file.Close() }()

		// The API accepts CSV uploads (text/csv); CreateFormFile would label
		// the part application/octet-stream.
		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fileUpload.FieldName, filepath.Base(fileUpload.FilePath)))
		partHeader.Set("Content-Type", "text/csv")
		part, partErr := writer.CreatePart(partHeader)
		if partErr != nil {
			return nil, partErr
		}

		if _, copyErr := io.Copy(part, file); copyErr != nil {
			return nil, copyErr
		}

		if closeErr := writer.Close(); closeErr != nil {
			return nil, closeErr
		}

		var reqErr error
		req, reqErr = http.NewRequest(methodStr, apiUrl, body)
		if reqErr != nil {
			return nil, reqErr
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else if methodStr == "GET" {
		// Handle GET requests with query parameters
		var reqErr error
		req, reqErr = http.NewRequest(methodStr, apiUrl, nil)
		if reqErr != nil {
			return nil, reqErr
		}

		if len(params) > 0 {
			q := req.URL.Query()
			for key, value := range params {
				switch v := value.(type) {
				case string:
					if v != "" {
						q.Add(key, v)
					}
				case int:
					if v > 0 {
						q.Add(key, strconv.Itoa(v))
					}
				case bool:
					q.Add(key, strconv.FormatBool(v))
				}
			}
			req.URL.RawQuery = q.Encode()
			req.Header.Set("Content-Type", "application/json")
		}
	} else {
		jsonData, marshalErr := json.Marshal(params)
		if marshalErr != nil {
			return nil, marshalErr
		}
		var reqErr error
		req, reqErr = http.NewRequest(methodStr, apiUrl, bytes.NewBuffer(jsonData))
		if reqErr != nil {
			return nil, reqErr
		}
	}

	// Common headers. A multipart request keeps its own Content-Type (with the
	// boundary): overwriting it made every file upload unreadable.
	conf.setAuthHeaders(req)
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// Parse rate limit headers from response
	conf.LastRateLimit = ParseRateLimit(resp)

	if resp.Body != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	body, readErr := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, newAPIError(resp, body)
	}
	return body, readErr
}

func (conf *Tomba) setAuthHeaders(req *http.Request) {
	req.Header.Set("X-Tomba-Key", conf.ApiKey)
	req.Header.Set("X-Tomba-Secret", conf.ApiSecret)
	req.Header.Set("User-Agent", "tomba-go-client/"+Version)
}

// doJSON sends body as JSON (or no body) and decodes the answer into out.
func (conf *Tomba) doJSON(method, path string, query url.Values, body interface{}, out interface{}) error {
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(data)
	}
	u := conf.baseURL() + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	req, err := http.NewRequest(method, u, reader)
	if err != nil {
		return err
	}
	conf.setAuthHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := (&http.Client{Timeout: 120 * time.Second}).Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	conf.LastRateLimit = ParseRateLimit(resp)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return newAPIError(resp, data)
	}
	if out == nil {
		return nil
	}
	return json.Unmarshal(data, out)
}
