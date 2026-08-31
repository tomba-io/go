package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// PhoneFinder searches for phone numbers based on an email, domain, or LinkedIn URL.
// At least one of email, domain, or linkedin must be provided in params.
//
// Supported params keys: email, domain, linkedin, full, webhook_url.
//
// See https://docs.tomba.io/api/phone#phone-finder
func (conf *Tomba) PhoneFinder(params Params) (models.Phone, error) {
	phone := models.Phone{}
	str, err := conf.TombaCall(PHONE_FINDER_PATH, params, nil, nil)
	if err != nil {
		return phone, err
	}
	data, err := models.UnmarshalPhone(str)
	if err != nil {
		return phone, err
	}
	return data, nil
}

// PhoneValidator validates a phone number and retrieves its associated information.
// See https://docs.tomba.io/api/phone#phone-validator
func (conf *Tomba) PhoneValidator(params Params) (models.PhoneValidator, error) {
	phoneValidator := models.PhoneValidator{}
	str, err := conf.TombaCall(PHONE_VALIDATOR_PATH, params, nil, nil)
	if err != nil {
		return phoneValidator, err
	}
	data, err := models.UnmarshalPhoneValidator(str)
	if err != nil {
		return phoneValidator, err
	}
	return data, nil
}
