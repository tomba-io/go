package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// EmailFormat retrieves the email format patterns used by a specific domain.
// See https://docs.tomba.io/api/finder#email-format
func (conf *Tomba) EmailFormat(domain string) (models.Format, error) {
	format := models.Format{}
	str, err := conf.TombaCall(FORMAT_PATH, Params{"domain": domain}, nil, nil)
	if err != nil {
		return format, err
	}
	data, err := models.UnmarshalFormat(str)
	if err != nil {
		return format, err
	}
	return data, nil
}
