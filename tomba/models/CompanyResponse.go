package models

import "encoding/json"

// UnmarshalCompanyResponse unmarshals a CompanyResponse from JSON bytes.
func UnmarshalCompanyResponse(data []byte) (CompanyResponse, error) {
	var r CompanyResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal returns the JSON encoding of CompanyResponse.
func (r *CompanyResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// CompanyResponse represents the response from the /companies/find endpoint.
type CompanyResponse struct {
	Data CompanyData `json:"data"`
}

// CompanyData represents company data returned by the Companies Find API.
type CompanyData struct {
	Organization CompanyOrganization `json:"organization,omitempty"`
}

// CompanyOrganization represents the organization details.
type CompanyOrganization struct {
	Organization  string            `json:"organization,omitempty"`
	WebsiteURL    *string           `json:"website_url,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Industries    *string           `json:"industries,omitempty"`
	EmployeeCount int64             `json:"employee_count,omitempty"`
	Founded       *string           `json:"founded,omitempty"`
	CompanySize   *string           `json:"company_size,omitempty"`
	Revenue       *string           `json:"revenue,omitempty"`
	Location      SearchLocation    `json:"location"`
	SocialLinks   SearchSocialLinks `json:"social_links"`
	Whois         SearchWhois       `json:"whois"`
	PhoneNumber   bool              `json:"phone_number,omitempty"`
	PhoneData     []PhoneData       `json:"phone_data,omitempty"`
	Disposable    bool              `json:"disposable,omitempty"`
	Webmail       bool              `json:"webmail,omitempty"`
	AcceptAll     bool              `json:"accept_all,omitempty"`
	Pattern       *string           `json:"pattern,omitempty"`
}
