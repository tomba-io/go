package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// Sources returns the sources where an email address has been found on the internet.
// See https://docs.tomba.io/api/source#email-sources
func (conf *Tomba) Sources(email string) (models.Source, error) {
	source := models.Source{}
	str, err := conf.TombaCall(SOURCES_PATH, Params{"email": email}, nil, nil)
	if err != nil {
		return source, err
	}
	data, err := models.UnmarshalSource(str)
	if err != nil {
		return source, err
	}
	return data, nil
}
