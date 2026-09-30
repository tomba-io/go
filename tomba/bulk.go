package tomba

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/tomba-io/go/tomba/models"
)

// getBulkPath converts a BulkType to its API path.
func getBulkPath(bulkType models.BulkType) string {
	return fmt.Sprintf(BULK_PATH, string(bulkType))
}

// The methods below predate CreateBulkJob and friends (bulk_jobs.go). They
// keep their signatures but send the API's current parameter names: the
// API ignores the old ones (sources, notifie, verify, valid, maximum,
// column_*...).

// legacyColumnParam is the field the one-based Column parameter meant for
// each type; it is sent as a zero-based <field>_field_index.
func legacyColumnParam(t models.BulkType) string {
	switch t {
	case models.BulkTypeSearch, models.BulkTypeSimilar, models.BulkTypeCompany, models.BulkTypePhoneFinder:
		return "domain_field_index"
	case models.BulkTypeEnrich, models.BulkTypeVerifier:
		return "email_field_index"
	case models.BulkTypeLinkedIn:
		return "linkedin_url_field_index"
	case models.BulkTypeAuthor:
		return "url_field_index"
	case models.BulkTypePhoneValidator:
		return "phone_field_index"
	case models.BulkTypeTechnology:
		return "technology_field_index"
	}
	return "" // finder: its columns are named
}

// legacyOptions are, per type, the options the old flags map to that the
// type takes (GET /bulk/types). The old API ignored a flag a type did not
// use; the current one rejects such an option, so it is not sent.
var legacyOptions = map[models.BulkType][]string{
	models.BulkTypeSearch:   {"include_sources", "verify_emails", "boost_score_from_sources"},
	models.BulkTypeSimilar:  {"include_sources", "verify_emails"},
	models.BulkTypeCompany:  {"verify_emails"},
	models.BulkTypeFinder:   {"include_sources", "verify_emails", "boost_score_from_sources"},
	models.BulkTypeEnrich:   {"include_sources", "verify_emails"},
	models.BulkTypeLinkedIn: {"include_sources", "verify_emails"},
	models.BulkTypeAuthor:   {"include_sources", "verify_emails"},
	models.BulkTypeVerifier: {"include_sources"},
	models.BulkTypeExport:   {"include_sources"},
}

// createParams translates BulkCreateParams. Options are only sent when true
// and taken by the type.
func createParams(t models.BulkType, p *models.BulkCreateParams) Params {
	out := make(Params)
	if p == nil {
		return out
	}
	out["name"] = p.Name
	if p.List != "" {
		out["list"] = p.List
	}
	if p.Delimiter != "" {
		out["delimiter"] = p.Delimiter
	}
	setTrue(out, "notify", p.Notifie)
	flags := map[string]bool{
		"include_sources":          p.Sources,
		"verify_emails":            p.Verify,
		"boost_score_from_sources": p.Valid,
	}
	for _, option := range legacyOptions[t] {
		setTrue(out, option, flags[option])
	}
	if name := legacyColumnParam(t); name != "" {
		setColumn(out, name, p.Column)
	}
	return out
}

func setTrue(p Params, key string, v bool) {
	if v {
		p[key] = true
	}
}

// setColumn sends a one-based column as the zero-based <field>_field_index.
func setColumn(p Params, key string, oneBased int) {
	if oneBased > 0 {
		p[key] = oneBased - 1
	}
}

