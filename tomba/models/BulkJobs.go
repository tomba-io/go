package models

import (
	"encoding/json"
	"time"
)

// BulkJobRequest creates a bulk job with the canonical parameters
// (https://docs.tomba.io/bulks). Send the input as exactly one of Data,
// List or FilePath. Field indexes are ZERO-based; nil means "auto-detect
// from the header row".
type BulkJobRequest struct {
	Name string `json:"name"`

	// Data is rows × columns; the first row is a header unless HasHeader is false.
	Data      [][]string `json:"data,omitempty"`
	HasHeader *bool      `json:"has_header,omitempty"`
	// List is one value per line (single-column input).
	List string `json:"list,omitempty"`
	// FilePath uploads a CSV file (multipart); Delimiter applies to it.
	FilePath  string `json:"-"`
	Delimiter string `json:"delimiter,omitempty"`

	DomainFieldIndex      *int `json:"domain_field_index,omitempty"`
	CompanyNameFieldIndex *int `json:"company_name_field_index,omitempty"`
	FirstNameFieldIndex   *int `json:"first_name_field_index,omitempty"`
	LastNameFieldIndex    *int `json:"last_name_field_index,omitempty"`
	FullNameFieldIndex    *int `json:"full_name_field_index,omitempty"`
	EmailFieldIndex       *int `json:"email_field_index,omitempty"`
	LinkedinURLFieldIndex *int `json:"linkedin_url_field_index,omitempty"`
	URLFieldIndex         *int `json:"url_field_index,omitempty"`
	PhoneFieldIndex       *int `json:"phone_field_index,omitempty"`
	CountryFieldIndex     *int `json:"country_field_index,omitempty"`
	TechnologyFieldIndex  *int `json:"technology_field_index,omitempty"`

	InputType             string   `json:"input_type,omitempty"` // "domain" | "company"
	MaxEmailsPerDomain    int      `json:"max_emails_per_domain,omitempty"`
	EmailType             string   `json:"email_type,omitempty"`      // all | personal | generic
	EmailTypeMode         string   `json:"email_type_mode,omitempty"` // only | priority
	Departments           []string `json:"departments,omitempty"`
	DepartmentsMode       string   `json:"departments_mode,omitempty"` // only | priority
	VerifyEmails          *bool    `json:"verify_emails,omitempty"`
	FindPhones            *bool    `json:"find_phones,omitempty"`
	IncludeSources        *bool    `json:"include_sources,omitempty"`
	BoostScoreFromSources *bool    `json:"boost_score_from_sources,omitempty"`
	FindAllPhones         *bool    `json:"find_all_phones,omitempty"`
	SkipRowsWithEmail     *bool    `json:"skip_rows_with_email,omitempty"`

	Notify     bool   `json:"notify,omitempty"`
	WebhookURL string `json:"webhook_url,omitempty"`
	// Launch starts the job right after creating it.
	Launch bool `json:"launch,omitempty"`
}

// BulkCreateResult is the answer to a create: the id and, when "launch"
// was requested but refused (quota, too many running jobs), why.
type BulkCreateResult struct {
	Data struct {
		ID      *int64 `json:"id"`
		Message string `json:"message,omitempty"`
	} `json:"data"`
}

// BulkConfig is a job's options (bulks.config).
type BulkConfig struct {
	V     int `json:"v"`
	Input struct {
		Source           string         `json:"source"`
		HasHeader        bool           `json:"has_header"`
		Delimiter        string         `json:"delimiter,omitempty"`
		OriginalFileName string         `json:"original_file_name,omitempty"`
		Rows             int            `json:"rows"`
		Fields           map[string]int `json:"fields,omitempty"`
	} `json:"input"`
	Options struct {
		InputType             string   `json:"input_type,omitempty"`
		MaxEmailsPerDomain    int      `json:"max_emails_per_domain,omitempty"`
		EmailType             string   `json:"email_type,omitempty"`
		EmailTypeMode         string   `json:"email_type_mode,omitempty"`
		Departments           []string `json:"departments,omitempty"`
		DepartmentsMode       string   `json:"departments_mode,omitempty"`
		VerifyEmails          bool     `json:"verify_emails"`
		FindPhones            bool     `json:"find_phones"`
		IncludeSources        bool     `json:"include_sources"`
		BoostScoreFromSources bool     `json:"boost_score_from_sources"`
		FindAllPhones         bool     `json:"find_all_phones"`
		SkipRowsWithEmail     bool     `json:"skip_rows_with_email"`
		Filter                string   `json:"filter,omitempty"`
	} `json:"options"`
	Notify  bool `json:"notify"`
	Webhook *struct {
		URL string `json:"url"`
	} `json:"webhook,omitempty"`
}

