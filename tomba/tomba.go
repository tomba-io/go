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

	"github.com/tomba-io/go/tomba/models"
)

// TombaCall to Tomba Api
// get data
// TombaCall makes API calls with support for file uploads
func (conf *Tomba) TombaCall(path string, params Params, method *string, fileUpload *models.BulkFileUpload) ([]byte, error) {
	var req *http.Request
	var err error
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
			switch v := any(value).(type) {
			case string:
				writer.WriteField(key, v)
			case int:
				writer.WriteField(key, strconv.Itoa(v))
			case bool:
				writer.WriteField(key, strconv.FormatBool(v))
			case float64:
				writer.WriteField(key, strconv.FormatFloat(v, 'f', -1, 64))
			}
		}

		// Add file
		file, err := os.Open(fileUpload.FilePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(fileUpload.FieldName, fileUpload.FilePath)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return nil, err
		}

		err = writer.Close()
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(methodStr, apiUrl, body)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else if methodStr == "GET" {
		// Handle GET requests with query parameters
		req, err = http.NewRequest(methodStr, apiUrl, nil)
		if err != nil {
			return nil, err
		}

		if len(params) > 0 {
			q := req.URL.Query()
			for key, value := range params {
				switch v := any(value).(type) {
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
		jsonData, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(methodStr, apiUrl, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, err
		}
	}

	// Set common headers
	req.Header.Set("X-Tomba-Key", conf.ApiKey)
	req.Header.Set("X-Tomba-Secret", conf.ApiSecret)
	req.Header.Set("User-Agent", "tomba-go-client/1.0")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Print("Request URL: ", req.URL.String(), "\n", "Method: ", req.Method, "\n", "Headers: ", req.Header, "\n")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// fmt.Print("Response Status: ", resp.Status, "\n", "Response Headers: ", resp.Header, "\n")
	if resp.Body != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Reset the body for further reading
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("error: %s, status code: %d", resp.Status, resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// Account Returns information about the current account.
func (conf *Tomba) Account() (models.Account, error) {
	account := models.Account{}
	str, err := conf.TombaCall(ACCOUNT_PATH, nil, nil, nil)
	if err != nil {
		return account, err
	}
	data, err := models.UnmarshalAccount([]byte(str))
	if err != nil {
		return account, err
	}
	return data, nil
}

// DomainSearch Search emails are based on the website You give one domain name and it returns all the email addresses found on the internet.
func (conf *Tomba) DomainSearch(params Params) (models.Search, error) {

	search := models.Search{}
	str, err := conf.TombaCall(SEARCH_PATH, params, nil, nil)
	if err != nil {
		return search, err
	}
	data, err := models.UnmarshalSearch([]byte(str))
	if err != nil {
		return search, err
	}
	return data, nil
}

// Count Returns total email addresses we have for one domain.
func (conf *Tomba) Count(domain string) (models.Count, error) {
	count := models.Count{}
	str, err := conf.TombaCall(COUNT_PATH, Params{"domain": domain}, nil, nil)
	if err != nil {
		return count, err
	}
	data, err := models.UnmarshalCount([]byte(str))
	if err != nil {
		return count, err
	}
	return data, nil
}

// Status Returns domain status if is webmail or disposable.
func (conf *Tomba) Status(domain string) (models.Status, error) {
	status := models.Status{}
	str, err := conf.TombaCall(STATUS_PATH, Params{"domain": domain}, nil, nil)
	if err != nil {
		return status, err
	}
	data, err := models.UnmarshalStatus([]byte(str))
	if err != nil {
		return status, err
	}
	return data, nil
}

// EmailFinder Generates or retrieves the most likely email address from a domain name, a first name and a last name.
func (conf *Tomba) EmailFinder(params Params) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(FINDER_PATH, params, nil, nil)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder([]byte(str))
	if err != nil {
		return finder, err
	}
	return data, nil
}

// Enrichment The API lets you look up person and company data based on an email, For example, you could retrieve a person’s name, location and social handles from an email
// Parameters: email (required), enrich_mobile (optional - set to true to get phone number)
func (conf *Tomba) Enrichment(params Params) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(ENRICHMENT_PATH, params, nil, nil)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder([]byte(str))
	if err != nil {
		return finder, err
	}
	return data, nil
}

// AuthorFinder This API generates or retrieves the most likely email address from a blog post url.
func (conf *Tomba) AuthorFinder(url string) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(AUTHOR_PATH, Params{"url": url}, nil, nil)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder([]byte(str))
	if err != nil {
		return finder, err
	}
	return data, nil
}

// LinkedinFinder This API point generates or retrieves the most likely email address from a Linkedin URL.
// Parameters: url (required), enrich_mobile (optional - set to true to get phone number)
func (conf *Tomba) LinkedinFinder(params Params) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(LINKEDIN_PATH, params, nil, nil)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder([]byte(str))
	if err != nil {
		return finder, err
	}
	return data, nil
}

