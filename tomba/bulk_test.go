package tomba

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/tomba-io/go/tomba/models"
)

// The deprecated bulk methods keep their signatures but must send the API's
// current parameter names: it ignores the old ones.

// capture serves reply and records the last request's path, query, JSON body
// and multipart form.
type capture struct {
	path  string
	query url.Values
	json  map[string]interface{}
	form  map[string][]string
}

func serve(t *testing.T, reply string) (*Tomba, *capture) {
	t.Helper()
	got := &capture{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.path, got.query = r.URL.Path, r.URL.Query()
		switch ct := r.Header.Get("Content-Type"); {
		case r.Method == http.MethodGet:
		case len(ct) >= 19 && ct[:19] == "multipart/form-data":
			if err := r.ParseMultipartForm(1 << 20); err != nil {
				t.Error(err)
			}
			got.form = r.MultipartForm.Value
		default:
			if err := json.NewDecoder(r.Body).Decode(&got.json); err != nil {
				t.Error(err)
			}
		}
		_, _ = w.Write([]byte(reply))
	}))
	t.Cleanup(srv.Close)
	return New("k", "s").WithBaseURL(srv.URL + "/v1/"), got
}

func csvFile(t *testing.T) string {
	path := filepath.Join(t.TempDir(), "leads.csv")
	if err := os.WriteFile(path, []byte("a,b,c\n1,2,3\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func sameJSON(t *testing.T, got map[string]interface{}, want string) {
	t.Helper()
	var w map[string]interface{}
	if err := json.Unmarshal([]byte(want), &w); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, w) {
		g, _ := json.Marshal(got)
		t.Errorf("body = %s\nwant   %s", g, want)
	}
}

func sameForm(t *testing.T, got, want map[string][]string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("form = %v\nwant   %v", got, want)
	}
}

const created = `{"data":{"id":12}}`

func TestCreateBulkSendsCurrentNames(t *testing.T) {
	c, got := serve(t, created)
	res, err := c.CreateBulk(models.BulkTypeEnrich, &models.BulkCreateParams{
		Name: "Emails", List: "jane@stripe.com", Sources: true, Notifie: true, Verify: true,
		Total: 1, Delimiter: ";", Column: 2,
	})
	if err != nil || *res.Data.ID != 12 {
		t.Fatalf("res=%+v err=%v", res, err)
	}
	if got.path != "/v1/bulk/enrich" {
		t.Errorf("path = %s", got.path)
	}
	// Column is one-based: column 2 is field index 1. Total is not sent.
	sameJSON(t, got.json, `{"name":"Emails","list":"jane@stripe.com","delimiter":";",
		"include_sources":true,"notify":true,"verify_emails":true,"email_field_index":1}`)
}

// The API rejects an option a type does not take, even false: the old flags
// are only sent where they apply, like the old API only used them there.
func TestCreateBulkSendsOnlyTheTypesOptions(t *testing.T) {
	all := &models.BulkCreateParams{Name: "F", Sources: true, Verify: true, Valid: true}
	for typ, want := range map[models.BulkType]string{
		models.BulkTypeSearch:         `{"name":"F","include_sources":true,"verify_emails":true,"boost_score_from_sources":true}`,
		models.BulkTypeSimilar:        `{"name":"F","include_sources":true,"verify_emails":true}`,
		models.BulkTypeCompany:        `{"name":"F","verify_emails":true}`,
		models.BulkTypeFinder:         `{"name":"F","include_sources":true,"verify_emails":true,"boost_score_from_sources":true}`,
		models.BulkTypeVerifier:       `{"name":"F","include_sources":true}`,
		models.BulkTypePhoneFinder:    `{"name":"F"}`,
		models.BulkTypePhoneValidator: `{"name":"F"}`,
		models.BulkTypeTechnology:     `{"name":"F"}`,
	} {
		c, got := serve(t, created)
		if _, err := c.CreateBulk(typ, all); err != nil {
			t.Fatal(err)
		}
		sameJSON(t, got.json, want)
	}
}

func TestCreateBulkColumnIsTheTypesPrimaryField(t *testing.T) {
	for typ, param := range map[models.BulkType]string{
		models.BulkTypeSearch:         "domain_field_index",
		models.BulkTypeSimilar:        "domain_field_index",
		models.BulkTypeCompany:        "domain_field_index",
		models.BulkTypePhoneFinder:    "domain_field_index",
		models.BulkTypeEnrich:         "email_field_index",
		models.BulkTypeVerifier:       "email_field_index",
		models.BulkTypeLinkedIn:       "linkedin_url_field_index",
		models.BulkTypeAuthor:         "url_field_index",
		models.BulkTypePhoneValidator: "phone_field_index",
		models.BulkTypeTechnology:     "technology_field_index",
	} {
		c, got := serve(t, created)
		if _, err := c.CreateBulkWithFile(typ, &models.BulkCreateParams{Name: "F", Column: 1}, csvFile(t)); err != nil {
			t.Fatal(err)
		}
		sameForm(t, got.form, map[string][]string{"name": {"F"}, param: {"0"}})
	}
}

