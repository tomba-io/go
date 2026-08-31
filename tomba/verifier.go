package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// EmailVerifier verifies the deliverability of an email address.
//
// Supported params keys: email, enrich_mobile, webhook_url.
//
// See https://docs.tomba.io/api/verifier#email-verifier
func (conf *Tomba) EmailVerifier(params Params) (models.Verifier, error) {
	verifier := models.Verifier{}
	str, err := conf.TombaCall(VERIFIER_PATH, params, nil, nil)
	if err != nil {
		return verifier, err
	}
	data, err := models.UnmarshalVerifier(str)
	if err != nil {
		return verifier, err
	}
	return data, nil
}
