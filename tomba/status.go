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
// See https://docs.tomba.io/api/domain#domain-status#autocomplete
func (conf *Tomba) AutoComplete(query string) (models.Status, error) {
	status := models.Status{}
	str, err := conf.TombaCall(AUTOCOMPLETE_PATH, Params{"query": query}, nil, nil)
	if err != nil {
		return status, err
	}
	data, err := models.UnmarshalStatus(str)
	if err != nil {
		return status, err
	}
	return data, nil
}
