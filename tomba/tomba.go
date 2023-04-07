package tomba

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/tomba-io/go/tomba/models"
)

// TombaCall to Tomba Api
// get data
func (conf *Tomba) TombaCall(path string) (string, error) {

	apiUrl := DEFAULT_BASE_URL + path

	req, _ := http.NewRequest("GET", apiUrl, nil)

	req.Header.Add("X-Tomba-Key", conf.ApiKey)
	req.Header.Add("X-Tomba-Secret", conf.ApiSecret)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Tomba go-client")

	res, _ := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return "", errors.New("\n  response response =>:\n" + string(body))
	}
	return string(body), err
}

// Account Returns information about the current account.
func (conf *Tomba) Account() (models.Account, error) {
	account := models.Account{}
	str, err := conf.TombaCall(ACCOUNT_PATH)
	if err != nil {
		return account, err
	}
	data, err := models.UnmarshalAccount([]byte(str))
	if err != nil {
		return account, err
	}
	return data, nil
}

//  Search emails are based on the website You give one domain name and it returns all the email addresses found on the internet.
func (conf *Tomba) DomainSearch(domain string) (models.Search, error) {
	search := models.Search{}
	str, err := conf.TombaCall(SEARCH_PATH + "?domain=" + domain)
	if err != nil {
		return search, err
	}
	data, err := models.UnmarshalSearch([]byte(str))
	if err != nil {
		return search, err
	}
	return data, nil
}

//  Returns total email addresses we have for one domain.
func (conf *Tomba) Count(domain string) (models.Count, error) {
	count := models.Count{}
	str, err := conf.TombaCall(COUNT_PATH + "?domain=" + domain)
	if err != nil {
		return count, err
	}
	data, err := models.UnmarshalCount([]byte(str))
	if err != nil {
		return count, err
	}
	return data, nil
}

//  Returns domain status if is webmail or disposable.
func (conf *Tomba) Status(domain string) (models.Status, error) {
	status := models.Status{}
	str, err := conf.TombaCall(STATUS_PATH + "?domain=" + domain)
	if err != nil {
		return status, err
	}
	data, err := models.UnmarshalStatus([]byte(str))
	if err != nil {
		return status, err
	}
	return data, nil
}

//  Company Autocomplete is an API that lets you auto-complete company names and retrieve logo and domain information.
func (conf *Tomba) Autocomplete(search string) (models.Autocomplete, error) {
	auto := models.Autocomplete{}
	str, err := conf.TombaCall(AUTOCOMPLETE_PATH + "?query=" + search)
	if err != nil {
		return auto, err
	}
	data, err := models.UnmarshalAutocomplete([]byte(str))
	if err != nil {
		return auto, err
	}
	return data, nil
}

//  Generates or retrieves the most likely email address from a domain name, a first name and a last name.
func (conf *Tomba) EmailFinder(domain string, fname string, lname string) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(FINDER_PATH + "?domain=" + domain + "&first_name=" + fname + "&last_name=" + lname)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder([]byte(str))
	if err != nil {
		return finder, err
	}
	return data, nil
}

//  The Enrichment API lets you look up person and company data based on an email, For example, you could retrieve a person’s name, location and social handles from an email
func (conf *Tomba) Enrichment(email string) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(ENRICHMENT_PATH + "?email=" + email)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder([]byte(str))
	if err != nil {
		return finder, err
	}
	return data, nil
}

//  This API }point generates or retrieves the most likely email address from a blog post url.
func (conf *Tomba) AuthorFinder(url string) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(AUTHOR_PATH + "?url=" + url)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder([]byte(str))
	if err != nil {
		return finder, err
	}
	return data, nil
}

//  This API }point generates or retrieves the most likely email address from a Linkedin URL.
func (conf *Tomba) LinkedinFinder(url string) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(LINKEDIN_PATH + "?url=" + url)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder([]byte(str))
	if err != nil {
		return finder, err
	}
	return data, nil
}

//  Verify the deliverability of an email address.
func (conf *Tomba) EmailVerifier(email string) (models.Verifier, error) {
	verifier := models.Verifier{}
	str, err := conf.TombaCall(VERIFIER_PATH + email)
	if err != nil {
		return verifier, err
	}
	data, err := models.UnmarshalVerifier([]byte(str))
	if err != nil {
		return verifier, err
	}
	return data, nil
}

//  Check your monthly requests.
func (conf *Tomba) Usage() (models.Usage, error) {
	usage := models.Usage{}
	str, err := conf.TombaCall(USAGE_PATH)
	if err != nil {
		return usage, err
	}
	data, err := models.UnmarshalUsage([]byte(str))
	if err != nil {
		return usage, err
	}
	return data, nil
}

//  Returns a your last 1,000 requests you made during the last 3 months.
func (conf *Tomba) Logs() (models.Logs, error) {
	logs := models.Logs{}
	str, err := conf.TombaCall(LOGS_PATH)
	if err != nil {
		return logs, err
	}
	data, err := models.UnmarshalLogs([]byte(str))
	if err != nil {
		return logs, err
	}
	return data, nil
}