// BulkMetrics is a job's result distribution and metrics (bulks.metrics).
type BulkMetrics struct {
	V       int  `json:"v"`
	Partial bool `json:"partial,omitempty"`
	Rows    struct {
		Input        int `json:"input"`
		ValidInput   int `json:"valid_input"`
		InvalidInput int `json:"invalid_input"`
		Duplicates   int `json:"duplicates"`
		Skipped      int `json:"skipped"`
		Processed    int `json:"processed"`
		Errors       int `json:"errors"`
	} `json:"rows"`
	Results struct {
		Found    int     `json:"found"`
		NotFound int     `json:"not_found"`
		HitRate  float64 `json:"hit_rate"`
		// Records is the number of rows in the full results file (emails,
		// websites, phone numbers... depending on the type).
		Records int `json:"records"`
	} `json:"results"`
	Distribution struct {
		Rows         []BulkBucket `json:"rows,omitempty"`
		Verification []BulkBucket `json:"verification,omitempty"`
		EmailType    []BulkBucket `json:"email_type,omitempty"`
		Confidence   []BulkBucket `json:"confidence,omitempty"`
		Departments  []BulkBucket `json:"departments,omitempty"`
		PhoneType    []BulkBucket `json:"phone_type,omitempty"`
	} `json:"distribution"`
	Emails *struct {
		Total           int     `json:"total"`
		PerFoundRow     float64 `json:"per_found_row"`
		WithSources     int     `json:"with_sources"`
		Webmail         int     `json:"webmail"`
		DeliverableRate float64 `json:"deliverable_rate"`
	} `json:"emails,omitempty"`
	Phones *struct {
		Total         int `json:"total"`
		RowsWithPhone int `json:"rows_with_phone"`
	} `json:"phones,omitempty"`
	Errors struct {
		Client      int `json:"client"`
		Server      int `json:"server"`
		RateLimited int `json:"rate_limited"`
	} `json:"errors"`
	Credits BulkCredits `json:"credits"`
	Timing  struct {
		QueuedMs      int64   `json:"queued_ms"`
		DurationMs    int64   `json:"duration_ms"`
		RowsPerMinute float64 `json:"rows_per_minute"`
	} `json:"timing"`
}

// BulkBucket is one slice of a distribution.
type BulkBucket struct {
	Key     string  `json:"key"`
	Label   string  `json:"label"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

// BulkCredits: phone credits are billed as searches and included in Search.
type BulkCredits struct {
	Search int `json:"search"`
	Verify int `json:"verify"`
	Phone  int `json:"phone"`
	Total  int `json:"total"`
}

// BulkEstimate is what a job costs (exact once completed, an upper bound before).
type BulkEstimate struct {
	BulkID    int64       `json:"bulk_id"`
	Status    string      `json:"status"`
	Basis     string      `json:"basis"` // exact | upper_bound
	Billed    bool        `json:"billed"`
	Rows      int         `json:"rows"`
	Credits   BulkCredits `json:"credits"`
	Remaining *struct {
		Search int64 `json:"search"`
		Verify int64 `json:"verify"`
	} `json:"remaining"`
	Sufficient *bool  `json:"sufficient"`
	Note       string `json:"note"`
}

// BulkEvent is one entry of a job's history.
type BulkEvent struct {
	ID        int64           `json:"id"`
	UserID    *int64          `json:"user_id"`
	Event     string          `json:"event"`
	Actor     string          `json:"actor"`
	Message   *string         `json:"message"`
	Meta      json.RawMessage `json:"meta"`
	CreatedAt time.Time       `json:"created_at"`
}

// BulkTypeInfo describes one bulk type (GET /bulk/types).
type BulkTypeInfo struct {
	Type         string   `json:"type"`
	Path         string   `json:"path"`
	Label        string   `json:"label"`
	Description  string   `json:"description"`
	AppPath      string   `json:"app_path"`
	Storage      string   `json:"storage"`
	AllowsUpload bool     `json:"allows_upload"`
	FileRequired bool     `json:"file_required"`
	PrimaryField string   `json:"primary_field"`
	Options      []string `json:"options"`
	MaxRows      int      `json:"max_rows"`
	Fields       []struct {
		Field          string   `json:"field"`
		Param          string   `json:"param"`
		Label          string   `json:"label"`
		HeaderSynonyms []string `json:"header_synonyms"`
	} `json:"fields"`
	CreatesPerDay         int      `json:"creates_per_day"`
	DownloadsPerJobPerDay int      `json:"downloads_per_job_per_day"`
	AutoLaunch            bool     `json:"auto_launch"`
	Downloads             []string `json:"downloads"`
	Distributions         []string `json:"distributions"`
}

// BulkStats is account-level bulk activity (GET /bulk/stats).
type BulkStats struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
	Jobs struct {
		Total     int `json:"total"`
		Pending   int `json:"pending"`
		Running   int `json:"running"`
		Completed int `json:"completed"`
		Failed    int `json:"failed"`
		Cancelled int `json:"cancelled"`
	} `json:"jobs"`
	Rows    int         `json:"rows_processed"`
	Found   int         `json:"found"`
	Emails  int         `json:"emails_found"`
	HitRate float64     `json:"hit_rate"`
	Credits BulkCredits `json:"credits"`
	ByType  []struct {
		Type    string      `json:"type"`
		Label   string      `json:"label"`
		Jobs    int         `json:"jobs"`
		Rows    int         `json:"rows_processed"`
		Found   int         `json:"found"`
		HitRate float64     `json:"hit_rate"`
		Credits BulkCredits `json:"credits"`
	} `json:"by_type"`
	Daily []struct {
		Date    string `json:"date"`
		Jobs    int    `json:"jobs"`
		Rows    int    `json:"rows_processed"`
		Credits int    `json:"credits"`
	} `json:"daily"`
}

// BulkWebhookSecret signs completion webhooks.
type BulkWebhookSecret struct {
	Secret    string     `json:"secret"`
	CreatedAt time.Time  `json:"created_at"`
	RotatedAt *time.Time `json:"rotated_at"`
}
