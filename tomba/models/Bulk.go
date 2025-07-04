package models

// BulkGetParams represents parameters for getting bulk operations
type BulkGetParams struct {
	Page      int    `json:"page,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	Direction string `json:"direction,omitempty"` // "desc" or "asc"
	Filter    string `json:"filter,omitempty"`    // "archived" or "all"
}

// BulkCreateParams represents parameters for creating bulk operations
type BulkCreateParams struct {
	Name      string `json:"name"`
	List      string `json:"list,omitempty"`
	Sources   bool   `json:"sources,omitempty"`
	Notifie   bool   `json:"notifie,omitempty"`
	Verify    bool   `json:"verify,omitempty"`
	Total     int    `json:"total,omitempty"`
	Delimiter string `json:"delimiter,omitempty"`
	Valid     bool   `json:"valid,omitempty"`
	Column    int    `json:"column,omitempty"`
}

// BulkSearchParams represents parameters specific to search bulk operations
type BulkSearchParams struct {
	Maximum    string `json:"maximum"` // "1", "5", "10", "20", "50", "100"
	EmailType  string `json:"email_type"`
	Department string `json:"department"`
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
	CreatedAt   string      `json:"created_at"`
	UserID      int64       `json:"user_id"`
	Launched    *bool       `json:"launched,omitempty"`
	Used        bool        `json:"used"`
	TimeTrack   *string     `json:"time_track,omitempty"`
	Progress    int         `json:"progress"`
	Verify      *bool       `json:"verify,omitempty"`
	VerifyCost  *int        `json:"verify_cost,omitempty"`
	TotalEmails *int        `json:"total_emails,omitempty"`
	Processed   int         `json:"processed"`
	ExpiredAt   string      `json:"expired_at"`
	Expired     bool        `json:"expired"`
	BulkType    string      `json:"bulk_type"`
}

// BulkProgress represents bulk operation progress
type BulkProgress struct {
	Status    bool `json:"status"`
	Progress  int  `json:"progress"`
	Processed int  `json:"processed"`
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
