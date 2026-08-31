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

// GetAllBulks retrieves all bulk operations for a specific type.
// See https://docs.tomba.io/api/bulk
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
		if params.Filter != "" {
			requestParams["filter"] = params.Filter
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

// GetBulk retrieves a specific bulk operation by its type and ID.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) GetBulk(bulkType models.BulkType, id int64) (*models.BulkDetailResponse, error) {
	path := fmt.Sprintf(BULK_PATH+"/%d", string(bulkType), id)

	method := "GET"
	resp, err := conf.TombaCall(path, Params{}, &method, nil)
	if err != nil {
		return nil, err
	}

	var response models.BulkDetailResponse
	err = json.Unmarshal(resp, &response)
	return &response, err
}

// CreateBulk creates a new bulk operation of the specified type.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) CreateBulk(bulkType models.BulkType, params *models.BulkCreateParams) (*models.BulkCreateResponse, error) {
	path := getBulkPath(bulkType)

	requestParams := make(Params)
	if params != nil {
		requestParams["name"] = params.Name
		if params.List != "" {
			requestParams["list"] = params.List
		}
		requestParams["sources"] = strconv.FormatBool(params.Sources)
		requestParams["notifie"] = strconv.FormatBool(params.Notifie)
		requestParams["verify"] = strconv.FormatBool(params.Verify)
		if params.Total > 0 {
			requestParams["total"] = strconv.Itoa(params.Total)
		}
		if params.Delimiter != "" {
			requestParams["delimiter"] = params.Delimiter
		}
		requestParams["valid"] = strconv.FormatBool(params.Valid)
		if params.Column > 0 {
			requestParams["column"] = strconv.Itoa(params.Column)
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

// CreateBulkWithFile creates a new bulk operation with a file upload.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) CreateBulkWithFile(bulkType models.BulkType, params *models.BulkCreateParams, filePath string) (*models.BulkCreateResponse, error) {
	path := getBulkPath(bulkType)

	requestParams := make(Params)
	if params != nil {
		requestParams["name"] = params.Name
		if params.List != "" {
			requestParams["list"] = params.List
		}
		requestParams["sources"] = strconv.FormatBool(params.Sources)
		requestParams["notifie"] = strconv.FormatBool(params.Notifie)
		requestParams["verify"] = strconv.FormatBool(params.Verify)
		if params.Total > 0 {
			requestParams["total"] = strconv.Itoa(params.Total)
		}
		if params.Delimiter != "" {
			requestParams["delimiter"] = params.Delimiter
		}
		requestParams["valid"] = strconv.FormatBool(params.Valid)
		if params.Column > 0 {
			requestParams["column"] = strconv.Itoa(params.Column)
		}
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

// CreateSearchBulk creates a new search bulk operation.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) CreateSearchBulk(params *models.BulkSearchCreateParams) (*models.BulkCreateResponse, error) {
	path := getBulkPath(models.BulkTypeSearch)
	requestParams := make(Params)
	if params != nil {
		// Basic bulk params
		requestParams["name"] = params.Name
		if params.List != "" {
			requestParams["list"] = params.List
		}
		requestParams["sources"] = (params.Sources)
		requestParams["verify"] = (params.Verify)

		// Search-specific params
		requestParams["maximum"] = params.Maximum
		requestParams["email_type"] = params.EmailType
		requestParams["department"] = params.Department
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

// CreateFinderBulk creates a new finder bulk operation with file upload.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) CreateFinderBulk(params *models.BulkFinderCreateParams, filePath string) (*models.BulkCreateResponse, error) {
	path := getBulkPath(models.BulkTypeFinder)

	requestParams := make(Params)
	if params != nil {
		// Basic bulk params
		requestParams["name"] = params.Name
		if params.Delimiter != "" {
			requestParams["delimiter"] = params.Delimiter
		}
		requestParams["verify"] = strconv.FormatBool(params.Verify)

		// Finder-specific params
		if params.ColumnFirst > 0 {
			requestParams["column_first"] = strconv.Itoa(params.ColumnFirst)
		}
		if params.ColumnLast > 0 {
			requestParams["column_last"] = strconv.Itoa(params.ColumnLast)
		}
		if params.ColumnName > 0 {
			requestParams["column_name"] = strconv.Itoa(params.ColumnName)
		}
		if params.ColumnDomain > 0 {
			requestParams["column_domain"] = strconv.Itoa(params.ColumnDomain)
		}
		requestParams["skip"] = strconv.FormatBool(params.Skip)
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

// CreatePhoneValidatorBulk creates a new phone validator bulk operation.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) CreatePhoneValidatorBulk(params *models.BulkPhoneValidatorCreateParams, filePath string) (*models.BulkCreateResponse, error) {
	path := getBulkPath(models.BulkTypePhoneValidator)

	requestParams := make(Params)
	if params != nil {
		// Basic bulk params
		requestParams["name"] = params.Name
		if params.Delimiter != "" {
			requestParams["delimiter"] = params.Delimiter
		}

		// Phone validator specific params
		if params.ColumnPhone > 0 {
			requestParams["column_phone"] = strconv.Itoa(params.ColumnPhone)
		}
		if params.ColumnCountry > 0 {
			requestParams["column_country"] = strconv.Itoa(params.ColumnCountry)
		}
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
// See https://docs.tomba.io/api/bulk
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

// DownloadBulk downloads the results of a bulk operation.
// See https://docs.tomba.io/api/bulk
func (conf *Tomba) DownloadBulk(bulkType models.BulkType, id int64, params *models.BulkDownloadParams) ([]byte, error) {
	path := fmt.Sprintf(BULK_PATH+"/%d/download", string(bulkType), id)

	requestParams := make(Params)
	if params != nil && params.Type != "" {
		requestParams["type"] = params.Type
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
