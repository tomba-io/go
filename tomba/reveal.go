package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// SearchCompanies searches for companies using natural language queries or structured filters.
// See https://docs.tomba.io/api/reveal#search-companies
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
	data, err := models.UnmarshalRevealSearch(str)
	if err != nil {
		return response, err
	}
	return data, nil
}
