package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// Usage returns information about your monthly API request usage.
// See https://docs.tomba.io/api/account#retrieve-api-usage#get-usage
func (conf *Tomba) Usage() (models.Usage, error) {
	usage := models.Usage{}
	str, err := conf.TombaCall(USAGE_PATH, nil, nil, nil)
	if err != nil {
		return usage, err
	}
	data, err := models.UnmarshalUsage(str)
	if err != nil {
		return usage, err
	}
	return data, nil
}
