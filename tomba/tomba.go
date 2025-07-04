package tomba

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tomba-io/go/tomba/models"
)

// TombaCall to Tomba Api
// get data
func (conf *Tomba) TombaCall(path string, params Params, method *string) (string, error) {
	apiUrl := DEFAULT_BASE_URL + path

	if len(params) != 0 {
		queryParams := url.Values{}
		for key, value := range params {
			queryParams.Add(key, value)
		}
		apiUrl += "?" + queryParams.Encode()
	}
	methodStr := "GET"
	if method != nil {
		methodStr = *method
	}
	req, _ := http.NewRequest(methodStr, apiUrl, nil)

	req.Header.Add("X-Tomba-Key", conf.ApiKey)
	req.Header.Add("X-Tomba-Secret", conf.ApiSecret)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Tomba go-client")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New("unexpected response status: " + res.Status + "\nResponse body: " + string(body))
	}
	return string(body), nil
}

// Account Returns information about the current account.
func (conf *Tomba) Account() (models.Account, error) {
	account := models.Account{}
	str, err := conf.TombaCall(ACCOUNT_PATH, nil, nil)
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
	str, err := conf.TombaCall(SEARCH_PATH, params, nil)
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
	str, err := conf.TombaCall(COUNT_PATH, Params{"domain": domain}, nil)
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
	str, err := conf.TombaCall(STATUS_PATH, Params{"domain": domain}, nil)
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
	str, err := conf.TombaCall(FINDER_PATH, params, nil)
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
func (conf *Tomba) Enrichment(email string) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(ENRICHMENT_PATH, Params{"email": email}, nil)
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
	str, err := conf.TombaCall(AUTHOR_PATH, Params{"url": url}, nil)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder([]byte(str))
	if err != nil {
		return finder, err
	}
	return data, nil
}

// LinkedinFinder  This API point generates or retrieves the most likely email address from a Linkedin URL.
func (conf *Tomba) LinkedinFinder(url string) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(LINKEDIN_PATH, Params{"url": url}, nil)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder([]byte(str))
	if err != nil {
		return finder, err
	}
	return data, nil
}

// EmailVerifier Verify the deliverability of an email address.
func (conf *Tomba) EmailVerifier(email string) (models.Verifier, error) {
	verifier := models.Verifier{}
	str, err := conf.TombaCall(VERIFIER_PATH+email, nil, nil)
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
	str, err := conf.TombaCall(SOURCES_PATH, Params{"email": email}, nil)
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
	str, err := conf.TombaCall(FORMAT_PATH, Params{"domain": domain}, nil)
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
	str, err := conf.TombaCall(EMPLOYEES_PATH, Params{"domain": domain}, nil)
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
	str, err := conf.TombaCall(SIMILAR_PATH, Params{"domain": domain}, nil)
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
	str, err := conf.TombaCall("/technology", Params{"domain": domain}, nil)
	if err != nil {
		return technology, err
	}
	data, err := models.UnmarshalTechnology([]byte(str))
	if err != nil {
		return technology, err
	}
	return data, nil
}

// Usage Check your monthly requests.
func (conf *Tomba) Usage() (models.Usage, error) {
	usage := models.Usage{}
	str, err := conf.TombaCall(USAGE_PATH, nil, nil)
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
	str, err := conf.TombaCall(LOGS_PATH, nil, nil)
	if err != nil {
		return logs, err
	}
	data, err := models.UnmarshalLogs([]byte(str))
	if err != nil {
		return logs, err
	}
	return data, nil
}
