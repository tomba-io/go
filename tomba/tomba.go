package tomba

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/tomba-io/go/tomba/models"
)

// TombaCall makes an HTTP request to the Tomba API with support for file uploads.
func (conf *Tomba) TombaCall(path string, params Params, method *string, fileUpload *models.BulkFileUpload) ([]byte, error) {
	var req *http.Request
	apiUrl := DEFAULT_BASE_URL + path
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

		part, partErr := writer.CreateFormFile(fileUpload.FieldName, fileUpload.FilePath)
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

	// Set common headers
	req.Header.Set("X-Tomba-Key", conf.ApiKey)
	req.Header.Set("X-Tomba-Secret", conf.ApiSecret)
	req.Header.Set("User-Agent", "tomba-go-client/1.0")
	req.Header.Set("Content-Type", "application/json")
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

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("error: %s, status code: %d", resp.Status, resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
