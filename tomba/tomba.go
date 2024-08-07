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
func (conf *Tomba) TombaCall(path string, params Params) (string, error) {

	apiUrl := DEFAULT_BASE_URL + path

	if len(params) != 0 {
		queryParams := url.Values{}
		for key, value := range params {
			queryParams.Add(key, value)
		}
		apiUrl += "?" + queryParams.Encode()
	}

	req, _ := http.NewRequest("GET", apiUrl, nil)

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
	str, err := conf.TombaCall(ACCOUNT_PATH, nil)
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
	str, err := conf.TombaCall(SEARCH_PATH, params)
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
	str, err := conf.TombaCall(COUNT_PATH, Params{"domain": domain})
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
	str, err := conf.TombaCall(STATUS_PATH, Params{"domain": domain})
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
	str, err := conf.TombaCall(FINDER_PATH, params)
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
	str, err := conf.TombaCall(ENRICHMENT_PATH, Params{"email": email})
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
	str, err := conf.TombaCall(AUTHOR_PATH, Params{"url": url})
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
	str, err := conf.TombaCall(LINKEDIN_PATH, Params{"url": url})
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
	str, err := conf.TombaCall(VERIFIER_PATH+email, nil)
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
	str, err := conf.TombaCall(SOURCES_PATH, Params{"email": email})
	if err != nil {
		return source, err
	}
	data, err := models.UnmarshalSource([]byte(str))
	if err != nil {
		return source, err
	}
	return data, nil
}

// Usage Check your monthly requests.
func (conf *Tomba) Usage() (models.Usage, error) {
	usage := models.Usage{}
	str, err := conf.TombaCall(USAGE_PATH, nil)
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
	str, err := conf.TombaCall(LOGS_PATH, nil)
	if err != nil {
		return logs, err
	}
	data, err := models.UnmarshalLogs([]byte(str))
	if err != nil {
		return logs, err
	}
	return data, nil
}