func TestCreateSearchBulkSendsFlatOptions(t *testing.T) {
	c, got := serve(t, created)
	_, err := c.CreateSearchBulk(&models.BulkSearchCreateParams{
		BulkCreateParams: models.BulkCreateParams{Name: "Domains", List: "stripe.com", Verify: true},
		BulkSearchParams: models.BulkSearchParams{
			Maximum:    "20",
			EmailType:  models.BulkSearchParamsType{Type: "personal", PriorityType: "priority"},
			Department: models.BulkSearchParamsDepartment{Name: []string{"sales", "it"}, PriorityType: "only"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	sameJSON(t, got.json, `{"name":"Domains","list":"stripe.com","verify_emails":true,
		"max_emails_per_domain":20,"email_type":"personal","email_type_mode":"priority",
		"departments":["sales","it"],"departments_mode":"only"}`)
}

func TestCreateSearchBulkLeavesUnsetOptionsToTheAPI(t *testing.T) {
	c, got := serve(t, created)
	_, err := c.CreateSearchBulk(&models.BulkSearchCreateParams{
		BulkCreateParams: models.BulkCreateParams{Name: "Domains", List: "stripe.com"},
	})
	if err != nil {
		t.Fatal(err)
	}
	sameJSON(t, got.json, `{"name":"Domains","list":"stripe.com"}`)
}

func TestCreateFinderBulkSendsZeroBasedFields(t *testing.T) {
	c, got := serve(t, created)
	_, err := c.CreateFinderBulk(&models.BulkFinderCreateParams{
		BulkCreateParams: models.BulkCreateParams{Name: "Leads", Verify: true, Column: 4},
		BulkFinderParams: models.BulkFinderParams{ColumnFirst: 1, ColumnLast: 2, ColumnDomain: 3, Skip: true},
	}, csvFile(t))
	if err != nil {
		t.Fatal(err)
	}
	if got.path != "/v1/bulk/finder" {
		t.Errorf("path = %s", got.path)
	}
	// Column means nothing for the finder: its columns are named.
	sameForm(t, got.form, map[string][]string{
		"name":                   {"Leads"},
		"verify_emails":          {"true"},
		"first_name_field_index": {"0"},
		"last_name_field_index":  {"1"},
		"domain_field_index":     {"2"},
		"skip_rows_with_email":   {"true"},
	})
}

func TestCreatePhoneValidatorBulkSendsZeroBasedFields(t *testing.T) {
	c, got := serve(t, created)
	_, err := c.CreatePhoneValidatorBulk(&models.BulkPhoneValidatorCreateParams{
		BulkCreateParams:         models.BulkCreateParams{Name: "Phones", Delimiter: ","},
		BulkPhoneValidatorParams: models.BulkPhoneValidatorParams{ColumnPhone: 1, ColumnCountry: 2},
	}, csvFile(t))
	if err != nil {
		t.Fatal(err)
	}
	sameForm(t, got.form, map[string][]string{
		"name": {"Phones"}, "delimiter": {","}, "phone_field_index": {"0"}, "country_field_index": {"1"},
	})
}

func TestGetAllBulksArchivedFilter(t *testing.T) {
	c, got := serve(t, `{"data":[],"meta":{"total":0,"pageSize":10,"current":2,"total_pages":0}}`)
	if _, err := c.GetAllBulks(models.BulkTypeVerifier, &models.BulkGetParams{Page: 2, Filter: "archived"}); err != nil {
		t.Fatal(err)
	}
	if want := (url.Values{"page": {"2"}, "archived": {"true"}}); !reflect.DeepEqual(got.query, want) {
		t.Errorf("query = %v, want %v", got.query, want)
	}

	if _, err := c.GetAllBulks(models.BulkTypeVerifier, &models.BulkGetParams{Filter: "all"}); err != nil {
		t.Fatal(err)
	}
	if len(got.query) != 0 {
		t.Errorf("query = %v, want none", got.query)
	}
}

func TestGetBulkWrapsTheJob(t *testing.T) {
	c, got := serve(t, detailJSON)
	res, err := c.GetBulk(models.BulkTypeSearch, 9)
	if err != nil {
		t.Fatal(err)
	}
	if got.path != "/v1/bulk/search/9" || len(res.Data) != 1 || res.Data[0].BulkID != 9 || !res.Data[0].Billed {
		t.Fatalf("path=%s res=%+v", got.path, res)
	}
}

func TestGetBulkProgressOfAPendingJob(t *testing.T) {
	c, _ := serve(t, `{"status":"pending","progress":0,"processed":0,"error_message":null,"rows_per_minute":null,"eta_seconds":null}`)
	p, err := c.GetBulkProgress(models.BulkTypeVerifier, 3)
	if err != nil || p.Status != models.BulkStatusPending {
		t.Fatalf("progress=%+v err=%v", p, err)
	}
}

func TestSaveBulkResultsAsksForTheFile(t *testing.T) {
	c, got := serve(t, "email\njane@stripe.com\n")
	path := filepath.Join(t.TempDir(), "out.csv")
	if err := c.SaveBulkResults(models.BulkTypeVerifier, 3, path, "not_found"); err != nil {
		t.Fatal(err)
	}
	if want := (url.Values{"file": {"not_found"}}); !reflect.DeepEqual(got.query, want) {
		t.Errorf("query = %v, want %v", got.query, want)
	}
	if b, _ := os.ReadFile(path); string(b) != "email\njane@stripe.com\n" {
		t.Errorf("file = %q", b)
	}
}
