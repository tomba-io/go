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
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/tomba-io/go/tomba/models"
)

// Bulk jobs with the canonical parameters. See https://docs.tomba.io/bulks.

func bulkJobPath(t models.BulkType, id int64, suffix string) string {
	return fmt.Sprintf(BULK_PATH+"/%d%s", string(t), id, suffix)
}

// CreateBulkJob creates a bulk job. The input is req.Data (rows of cells),
// req.List (one value per line) or req.FilePath (a CSV upload). With
// req.Launch the job starts right away; if the launch is refused (quota,
// running jobs) the job still exists and the result's Message says why.
func (conf *Tomba) CreateBulkJob(t models.BulkType, req *models.BulkJobRequest) (*models.BulkCreateResult, error) {
	var out models.BulkCreateResult
	if req.FilePath == "" {
		err := conf.doJSON(http.MethodPost, getBulkPath(t), nil, req, &out)
		return &out, err
	}
	body, contentType, err := multipartJobBody(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest(http.MethodPost, conf.baseURL()+getBulkPath(t), body)
	if err != nil {
		return nil, err
	}
	conf.setAuthHeaders(httpReq)
	httpReq.Header.Set("Content-Type", contentType)
	err = conf.send(httpReq, &out)
	return &out, err
}

// multipartJobBody encodes req as a form: scalar fields as values, slices
// repeated, Data as JSON, and the CSV file as a text/csv part.
func multipartJobBody(req *models.BulkJobRequest) (io.Reader, string, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fields := map[string]interface{}{}
	raw, _ := json.Marshal(req)
	_ = json.Unmarshal(raw, &fields)
	for k, v := range fields {
		switch x := v.(type) {
		case []interface{}:
			if k == "data" {
				b, _ := json.Marshal(x)
				_ = w.WriteField(k, string(b))
				continue
			}
			for _, e := range x {
				_ = w.WriteField(k, fmt.Sprint(e))
			}
		case float64:
			_ = w.WriteField(k, strconv.FormatFloat(x, 'f', -1, 64))
		default:
			_ = w.WriteField(k, fmt.Sprint(x))
		}
	}
	f, err := os.Open(req.FilePath)
	if err != nil {
		return nil, "", err
	}
	defer func() { _ = f.Close() }()
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filepath.Base(req.FilePath)))
	h.Set("Content-Type", "text/csv")
	part, err := w.CreatePart(h)
	if err != nil {
		return nil, "", err
	}
	if _, err := io.Copy(part, f); err != nil {
		return nil, "", err
	}
	if err := w.Close(); err != nil {
		return nil, "", err
	}
	return &buf, w.FormDataContentType(), nil
}

// GetBulkJob returns one job, with its preview rows.
func (conf *Tomba) GetBulkJob(t models.BulkType, id int64) (*models.BulkItem, error) {
	var out struct {
		Data *models.BulkItem `json:"data"`
	}
	if err := conf.doJSON(http.MethodGet, bulkJobPath(t, id, ""), nil, nil, &out); err != nil {
		return nil, err
	}
	if out.Data == nil {
		return nil, &APIError{HTTPStatus: http.StatusNotFound, Type: "unknown_record", Message: "Record was not found."}
	}
	return out.Data, nil
}

// ListBulkJobs lists jobs of a type. status and q are optional filters.
func (conf *Tomba) ListBulkJobs(t models.BulkType, page, limit int, archived bool, status, q string) (*models.BulkListResponse, error) {
	query := url.Values{}
	if page > 0 {
		query.Set("page", strconv.Itoa(page))
	}
	if limit > 0 {
		query.Set("limit", strconv.Itoa(limit))
	}
	if archived {
		query.Set("archived", "true")
	}
	if status != "" {
		query.Set("status", status)
	}
	if q != "" {
		query.Set("q", q)
	}
	var out models.BulkListResponse
	err := conf.doJSON(http.MethodGet, getBulkPath(t), query, nil, &out)
	return &out, err
}

// RestoreBulk un-archives a job.
func (conf *Tomba) RestoreBulk(t models.BulkType, id int64) (*models.BulkSuccessResponse, error) {
	return conf.bulkAction(http.MethodPut, t, id, "/restore")
}

// CancelBulk stops a pending (or, when enabled on the server, running) job.
func (conf *Tomba) CancelBulk(t models.BulkType, id int64) (*models.BulkSuccessResponse, error) {
	return conf.bulkAction(http.MethodPost, t, id, "/cancel")
}

// RetryBulk runs a failed or cancelled job again from its input.
func (conf *Tomba) RetryBulk(t models.BulkType, id int64) (*models.BulkSuccessResponse, error) {
	return conf.bulkAction(http.MethodPost, t, id, "/retry")
}

