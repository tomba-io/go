package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// DomainSearch searches for email addresses based on a website domain.
// You provide a domain name and it returns all the email addresses found on the internet.
//
// Supported params keys: domain, company, page, limit, department, country, enrich_mobile, webhook_url.
//
// See https://docs.tomba.io/api/finder#domain-search#search
func (conf *Tomba) DomainSearch(params Params) (models.Search, error) {

	search := models.Search{}
	str, err := conf.TombaCall(SEARCH_PATH, params, nil, nil)
	if err != nil {
		return search, err
	}
	data, err := models.UnmarshalSearch(str)
	if err != nil {
		return search, err
	}
	return data, nil
}
