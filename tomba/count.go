package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// Count returns the total number of email addresses found for a given domain.
// See https://docs.tomba.io/api/finder#email-count#get-email-count
func (conf *Tomba) Count(domain string) (models.Count, error) {
	count := models.Count{}
	str, err := conf.TombaCall(COUNT_PATH, Params{"domain": domain}, nil, nil)
	if err != nil {
		return count, err
	}
	data, err := models.UnmarshalCount(str)
	if err != nil {
		return count, err
	}
	return data, nil
}
