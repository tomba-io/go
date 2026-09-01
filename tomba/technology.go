package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// TechnologyCheck retrieves the technologies used by a specific domain.
// See https://docs.tomba.io/api/domain#technology
func (conf *Tomba) TechnologyCheck(domain string) (models.Technology, error) {
	technology := models.Technology{}
	str, err := conf.TombaCall("/technology", Params{"domain": domain}, nil, nil)
	if err != nil {
		return technology, err
	}
	data, err := models.UnmarshalTechnology(str)
	if err != nil {
		return technology, err
	}
	return data, nil
}
