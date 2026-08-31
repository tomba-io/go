package models

import "encoding/json"

// UnmarshalCombinedResponse unmarshals a CombinedResponse from JSON bytes.
func UnmarshalCombinedResponse(data []byte) (CombinedResponse, error) {
	var r CombinedResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal returns the JSON encoding of CombinedResponse.
func (r *CombinedResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// CombinedResponse represents the response from the /combined/find endpoint.
type CombinedResponse struct {
	Data CombinedData `json:"data"`
}

// CombinedData represents the combined person and company data.
type CombinedData struct {
	Person  PersonData  `json:"person,omitempty"`
	Company CompanyData `json:"company,omitempty"`
}
