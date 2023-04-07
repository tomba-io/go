package models

import "encoding/json"

func UnmarshalSearch(data []byte) (Search, error) {
	var r Search
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Search) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Search struct {
	Data SearchData `json:"data"`
	Meta SearchMeta `json:"meta"`
}

type SearchData struct {
	Organization SearchOrganization `json:"organization"`
	Emails       []SearchEmail      `json:"emails"`
}

type SearchEmail struct {
	Email        string             `json:"email"`
	FirstName    *string            `json:"first_name"`
	LastName     *string            `json:"last_name"`
	FullName     *string            `json:"full_name"`
	Gender       *string            `json:"gender"`
	PhoneNumber  *bool              `json:"phone_number"`
	Type         string             `json:"type"`
	Country      *string            `json:"country"`
	Position     *string            `json:"position"`
	Department   *string            `json:"department"`
	Seniority    *string            `json:"seniority"`
	Twitter      interface{}        `json:"twitter"`
	Linkedin     *string            `json:"linkedin"`
	AcceptAll    bool               `json:"accept_all"`
	Pattern      *string            `json:"pattern"`
	Score        int64              `json:"score"`
	Verification SearchVerification `json:"verification"`
	LastUpdated  string             `json:"last_updated"`
	Sources      []SearchSource     `json:"sources"`
}

type SearchSource struct {
	URI         string `json:"uri"`
	WebsiteURL  string `json:"website_url"`
	ExtractedOn string `json:"extracted_on"`
	LastSeenOn  string `json:"last_seen_on"`
	StillOnPage bool   `json:"still_on_page"`
}

type SearchVerification struct {
	Date   *string `json:"date"`
	Status *string `json:"status"`
}

type SearchOrganization struct {
	Location      SearchLocation    `json:"location"`
	SocialLinks   SearchSocialLinks `json:"social_links"`
	Disposable    bool              `json:"disposable"`
	Webmail       bool              `json:"webmail"`
	WebsiteURL    *string           `json:"website_url"`
	PhoneNumber   string            `json:"phone_number"`
	Industries    interface{}       `json:"industries"`
	PostalCode    interface{}       `json:"postal_code"`
	EmployeeCount int64             `json:"employee_count"`
	Founded       interface{}       `json:"founded"`
	CompanySize   interface{}       `json:"company_size"`
	LastUpdated   string            `json:"last_updated"`
	Revenue       interface{}       `json:"revenue"`
	AcceptAll     bool              `json:"accept_all"`
	Description   interface{}       `json:"description"`
	Pattern       *string           `json:"pattern"`
	DomainScore   int64             `json:"domain_score"`
	Organization  string            `json:"organization"`
	Whois         SearchWhois       `json:"whois"`
}

type SearchLocation struct {
	Country       interface{} `json:"country"`
	City          interface{} `json:"city"`
	State         interface{} `json:"state"`
	StreetAddress interface{} `json:"street_address"`
}

type SearchSocialLinks struct {
	TwitterURL  interface{} `json:"twitter_url"`
	FacebookURL string      `json:"facebook_url"`
	LinkedinURL string      `json:"linkedin_url"`
}

type SearchWhois struct {
	RegistrarName string `json:"registrar_name"`
	CreatedDate   string `json:"created_date"`
	ReferralURL   string `json:"referral_url"`
}

type SearchMeta struct {
	Total      int64 `json:"total"`
	PageSize   int64 `json:"pageSize"`
	Current    int64 `json:"current"`
	TotalPages int64 `json:"total_pages"`
}