// PhoneFinder Search for phone numbers based on an email, domain, or LinkedIn URL.
// Parameters:
//   - email (optional): The email address you want to find phone for
//   - domain (optional): Domain name from which you want to find the phone numbers (e.g., stripe.com)
//   - linkedin (optional): The URL of the LinkedIn profile (e.g., https://www.linkedin.com/in/alex-maccaw-ab592978)
//   - full (optional): Set to true to get an array of all phone numbers associated with the email/domain/LinkedIn URL
//
// At least one of email, domain, or linkedin must be provided.
func (conf *Tomba) PhoneFinder(params Params) (models.Phone, error) {
	phone := models.Phone{}
	str, err := conf.TombaCall(PHONE_FINDER_PATH, params, nil, nil)
	if err != nil {
		return phone, err
	}
	data, err := models.UnmarshalPhone([]byte(str))
	if err != nil {
		return phone, err
	}
	return data, nil
}

// PhoneValidator validates a phone number and retrieves its associated information.
// Parameters:
//   - phone (required): The phone number you want to validate
//   - country_code (optional): Country code of the phone number (e.g., "US")
//
// see https://docs.tomba.io/api/phone#phone-validator
func (conf *Tomba) PhoneValidator(params Params) (models.PhoneValidator, error) {
	phoneValidator := models.PhoneValidator{}
	str, err := conf.TombaCall(PHONE_VALIDATOR_PATH, params, nil, nil)
	if err != nil {
		return phoneValidator, err
	}
	data, err := models.UnmarshalPhoneValidator([]byte(str))
	if err != nil {
		return phoneValidator, err
	}
	return data, nil
}

// EmailVerifier Verify the deliverability of an email address.
// Parameters: email (required), enrich_mobile (optional - set to true to get phone number)
func (conf *Tomba) EmailVerifier(params Params) (models.Verifier, error) {
	verifier := models.Verifier{}
	str, err := conf.TombaCall(VERIFIER_PATH, params, nil, nil)
	if err != nil {
		return verifier, err
	}
	data, err := models.UnmarshalVerifier([]byte(str))
	if err != nil {
		return verifier, err
	}
	return data, nil
}

// Status Returns domain status if is webmail or disposable.
func (conf *Tomba) Sources(email string) (models.Source, error) {
	source := models.Source{}
	str, err := conf.TombaCall(SOURCES_PATH, Params{"email": email}, nil, nil)
	if err != nil {
		return source, err
	}
	data, err := models.UnmarshalSource([]byte(str))
	if err != nil {
		return source, err
	}
	return data, nil
}

// Email Format Check
// This endpoint allows you to retrieve the email format patterns used by a specific domain.
// see https://docs.tomba.io/api/finder#email-format
func (conf *Tomba) EmailFormat(domain string) (models.Format, error) {
	format := models.Format{}
	str, err := conf.TombaCall(FORMAT_PATH, Params{"domain": domain}, nil, nil)
	if err != nil {
		return format, err
	}
	data, err := models.UnmarshalFormat([]byte(str))
	if err != nil {
		return format, err
	}
	return data, nil
}

// Employees location count
// This endpoint allows you to retrieve the number of employees in a specific country for a given domain
// see https://docs.tomba.io/api/finder#employees
func (conf *Tomba) EmployeesCount(domain string) (models.Employees, error) {
	employees := models.Employees{}
	str, err := conf.TombaCall(EMPLOYEES_PATH, Params{"domain": domain}, nil, nil)
	if err != nil {
		return employees, err
	}
	data, err := models.UnmarshalEmployees([]byte(str))
	if err != nil {
		return employees, err
	}
	return data, nil
}

// Similar Domains
// This endpoint allows you to retrieve a list of similar domains based on a given domain name.
// see https://docs.tomba.io/api/~endpoints#similar
func (conf *Tomba) SimilarDomains(domain string) (models.Similar, error) {
	similarDomains := models.Similar{}
	str, err := conf.TombaCall(SIMILAR_PATH, Params{"domain": domain}, nil, nil)
	if err != nil {
		return similarDomains, err
	}
	data, err := models.UnmarshalSimilar([]byte(str))
	if err != nil {
		return similarDomains, err
	}
	return data, nil
}

// Technology Check
// This endpoint allows you to retrieve the technologies used by a specific domain.
// see https://docs.tomba.io/api/~endpoints#technology
func (conf *Tomba) TechnologyCheck(domain string) (models.Technology, error) {
	technology := models.Technology{}
	str, err := conf.TombaCall("/technology", Params{"domain": domain}, nil, nil)
	if err != nil {
		return technology, err
	}
	data, err := models.UnmarshalTechnology([]byte(str))
	if err != nil {
		return technology, err
	}
	return data, nil
}

