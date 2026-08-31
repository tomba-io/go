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
	Name          string  `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
	WebsiteURL    *string `json:"website_url,omitempty"`
	Industries    *string `json:"industries,omitempty"`
	Country       *string `json:"country,omitempty"`
	City          *string `json:"city,omitempty"`
	State         *string `json:"state,omitempty"`
	StreetAddress *string `json:"street_address,omitempty"`
	PostalCode    *string `json:"postal_code,omitempty"`
	EmployeeCount int64   `json:"employee_count,omitempty"`
	Founded       *string `json:"founded,omitempty"`
	CompanySize   *string `json:"company_size,omitempty"`
	Revenue       *string `json:"revenue,omitempty"`
	LinkedinURL   *string `json:"linkedin_url,omitempty"`
	TwitterURL    *string `json:"twitter_url,omitempty"`
	FacebookURL   *string `json:"facebook_url,omitempty"`
}
