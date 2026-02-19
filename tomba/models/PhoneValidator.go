package models

import "encoding/json"

func UnmarshalPhoneValidator(data []byte) (PhoneValidator, error) {
	var r PhoneValidator
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PhoneValidator) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// PhoneValidator represents the phone validator response
type PhoneValidator struct {
	Data PhoneValidatorData `json:"data"`
}

// PhoneValidatorData represents validated phone information
type PhoneValidatorData struct {
	Valid         bool     `json:"valid"`
	LocalFormat   string   `json:"local_format"`
	IntlFormat    string   `json:"intl_format"`
	E164Format    string   `json:"e164_format"`
	RFC3966Format string   `json:"rfc3966_format"`
	CountryCode   string   `json:"country_code"`
	LineType      string   `json:"line_type"`
	Carrier       string   `json:"carrier"`
	Timezones     []string `json:"timezones"`
}
