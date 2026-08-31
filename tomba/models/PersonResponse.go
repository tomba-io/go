package models

import "encoding/json"

// UnmarshalPersonResponse unmarshals a PersonResponse from JSON bytes.
func UnmarshalPersonResponse(data []byte) (PersonResponse, error) {
	var r PersonResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal returns the JSON encoding of PersonResponse.
func (r *PersonResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// PersonResponse represents the response from the /people/find endpoint.
type PersonResponse struct {
	Data PersonData `json:"data"`
}

// PersonData represents person data returned by the People Find API.
type PersonData struct {
	FullName   string `json:"full_name,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Email      string `json:"email,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Phone      any    `json:"phone_number,omitempty"`
	Country    any    `json:"country,omitempty"`
	Position   string `json:"position,omitempty"`
	Department string `json:"department,omitempty"`
	Seniority  string `json:"seniority,omitempty"`
	Twitter    any    `json:"twitter,omitempty"`
	Linkedin   any    `json:"linkedin,omitempty"`
	Company    string `json:"company,omitempty"`
	Score      int64  `json:"score,omitempty"`
}
