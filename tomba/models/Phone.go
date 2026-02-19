package models

import "encoding/json"

func UnmarshalPhone(data []byte) (Phone, error) {
	var r Phone
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Phone) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Phone represents the phone finder response
// Data can be either a single PhoneData object or an array of PhoneData objects
// depending on the 'full' parameter in the request
type Phone struct {
	Data json.RawMessage `json:"data"`
}

// PhoneData represents phone information for an email/domain/linkedin
type PhoneData struct {
	Email         string   `json:"email,omitempty"`
	Domain        string   `json:"domain,omitempty"`
	LinkedIn      string   `json:"linkedin,omitempty"`
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

// GetSingleData returns the data as a single PhoneData object
// Use this when 'full' parameter was not set or set to false
func (p *Phone) GetSingleData() (PhoneData, error) {
	var data PhoneData
	err := json.Unmarshal(p.Data, &data)
	return data, err
}

// GetFullData returns the data as an array of PhoneData objects
// Use this when 'full' parameter was set to true
func (p *Phone) GetFullData() ([]PhoneData, error) {
	var data []PhoneData
	err := json.Unmarshal(p.Data, &data)
	return data, err
}
