package models

import (
	"encoding/json"
	"time"
)

type BulkType string

const (
	BulkTypeSearch         BulkType = "search"
	BulkTypeSimilar        BulkType = "similar"
	BulkTypeCompany        BulkType = "company"
	BulkTypeFinder         BulkType = "finder"
	BulkTypeEnrich         BulkType = "enrich"
	BulkTypeLinkedIn       BulkType = "linkedin"
	BulkTypeAuthor         BulkType = "author"
	BulkTypeVerifier       BulkType = "verifier"
	BulkTypePhoneFinder    BulkType = "phone-finder"
	BulkTypePhoneValidator BulkType = "phone-validator"
)

// Request Models

// BulkGetParams represents parameters for getting bulk operations
type BulkGetParams struct {
	Page      int    `json:"page,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	Direction string `json:"direction,omitempty"` // "desc" or "asc"
	Filter    string `json:"filter,omitempty"`    // "archived" or "all"
}

// BulkCreateParams represents parameters for creating bulk operations
type BulkCreateParams struct {
	Name      string `json:"name" form:"name"`
	List      string `json:"list,omitempty" form:"list"`
	Sources   bool   `json:"sources,omitempty" form:"sources"`
	Notifie   bool   `json:"notifie,omitempty" form:"notifie"`
	Verify    bool   `json:"verify,omitempty" form:"verify"`
	Total     int    `json:"total,omitempty" form:"total"`
	Delimiter string `json:"delimiter,omitempty" form:"delimiter"`
	Valid     bool   `json:"valid,omitempty" form:"valid"`
	Column    int    `json:"column,omitempty" form:"column"`
}

// BulkSearchParams represents parameters specific to search bulk operations
type BulkSearchParams struct {
	Maximum    string                     `json:"maximum" form:"maximum"` // "1", "5", "10", "20", "50", "100"
	EmailType  BulkSearchParamsType       `json:"email_type" form:"email_type"`
	Department BulkSearchParamsDepartment `json:"department" form:"department"`
}
type BulkSearchParamsType struct {
	Type         string `json:"type" form:"type"`                   // "all", "personal", "generic"
	PriorityType string `json:"priority_type" form:"priority_type"` // "only" or "priority"
}
type BulkSearchParamsDepartment struct {
	Name         []string `json:"name" form:"name"`                   // oneof=all engineering sales finance hr it marketing operations management executive legal support communication software security pr warehouse diversity administrative facilities accounting"
	PriorityType string   `json:"priority_type" form:"priority_type"` // "only" or "exclude"
}

func (r *BulkSearchParamsType) BulkSearchParamsTypeMarshal() ([]byte, error) {
	return json.Marshal(r)
}

// BulkFinderParams represents parameters specific to finder bulk operations
type BulkFinderParams struct {
	ColumnFirst  int  `json:"column_first,omitempty"`
	ColumnLast   int  `json:"column_last,omitempty"`
	ColumnName   int  `json:"column_name,omitempty"`
	ColumnDomain int  `json:"column_domain,omitempty"`
	Skip         bool `json:"skip,omitempty"`
}

// BulkPhoneValidatorParams represents parameters specific to phone validator bulk operations
type BulkPhoneValidatorParams struct {
	ColumnPhone   int `json:"column_phone,omitempty"`
	ColumnCountry int `json:"column_country,omitempty"`
}

// BulkExportParams represents parameters specific to export bulk operations
type BulkExportParams struct {
	Domain  string `json:"domain"`
	Type    string `json:"type,omitempty"` // "all", "personal", "generic", etc.
	Sources bool   `json:"sources,omitempty"`
}

// BulkRenameParams represents parameters for renaming bulk operations
type BulkRenameParams struct {
	Name string `json:"name"`
}

// BulkDownloadParams represents parameters for downloading bulk results
type BulkDownloadParams struct {
	Type string `json:"type,omitempty"` // "full", "not_found", "valid"
}

// BulkFileUpload represents file upload parameters
type BulkFileUpload struct {
	FilePath  string
	FieldName string
}

// Combined parameter types for specific bulk operations
type BulkSearchCreateParams struct {
	BulkCreateParams
	BulkSearchParams
}

type BulkFinderCreateParams struct {
	BulkCreateParams
	BulkFinderParams
}

type BulkPhoneValidatorCreateParams struct {
	BulkCreateParams
	BulkPhoneValidatorParams
}

type BulkExportCreateParams struct {
	BulkCreateParams
}

// Response Models

// BulkItem represents a bulk operation item
type BulkItem struct {
	BulkID      int64       `json:"bulk_id"`
	Name        string      `json:"name"`
	Maximum     *string     `json:"maximum,omitempty"`
	EmailType   *string     `json:"email_type,omitempty"`
	Department  *string     `json:"department,omitempty"`
	Sources     *bool       `json:"sources,omitempty"`
	FileName    string      `json:"file_name"`
	Total       *int        `json:"total,omitempty"`
	TotalList   *int        `json:"total_list,omitempty"`
	Status      bool        `json:"status"`
	Chart       interface{} `json:"chart,omitempty"`
	Table       interface{} `json:"table,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	UserID      int64       `json:"user_id"`
	Launched    *bool       `json:"launched,omitempty"`
	Used        bool        `json:"used"`
	TimeTrack   *string     `json:"time_track,omitempty"`
	Progress    int         `json:"progress"`
	Verify      *bool       `json:"verify,omitempty"`
	VerifyCost  *int        `json:"verify_cost,omitempty"`
	TotalEmails *int        `json:"total_emails,omitempty"`
	Processed   int         `json:"processed"`
	ExpiredAt   time.Time   `json:"expired_at"`
	Expired     bool        `json:"expired"`
	BulkType    string      `json:"bulk_type"`
}

// BulkProgress represents bulk operation progress
type BulkProgress struct {
	Status         bool `json:"status"`
	Progress       int  `json:"progress"`
	Processed      int  `json:"processed"`
	ProcessedEmail int  `json:"processed_email"`
}

// BulkListResponse represents the response for bulk list operations
type BulkListResponse struct {
	Data []BulkItem `json:"data"`
	Meta struct {
		Total      int64   `json:"total"`
		PageSize   int     `json:"pageSize"`
		Current    int     `json:"current"`
		TotalPages float64 `json:"total_pages"`
	} `json:"meta"`
}

// BulkDetailResponse represents the response for bulk detail operations
type BulkDetailResponse struct {
	Data []BulkItem `json:"data"`
}

// BulkCreateResponse represents the response for bulk creation
type BulkCreateResponse struct {
	Data struct {
		ID *int64 `json:"id"`
	} `json:"data"`
}

// BulkSuccessResponse represents a successful bulk operation response
type BulkSuccessResponse struct {
	Data struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"data"`
}
