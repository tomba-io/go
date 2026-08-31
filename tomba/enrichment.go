package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// PersonFind looks up person data based on an email address.
// See https://docs.tomba.io/api/people#person-find
func (conf *Tomba) PersonFind(params Params) (models.PersonResponse, error) {
	person := models.PersonResponse{}
	str, err := conf.TombaCall(PEOPLE_FIND_PATH, params, nil, nil)
	if err != nil {
		return person, err
	}
	data, err := models.UnmarshalPersonResponse(str)
	if err != nil {
		return person, err
	}
	return data, nil
}

// CompanyFind looks up company data based on a domain name.
// See https://docs.tomba.io/api/companies#company-find
func (conf *Tomba) CompanyFind(params Params) (models.CompanyResponse, error) {
	company := models.CompanyResponse{}
	str, err := conf.TombaCall(COMPANIES_FIND_PATH, params, nil, nil)
	if err != nil {
		return company, err
	}
	data, err := models.UnmarshalCompanyResponse(str)
	if err != nil {
		return company, err
	}
	return data, nil
}

// CombinedFind looks up both person and company data based on an email address.
// See https://docs.tomba.io/api/combined#combined-find
func (conf *Tomba) CombinedFind(params Params) (models.CombinedResponse, error) {
	combined := models.CombinedResponse{}
	str, err := conf.TombaCall(COMBINED_FIND_PATH, params, nil, nil)
	if err != nil {
		return combined, err
	}
	data, err := models.UnmarshalCombinedResponse(str)
	if err != nil {
		return combined, err
	}
	return data, nil
}
