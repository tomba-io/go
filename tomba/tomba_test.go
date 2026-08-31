package tomba

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/tomba-io/go/tomba/models"
)

const (
	testDomain = "tomba.io"
	testEmail  = "b.mohamed@tomba.io"
	testQuery  = "tomba"
)

func skipWithoutCredentials(t *testing.T) {
	t.Helper()
	if os.Getenv("TOMBA_API_KEY") == "" {
		t.Skip("Skipping integration test: TOMBA_API_KEY not set")
	}
}

func newTestClient(t *testing.T) *Tomba {
	t.Helper()
	skipWithoutCredentials(t)
	key := os.Getenv("TOMBA_API_KEY")
	secret := os.Getenv("TOMBA_SECRET_KEY")
	return New(key, secret)
}

func TestNew(t *testing.T) {
	client := New("ta_test_key", "ts_test_secret")
	if client == nil {
		t.Fatal("New() returned nil")
	}
	if client.ApiKey != "ta_test_key" {
		t.Errorf("expected ApiKey 'ta_test_key', got '%s'", client.ApiKey)
	}
	if client.ApiSecret != "ts_test_secret" {
		t.Errorf("expected ApiSecret 'ts_test_secret', got '%s'", client.ApiSecret)
	}
}

func TestConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		expected string
	}{
		{"DEFAULT_BASE_URL", DEFAULT_BASE_URL, "https://api.tomba.io/v1"},
		{"ACCOUNT_PATH", ACCOUNT_PATH, "/me"},
		{"SEARCH_PATH", SEARCH_PATH, "/domain-search"},
		{"FINDER_PATH", FINDER_PATH, "/email-finder"},
		{"ENRICHMENT_PATH", ENRICHMENT_PATH, "/enrich"},
		{"AUTHOR_PATH", AUTHOR_PATH, "/author-finder"},
		{"LINKEDIN_PATH", LINKEDIN_PATH, "/linkedin"},
		{"VERIFIER_PATH", VERIFIER_PATH, "/email-verifier"},
		{"SOURCES_PATH", SOURCES_PATH, "/email-sources"},
		{"COUNT_PATH", COUNT_PATH, "/email-count"},
		{"STATUS_PATH", STATUS_PATH, "/domain-status"},
		{"AUTOCOMPLETE_PATH", AUTOCOMPLETE_PATH, "/domain-suggestions"},
		{"FORMAT_PATH", FORMAT_PATH, "/email-format"},
		{"EMPLOYEES_PATH", EMPLOYEES_PATH, "/location"},
		{"SIMILAR_PATH", SIMILAR_PATH, "/similar"},
		{"TECHNOLOGY_PATH", TECHNOLOGY_PATH, "/technology"},
		{"USAGE_PATH", USAGE_PATH, "/usage"},
		{"LOGS_PATH", LOGS_PATH, "/logs"},
		{"PHONE_FINDER_PATH", PHONE_FINDER_PATH, "/phone-finder"},
		{"PHONE_VALIDATOR_PATH", PHONE_VALIDATOR_PATH, "/phone-validator"},
		{"PEOPLE_FIND_PATH", PEOPLE_FIND_PATH, "/people/find"},
		{"COMPANIES_FIND_PATH", COMPANIES_FIND_PATH, "/companies/find"},
		{"COMBINED_FIND_PATH", COMBINED_FIND_PATH, "/combined/find"},
		{"LEADS_PATH", LEADS_PATH, "/leads"},
		{"LEADS_LISTS_PATH", LEADS_LISTS_PATH, "/leads_lists"},
		{"ATTRIBUTES_PATH", ATTRIBUTES_PATH, "/attributes"},
		{"KEYS_PATH", KEYS_PATH, "/keys"},
		{"FLAG_PATH", FLAG_PATH, "/flag"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %q, want %q", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestAccount(t *testing.T) {
	client := newTestClient(t)
	account, err := client.Account()
	if err != nil {
		t.Fatalf("Account() error: %v", err)
	}
	if account.Data.Email == "" {
		t.Error("Account() returned empty email")
	}
}

func TestDomainSearch(t *testing.T) {
	client := newTestClient(t)
	result, err := client.DomainSearch(Params{"domain": testDomain})
	if err != nil {
		t.Fatalf("DomainSearch() error: %v", err)
	}
	if len(result.Data.Emails) == 0 {
		t.Error("DomainSearch() returned no emails")
	}
}

func TestCount(t *testing.T) {
	client := newTestClient(t)
	result, err := client.Count(testDomain)
	if err != nil {
		t.Fatalf("Count() error: %v", err)
	}
	if result.Data.Total <= 0 {
		t.Error("Count() returned zero or negative total")
	}
}

func TestStatus(t *testing.T) {
	client := newTestClient(t)
	result, err := client.Status(testDomain)
	if err != nil {
		t.Fatalf("Status() error: %v", err)
	}
	if result.Domain == "" {
		t.Error("Status() returned empty domain")
	}
}

func TestEmailFinder(t *testing.T) {
	client := newTestClient(t)
	result, err := client.EmailFinder(Params{
		"domain":     testDomain,
		"first_name": "Mohamed",
		"last_name":  "Ben rebia",
	})
	if err != nil {
		t.Fatalf("EmailFinder() error: %v", err)
	}
	if result.Data == nil {
		t.Error("EmailFinder() returned nil data")
	}
}

func TestEnrichment(t *testing.T) {
	client := newTestClient(t)
	result, err := client.Enrichment(Params{"email": testEmail})
	if err != nil {
		t.Fatalf("Enrichment() error: %v", err)
	}
	if result.Data == nil {
		t.Error("Enrichment() returned nil data")
	}
}

func TestEmailVerifier(t *testing.T) {
	client := newTestClient(t)
	result, err := client.EmailVerifier(Params{"email": testEmail})
	if err != nil {
		t.Fatalf("EmailVerifier() error: %v", err)
	}
	if result.Data.Email.Email == "" {
		t.Error("EmailVerifier() returned empty email")
	}
}

func TestSources(t *testing.T) {
	client := newTestClient(t)
	_, err := client.Sources(testEmail)
	if err != nil {
		t.Fatalf("Sources() error: %v", err)
	}
}

func TestEmailFormat(t *testing.T) {
	client := newTestClient(t)
	_, err := client.EmailFormat(testDomain)
	if err != nil {
		t.Fatalf("EmailFormat() error: %v", err)
	}
}

func TestEmployeesCount(t *testing.T) {
	client := newTestClient(t)
	_, err := client.EmployeesCount(testDomain)
	if err != nil {
		t.Fatalf("EmployeesCount() error: %v", err)
	}
}

func TestSimilarDomains(t *testing.T) {
	client := newTestClient(t)
	_, err := client.SimilarDomains(testDomain)
	if err != nil {
		t.Fatalf("SimilarDomains() error: %v", err)
	}
}

func TestTechnologyCheck(t *testing.T) {
	client := newTestClient(t)
	_, err := client.TechnologyCheck(testDomain)
	if err != nil {
		t.Fatalf("TechnologyCheck() error: %v", err)
	}
}

func TestUsage(t *testing.T) {
	client := newTestClient(t)
	_, err := client.Usage()
	if err != nil {
		t.Fatalf("Usage() error: %v", err)
	}
}

func TestLogs(t *testing.T) {
	client := newTestClient(t)
	_, err := client.Logs()
	if err != nil {
		t.Fatalf("Logs() error: %v", err)
	}
}

func TestSearchCompanies(t *testing.T) {
	client := newTestClient(t)
	request := &models.RevealSearchRequest{
		Filters: &models.RevealFilters{
			LocationCountry: &models.SearchFilter{
				Include: []string{"US", "UK"},
			},
		},
		Page: 1,
	}
	result, err := client.SearchCompanies(request)
	if err != nil {
		t.Fatalf("SearchCompanies() error: %v", err)
	}
	if !result.Success {
		t.Error("SearchCompanies() returned success=false")
	}
}

func TestPersonFind(t *testing.T) {
	client := newTestClient(t)
	result, err := client.PersonFind(Params{"email": testEmail})
	if err != nil {
		t.Fatalf("PersonFind() error: %v", err)
	}
	if result.Data.Email == "" && result.Data.FullName == "" {
		t.Error("PersonFind() returned empty person data")
	}
}

func TestCompanyFind(t *testing.T) {
	client := newTestClient(t)
	_, err := client.CompanyFind(Params{"domain": testDomain})
	if err != nil {
		t.Fatalf("CompanyFind() error: %v", err)
	}
}

func TestCombinedFind(t *testing.T) {
	client := newTestClient(t)
	result, err := client.CombinedFind(Params{"email": testEmail})
	if err != nil {
		t.Fatalf("CombinedFind() error: %v", err)
	}
	_ = result
}

func TestListFlags(t *testing.T) {
	client := newTestClient(t)
	result, err := client.ListFlags(nil)
	if err != nil {
		t.Fatalf("ListFlags() error: %v", err)
	}
	if !json.Valid(result) {
		t.Error("ListFlags() returned invalid JSON")
	}
}

func TestListLeads(t *testing.T) {
	client := newTestClient(t)
	result, err := client.ListLeads(nil)
	if err != nil {
		t.Fatalf("ListLeads() error: %v", err)
	}
	if !json.Valid(result) {
		t.Error("ListLeads() returned invalid JSON")
	}
}

func TestListLeadsLists(t *testing.T) {
	client := newTestClient(t)
	result, err := client.ListLeadsLists(nil)
	if err != nil {
		t.Fatalf("ListLeadsLists() error: %v", err)
	}
	if !json.Valid(result) {
		t.Error("ListLeadsLists() returned invalid JSON")
	}
}

func TestListAttributes(t *testing.T) {
	client := newTestClient(t)
	result, err := client.ListAttributes(nil)
	if err != nil {
		t.Fatalf("ListAttributes() error: %v", err)
	}
	if !json.Valid(result) {
		t.Error("ListAttributes() returned invalid JSON")
	}
}

func TestListKeys(t *testing.T) {
	client := newTestClient(t)
	result, err := client.ListKeys()
	if err != nil {
		t.Fatalf("ListKeys() error: %v", err)
	}
	if !json.Valid(result) {
		t.Error("ListKeys() returned invalid JSON")
	}
}

func TestPhoneFinder(t *testing.T) {
	client := newTestClient(t)
	_, err := client.PhoneFinder(Params{"email": testEmail})
	if err != nil {
		t.Fatalf("PhoneFinder() error: %v", err)
	}
}

func TestPhoneValidator(t *testing.T) {
	client := newTestClient(t)
	_, err := client.PhoneValidator(Params{"phone": "+14155552671"})
	if err != nil {
		t.Fatalf("PhoneValidator() error: %v", err)
	}
}