func (conf *Tomba) bulkAction(method string, t models.BulkType, id int64, suffix string) (*models.BulkSuccessResponse, error) {
	var out models.BulkSuccessResponse
	err := conf.doJSON(method, bulkJobPath(t, id, suffix), nil, nil, &out)
	return &out, err
}

// GetBulkEstimate returns the credits a job costs.
func (conf *Tomba) GetBulkEstimate(t models.BulkType, id int64) (*models.BulkEstimate, error) {
	var out struct {
		Data models.BulkEstimate `json:"data"`
	}
	err := conf.doJSON(http.MethodGet, bulkJobPath(t, id, "/estimate"), nil, nil, &out)
	return &out.Data, err
}

// GetBulkHistory returns a job's events, oldest first.
func (conf *Tomba) GetBulkHistory(t models.BulkType, id int64, page, limit int) ([]models.BulkEvent, error) {
	query := url.Values{}
	if page > 0 {
		query.Set("page", strconv.Itoa(page))
	}
	if limit > 0 {
		query.Set("limit", strconv.Itoa(limit))
	}
	var out struct {
		Data []models.BulkEvent `json:"data"`
	}
	err := conf.doJSON(http.MethodGet, bulkJobPath(t, id, "/history"), query, nil, &out)
	return out.Data, err
}

// GetBulkJobProgress returns a job's progress with live metrics and ETA.
func (conf *Tomba) GetBulkJobProgress(t models.BulkType, id int64) (*models.BulkProgress, error) {
	var out models.BulkProgress
	err := conf.doJSON(http.MethodGet, bulkJobPath(t, id, "/progress"), nil, nil, &out)
	return &out, err
}

// GetBulkTypes returns every bulk type with its fields, options and limits.
func (conf *Tomba) GetBulkTypes() ([]models.BulkTypeInfo, error) {
	var out struct {
		Data []models.BulkTypeInfo `json:"data"`
	}
	err := conf.doJSON(http.MethodGet, "/bulk/types", nil, nil, &out)
	return out.Data, err
}

// GetBulkStats returns account-level bulk activity. typ, from and to
// (YYYY-MM-DD) are optional; the default period is the last 30 days.
func (conf *Tomba) GetBulkStats(typ, from, to string) (*models.BulkStats, error) {
	query := url.Values{}
	for k, v := range map[string]string{"type": typ, "from": from, "to": to} {
		if v != "" {
			query.Set(k, v)
		}
	}
	var out struct {
		Data models.BulkStats `json:"data"`
	}
	err := conf.doJSON(http.MethodGet, "/bulk/stats", query, nil, &out)
	return &out.Data, err
}

// GetBulkWebhookSecret returns the key completion webhooks are signed with.
func (conf *Tomba) GetBulkWebhookSecret() (*models.BulkWebhookSecret, error) {
	var out struct {
		Data models.BulkWebhookSecret `json:"data"`
	}
	err := conf.doJSON(http.MethodGet, "/bulk/webhook-secret", nil, nil, &out)
	return &out.Data, err
}

// RotateBulkWebhookSecret replaces the webhook signing key.
func (conf *Tomba) RotateBulkWebhookSecret() (*models.BulkWebhookSecret, error) {
	var out struct {
		Data models.BulkWebhookSecret `json:"data"`
	}
	err := conf.doJSON(http.MethodPost, "/bulk/webhook-secret/rotate", nil, nil, &out)
	return &out.Data, err
}

// DownloadBulkTo streams a result file ("full", "valid" or "not_found") to
// w. The first download of a job is billed.
func (conf *Tomba) DownloadBulkTo(t models.BulkType, id int64, kind string, w io.Writer) (int64, error) {
	u := conf.baseURL() + bulkJobPath(t, id, "/download")
	if kind != "" {
		u += "?file=" + url.QueryEscape(kind)
	}
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return 0, err
	}
	conf.setAuthHeaders(req)
	resp, err := (&http.Client{Timeout: 10 * time.Minute}).Do(req)
	if err != nil {
		return 0, err
	}
	defer func() { _ = resp.Body.Close() }()
	conf.LastRateLimit = ParseRateLimit(resp)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 64<<10))
		return 0, newAPIError(resp, body)
	}
	return io.Copy(w, resp.Body)
}

// send executes req and decodes a JSON answer into out.
func (conf *Tomba) send(req *http.Request, out interface{}) error {
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
	if out == nil || reflect.ValueOf(out).IsNil() {
		return nil
	}
	return json.Unmarshal(data, out)
}

// IsBulkTerminal reports whether a job status is final.
func IsBulkTerminal(status string) bool {
	switch strings.ToLower(status) {
	case models.BulkStatusCompleted, models.BulkStatusFailed, models.BulkStatusCancelled:
		return true
	}
	return false
}