// GetAllBulks retrieves all bulk operations for a specific type.
// Filter "archived" lists the archived jobs.
//
// Deprecated: use ListBulkJobs.
func (conf *Tomba) GetAllBulks(bulkType models.BulkType, params *models.BulkGetParams) (*models.BulkListResponse, error) {
	path := getBulkPath(bulkType)

	requestParams := make(Params)
	if params != nil {
		if params.Page > 0 {
			requestParams["page"] = strconv.Itoa(params.Page)
		}
		if params.Limit > 0 {
			requestParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Direction != "" {
			requestParams["direction"] = params.Direction
		}
		if params.Filter == "archived" {
			requestParams["archived"] = "true"
		}
	}

	method := "GET"
	resp, err := conf.TombaCall(path, requestParams, &method, nil)
	if err != nil {
		return nil, err
	}

	var response models.BulkListResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// GetBulk retrieves a specific bulk operation by its type and ID, as a
// one-element Data list.
//
// Deprecated: use GetBulkJob.
func (conf *Tomba) GetBulk(bulkType models.BulkType, id int64) (*models.BulkDetailResponse, error) {
	path := fmt.Sprintf(BULK_PATH+"/%d", string(bulkType), id)

	method := "GET"
	resp, err := conf.TombaCall(path, Params{}, &method, nil)
	if err != nil {
		return nil, err
	}

	var one struct {
		Data models.BulkItem `json:"data"`
	}
	if err := json.Unmarshal(resp, &one); err != nil {
		return nil, err
	}
	return &models.BulkDetailResponse{Data: []models.BulkItem{one.Data}}, nil
}

// CreateBulk creates a new bulk operation of the specified type (not
// launched: call LaunchBulk).
//
// Deprecated: use CreateBulkJob.
func (conf *Tomba) CreateBulk(bulkType models.BulkType, params *models.BulkCreateParams) (*models.BulkCreateResponse, error) {
	path := getBulkPath(bulkType)
	requestParams := createParams(bulkType, params)

	method := "POST"
	resp, err := conf.TombaCall(path, requestParams, &method, nil)
	if err != nil {
		return nil, err
	}

	var response models.BulkCreateResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// CreateBulkWithFile creates a new bulk operation with a file upload (not
// launched: call LaunchBulk).
//
// Deprecated: use CreateBulkJob with FilePath.
func (conf *Tomba) CreateBulkWithFile(bulkType models.BulkType, params *models.BulkCreateParams, filePath string) (*models.BulkCreateResponse, error) {
	path := getBulkPath(bulkType)
	requestParams := createParams(bulkType, params)

	fileUpload := &models.BulkFileUpload{
		FilePath:  filePath,
		FieldName: "file",
	}

	method := "POST"
	resp, err := conf.TombaCall(path, requestParams, &method, fileUpload)
	if err != nil {
		return nil, err
	}

	var response models.BulkCreateResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// CreateSearchBulk creates a new search bulk operation (not launched: call
// LaunchBulk).
//
// Deprecated: use CreateBulkJob.
func (conf *Tomba) CreateSearchBulk(params *models.BulkSearchCreateParams) (*models.BulkCreateResponse, error) {
	path := getBulkPath(models.BulkTypeSearch)
	var requestParams Params
	if params != nil {
		requestParams = createParams(models.BulkTypeSearch, &params.BulkCreateParams)
		if n, err := strconv.Atoi(params.Maximum); err == nil && n > 0 {
			requestParams["max_emails_per_domain"] = n
		}
		if params.EmailType.Type != "" {
			requestParams["email_type"] = params.EmailType.Type
		}
		if params.EmailType.PriorityType != "" {
			requestParams["email_type_mode"] = params.EmailType.PriorityType
		}
		if len(params.Department.Name) > 0 {
			requestParams["departments"] = params.Department.Name
		}
		if params.Department.PriorityType != "" {
			requestParams["departments_mode"] = params.Department.PriorityType
		}
	}

	method := "POST"
	resp, err := conf.TombaCall(path, requestParams, &method, nil)
	if err != nil {
		return nil, err
	}

	var response models.BulkCreateResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// CreateFinderBulk creates a new finder bulk operation with file upload (not
// launched: call LaunchBulk). Its Column* fields are one-based.
//
// Deprecated: use CreateBulkJob.
func (conf *Tomba) CreateFinderBulk(params *models.BulkFinderCreateParams, filePath string) (*models.BulkCreateResponse, error) {
	path := getBulkPath(models.BulkTypeFinder)

	var requestParams Params
	if params != nil {
		requestParams = createParams(models.BulkTypeFinder, &params.BulkCreateParams)
		setColumn(requestParams, "first_name_field_index", params.ColumnFirst)
		setColumn(requestParams, "last_name_field_index", params.ColumnLast)
		setColumn(requestParams, "full_name_field_index", params.ColumnName)
		setColumn(requestParams, "domain_field_index", params.ColumnDomain)
		setTrue(requestParams, "skip_rows_with_email", params.Skip)
	}

	fileUpload := &models.BulkFileUpload{
		FilePath:  filePath,
		FieldName: "file",
	}

	method := "POST"
	resp, err := conf.TombaCall(path, requestParams, &method, fileUpload)
	if err != nil {
		return nil, err
	}

	var response models.BulkCreateResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// CreatePhoneValidatorBulk creates a new phone validator bulk operation (not
// launched: call LaunchBulk). Its Column* fields are one-based.
//
// Deprecated: use CreateBulkJob.
func (conf *Tomba) CreatePhoneValidatorBulk(params *models.BulkPhoneValidatorCreateParams, filePath string) (*models.BulkCreateResponse, error) {
	path := getBulkPath(models.BulkTypePhoneValidator)

	var requestParams Params
	if params != nil {
		requestParams = createParams(models.BulkTypePhoneValidator, &params.BulkCreateParams)
		setColumn(requestParams, "phone_field_index", params.ColumnPhone)
		setColumn(requestParams, "country_field_index", params.ColumnCountry)
	}

	fileUpload := &models.BulkFileUpload{
		FilePath:  filePath,
		FieldName: "file",
	}

	method := "POST"
	resp, err := conf.TombaCall(path, requestParams, &method, fileUpload)
	if err != nil {
		return nil, err
	}

	var response models.BulkCreateResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// LaunchBulk launches a bulk operation for processing.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) LaunchBulk(bulkType models.BulkType, id int64) (*models.BulkSuccessResponse, error) {
	path := fmt.Sprintf(BULK_PATH+"/%d", string(bulkType), id)

	method := "PUT"
	resp, err := conf.TombaCall(path, Params{}, &method, nil)
	if err != nil {
		return nil, err
	}

	var response models.BulkSuccessResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// DeleteBulk deletes a bulk operation by type and ID.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) DeleteBulk(bulkType models.BulkType, id int64) (*models.BulkSuccessResponse, error) {
	path := fmt.Sprintf(BULK_PATH+"/%d/delete", string(bulkType), id)

	method := "DELETE"
	resp, err := conf.TombaCall(path, Params{}, &method, nil)
	if err != nil {
		return nil, err
	}

	var response models.BulkSuccessResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// ArchiveBulk archives a bulk operation by type and ID.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) ArchiveBulk(bulkType models.BulkType, id int64) (*models.BulkSuccessResponse, error) {
	path := fmt.Sprintf(BULK_PATH+"/%d/archive", string(bulkType), id)

	method := "DELETE"
	resp, err := conf.TombaCall(path, Params{}, &method, nil)
	if err != nil {
		return nil, err
	}

	var response models.BulkSuccessResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// RenameBulk renames a bulk operation by type and ID.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) RenameBulk(bulkType models.BulkType, id int64, params *models.BulkRenameParams) (*models.BulkSuccessResponse, error) {
	path := fmt.Sprintf(BULK_PATH+"/%d/rename", string(bulkType), id)

	requestParams := make(Params)
	if params != nil {
		requestParams["name"] = params.Name
	}

	method := "PUT"
	resp, err := conf.TombaCall(path, requestParams, &method, nil)
	if err != nil {
		return nil, err
	}

	var response models.BulkSuccessResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// GetBulkProgress returns the progress of a bulk operation.
//
// Deprecated: use GetBulkJobProgress.
func (conf *Tomba) GetBulkProgress(bulkType models.BulkType, id int64) (*models.BulkProgress, error) {
	path := fmt.Sprintf(BULK_PATH+"/%d/progress", string(bulkType), id)

	method := "GET"
	resp, err := conf.TombaCall(path, Params{}, &method, nil)
	if err != nil {
		return nil, err
	}

	var response models.BulkProgress
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// DownloadBulk downloads the results of a bulk operation. params.Type is the
// file: "full" (default), "valid" or "not_found".
//
// Deprecated: use DownloadBulkTo.
func (conf *Tomba) DownloadBulk(bulkType models.BulkType, id int64, params *models.BulkDownloadParams) ([]byte, error) {
	path := fmt.Sprintf(BULK_PATH+"/%d/download", string(bulkType), id)

	requestParams := make(Params)
	if params != nil && params.Type != "" {
		requestParams["file"] = params.Type
	}

	method := "GET"
	return conf.TombaCall(path, requestParams, &method, nil)
}

// GetAllSearchBulks retrieves all search bulk operations.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetAllSearchBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeSearch, params)
}

// GetAllSimilarBulks retrieves all similar bulk operations.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetAllSimilarBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeSimilar, params)
}

// GetAllCompanyBulks retrieves all company bulk operations.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetAllCompanyBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeCompany, params)
}

// GetAllFinderBulks retrieves all finder bulk operations.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetAllFinderBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeFinder, params)
}

// GetAllEnrichBulks retrieves all enrich bulk operations.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetAllEnrichBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeEnrich, params)
}

// GetAllLinkedInBulks retrieves all LinkedIn bulk operations.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetAllLinkedInBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeLinkedIn, params)
}

// GetAllAuthorBulks retrieves all author bulk operations.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetAllAuthorBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeAuthor, params)
}

// GetAllVerifierBulks retrieves all verifier bulk operations.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetAllVerifierBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeVerifier, params)
}

// GetAllPhoneFinderBulks retrieves all phone finder bulk operations.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetAllPhoneFinderBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypePhoneFinder, params)
}

// GetAllPhoneValidatorBulks retrieves all phone validator bulk operations.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetAllPhoneValidatorBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypePhoneValidator, params)
}

// SaveBulkResults downloads bulk results and saves them to a local file.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) SaveBulkResults(bulkType models.BulkType, id int64, filePath string, downloadType string) error {
	params := &models.BulkDownloadParams{}
	if downloadType != "" {
		params.Type = downloadType
	}

	data, err := conf.DownloadBulk(bulkType, id, params)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = file.Write(data)
	return err
}
