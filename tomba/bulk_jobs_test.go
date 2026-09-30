package tomba

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tomba-io/go/tomba/models"
)

func intp(i int) *int    { return &i }
func boolp(b bool) *bool { return &b }

func TestCreateBulkJobJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/bulk/finder" || r.Header.Get("X-Tomba-Key") != "ta_k" {
			t.Errorf("path=%s key=%s", r.URL.Path, r.Header.Get("X-Tomba-Key"))
		}
		var body map[string]interface{}
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["domain_field_index"].(float64) != 0 || body["verify_emails"] != true || len(body["data"].([]interface{})) != 2 {
			t.Errorf("body = %v", body)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"data":{"id":4521,"message":"Bulk created but not launched: Insufficient search quota"}}`))
	}))
	defer srv.Close()

	c := New("ta_k", "ts_s").WithBaseURL(srv.URL + "/v1/")
	res, err := c.CreateBulkJob(models.BulkTypeFinder, &models.BulkJobRequest{
		Name: "Leads", Data: [][]string{{"domain", "first", "last"}, {"stripe.com", "Jane", "Doe"}},
		DomainFieldIndex: intp(0), VerifyEmails: boolp(true), Launch: true,
	})
	if err != nil || *res.Data.ID != 4521 || !strings.Contains(res.Data.Message, "not launched") {
		t.Fatalf("res=%+v err=%v", res, err)
	}
}

// Regression: uploads used to be sent with Content-Type application/json and
// an application/octet-stream file part, which the API rejected.
func TestCreateBulkJobUploadIsMultipartCSV(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "leads.csv")
	_ = os.WriteFile(path, []byte("email\njane@stripe.com\n"), 0o644)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data; boundary=") {
			t.Fatalf("content-type = %q", r.Header.Get("Content-Type"))
		}
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			t.Fatal(err)
		}
		fh := r.MultipartForm.File["file"][0]
		if fh.Header.Get("Content-Type") != "text/csv" || fh.Filename != "leads.csv" {
			t.Errorf("file part: %v %s", fh.Header, fh.Filename)
		}
		if r.FormValue("email_field_index") != "0" || r.FormValue("name") != "Emails" {
			t.Errorf("form = %v", r.MultipartForm.Value)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"data":{"id":7}}`))
	}))
	defer srv.Close()
	c := New("k", "s").WithBaseURL(srv.URL)
	res, err := c.CreateBulkJob(models.BulkTypeVerifier, &models.BulkJobRequest{Name: "Emails", FilePath: path, EmailFieldIndex: intp(0)})
	if err != nil || *res.Data.ID != 7 {
		t.Fatalf("res=%+v err=%v", res, err)
	}
}

func TestAPIErrorCarriesTheEnvelope(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Retry-After", "60")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"errors":{"type":"rate_limit","message":"You already have 2 Email Finder bulks running.","code":429}}`))
	}))
	defer srv.Close()
	_, err := New("k", "s").WithBaseURL(srv.URL).RetryBulk(models.BulkTypeFinder, 1)
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.HTTPStatus != 429 || apiErr.Type != "rate_limit" || apiErr.RetryAfter != 60 {
		t.Fatalf("err = %#v", err)
	}
	if !strings.Contains(err.Error(), "2 Email Finder bulks running") {
		t.Fatalf("message = %q", err.Error())
	}
}

// The detail is {"data": job}, with the preview rows once billed.
func TestGetBulkJobDecodesDetail(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(detailJSON))
	}))
	defer srv.Close()
	job, err := New("k", "s").WithBaseURL(srv.URL).GetBulkJob(models.BulkTypeSearch, 9)
	if err != nil {
		t.Fatal(err)
	}
	if job.Status != "completed" || !job.Billed || job.Config.Options.MaxEmailsPerDomain != 10 ||
		job.Metrics.Results.Found != 2 || job.Metrics.Results.Records != 20 {
		t.Fatalf("job = %+v", job)
	}
	var preview []map[string]interface{}
	if err := json.Unmarshal(job.Preview, &preview); err != nil || preview[0]["email"] != "jane@stripe.com" {
		t.Fatalf("preview = %s (%v)", job.Preview, err)
	}
}

const detailJSON = `{"data":{"bulk_id":9,"user_id":7,"bulk_type":"search","name":"Q3","status":"completed",
	"archived":false,"progress":100,"processed":2,"billed":true,"error_message":null,"retry_count":0,
	"created_at":"2026-09-30T10:00:00Z","started_at":"2026-09-30T10:00:04Z","completed_at":"2026-09-30T10:00:46Z",
	"cancelled_at":null,"expired":false,"expired_at":"2027-03-29T10:00:00Z",
	"config":{"v":1,"input":{"source":"list","has_header":false,"rows":2},"options":{"max_emails_per_domain":10,"verify_emails":true},"notify":false},
	"metrics":{"v":1,"results":{"found":2,"not_found":0,"hit_rate":100,"records":20},"credits":{"search":2,"verify":20,"phone":0,"total":22}},
	"preview":[{"email":"jane@stripe.com","score":98}]}}`

func TestDownloadBulkToStreams(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("file") != "valid" {
			t.Errorf("query = %v", r.URL.Query())
		}
		_, _ = io.WriteString(w, "email\njane@stripe.com\n")
	}))
	defer srv.Close()
	var buf bytes.Buffer
	n, err := New("k", "s").WithBaseURL(srv.URL).DownloadBulkTo(models.BulkTypeVerifier, 3, "valid", &buf)
	if err != nil || n != int64(buf.Len()) || !strings.Contains(buf.String(), "jane@stripe.com") {
		t.Fatalf("n=%d err=%v body=%q", n, err, buf.String())
	}
}