// SearchCompanies searches for companies using natural language queries or structured filters.
// The AI assistant will automatically generate appropriate filters from your query.
// see https://docs.tomba.io/api/reveal#search-companies
func (conf *Tomba) SearchCompanies(request *models.RevealSearchRequest) (models.RevealSearchResponse, error) {
	response := models.RevealSearchResponse{}

	requestParams := make(Params)
	if request != nil {
		if request.Query != "" {
			requestParams["query"] = request.Query
		}
		if request.Page > 0 {
			requestParams["page"] = request.Page
		}
		if request.Filters != nil {
			requestParams["filters"] = request.Filters
		}
	}

	method := "POST"
	str, err := conf.TombaCall(REVEAL_SEARCH_PATH, requestParams, &method, nil)
	if err != nil {
		return response, err
	}
	data, err := models.UnmarshalRevealSearch([]byte(str))
	if err != nil {
		return response, err
	}
	return data, nil
}

// Usage Check your monthly requests.
func (conf *Tomba) Usage() (models.Usage, error) {
	usage := models.Usage{}
	str, err := conf.TombaCall(USAGE_PATH, nil, nil, nil)
	if err != nil {
		return usage, err
	}
	data, err := models.UnmarshalUsage([]byte(str))
	if err != nil {
		return usage, err
	}
	return data, nil
}

// Logs Returns a your last 1,000 requests you made during the last 3 months.
func (conf *Tomba) Logs() (models.Logs, error) {
	logs := models.Logs{}
	str, err := conf.TombaCall(LOGS_PATH, nil, nil, nil)
	if err != nil {
		return logs, err
	}
	data, err := models.UnmarshalLogs([]byte(str))
	if err != nil {
		return logs, err
	}
	return data, nil
}

// Helper function to convert BulkType to path
func getBulkPath(bulkType models.BulkType) string {
	return fmt.Sprintf(BULK_PATH, string(bulkType))
}

// GetAllBulks retrieves all bulk operations for a specific type
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

// GetBulk retrieves a specific bulk operation
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

// CreateBulk creates a new bulk operation
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

// CreateBulkWithFile creates a new bulk operation with file upload
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

// CreateSearchBulk creates a new search bulk operation
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

// CreateFinderBulk creates a new finder bulk operation with file upload
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

// CreatePhoneValidatorBulk creates a new phone validator bulk operation
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

// LaunchBulk launches a bulk operation
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

// DeleteBulk deletes a bulk operation
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

// ArchiveBulk archives a bulk operation
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

// RenameBulk renames a bulk operation
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

// GetBulkProgress gets the progress of a bulk operation
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

// DownloadBulk downloads bulk operation results
func (conf *Tomba) DownloadBulk(bulkType models.BulkType, id int64, params *models.BulkDownloadParams) ([]byte, error) {
	path := fmt.Sprintf(BULK_PATH+"/%d/download", string(bulkType), id)

	requestParams := make(Params)
	if params != nil && params.Type != "" {
		requestParams["type"] = params.Type
	}

	method := "GET"
	return conf.TombaCall(path, requestParams, &method, nil)
}

// Convenience methods for specific bulk types

// GetAllSearchBulks gets all search bulk operations
func (conf *Tomba) GetAllSearchBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeSearch, params)
}

// GetAllSimilarBulks gets all similar bulk operations
func (conf *Tomba) GetAllSimilarBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeSimilar, params)
}

// GetAllCompanyBulks gets all company bulk operations
func (conf *Tomba) GetAllCompanyBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeCompany, params)
}

// GetAllFinderBulks gets all finder bulk operations
func (conf *Tomba) GetAllFinderBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeFinder, params)
}

// GetAllEnrichBulks gets all enrich bulk operations
func (conf *Tomba) GetAllEnrichBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeEnrich, params)
}

// GetAllLinkedInBulks gets all LinkedIn bulk operations
func (conf *Tomba) GetAllLinkedInBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeLinkedIn, params)
}

// GetAllAuthorBulks gets all author bulk operations
func (conf *Tomba) GetAllAuthorBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeAuthor, params)
}

// GetAllVerifierBulks gets all verifier bulk operations
func (conf *Tomba) GetAllVerifierBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypeVerifier, params)
}

// GetAllPhoneFinderBulks gets all phone finder bulk operations
func (conf *Tomba) GetAllPhoneFinderBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypePhoneFinder, params)
}

// GetAllPhoneValidatorBulks gets all phone validator bulk operations
func (conf *Tomba) GetAllPhoneValidatorBulks(params *models.BulkGetParams) (*models.BulkListResponse, error) {
	return conf.GetAllBulks(models.BulkTypePhoneValidator, params)
}

// SaveBulkResults saves bulk results to a file
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
	defer file.Close()

	_, err = file.Write(data)
	return err
}
