package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// Status returns a domain's status indicating whether it is webmail or disposable.
// See https://docs.tomba.io/api/domain#domain-status#get-domain-status
func (conf *Tomba) Status(domain string) (models.Status, error) {
	status := models.Status{}
	str, err := conf.TombaCall(STATUS_PATH, Params{"domain": domain}, nil, nil)
	if err != nil {
		return status, err
	}
	data, err := models.UnmarshalStatus(str)
	if err != nil {
		return status, err
	}
	return data, nil
}

// AutoComplete returns suggestions for domain names based on a partial query.
// See https://docs.tomba.io/api/domain-suggestions#get-domain-suggestions
func (conf *Tomba) AutoComplete(query string) (models.DomainSuggestions, error) {
	result := models.DomainSuggestions{}
	str, err := conf.TombaCall(AUTOCOMPLETE_PATH, Params{"query": query}, nil, nil)
	if err != nil {
		return result, err
	}
	data, err := models.UnmarshalDomainSuggestions(str)
	if err != nil {
		return result, err
	}
	return data, nil
}
