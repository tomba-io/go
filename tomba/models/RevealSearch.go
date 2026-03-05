package models

import "encoding/json"

func UnmarshalRevealSearch(data []byte) (RevealSearchResponse, error) {
	var r RevealSearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RevealSearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// RevealSearchRequest represents the request body for Search Companies
type RevealSearchRequest struct {
	Query   string               `json:"query,omitempty"`
	Filters *RevealSearchFilters `json:"filters,omitempty"`
	Page    int                  `json:"page,omitempty"`
}

// RevealSearchFilters contains structured filters for company search
type RevealSearchFilters struct {
	Company *RevealCompanyFilters `json:"company,omitempty"`
}

// RevealCompanyFilters contains include/exclude filter options
type RevealCompanyFilters struct {
	LocationCountry *RevealCircularFilter `json:"location_country,omitempty"`
	LocationCity    *RevealCircularFilter `json:"location_city,omitempty"`
	LocationState   *RevealCircularFilter `json:"location_state,omitempty"`
	Industry        *RevealCircularFilter `json:"industry,omitempty"`
	Size            *RevealCircularFilter `json:"size,omitempty"`
	Type            *RevealCircularFilter `json:"type,omitempty"`
	Keywords        *RevealCircularFilter `json:"keywords,omitempty"`
	Founded         *RevealCircularFilter `json:"founded,omitempty"`
	Technologies    *RevealCircularFilter `json:"technologies,omitempty"`
	Similar         *RevealCircularFilter `json:"similar,omitempty"`
	Revenue         *RevealCircularFilter `json:"revenue,omitempty"`
	SIC             *RevealCircularFilter `json:"sic,omitempty"`
	NAICS           *RevealCircularFilter `json:"naics,omitempty"`
}

// RevealCircularFilter represents include/exclude filter options
type RevealCircularFilter struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

// RevealSearchResponse represents the response from Search Companies
type RevealSearchResponse struct {
	Success bool              `json:"success"`
	Data    RevealSearchData  `json:"data"`
	Meta    *RevealSearchMeta `json:"meta,omitempty"`
	Message string            `json:"message,omitempty"`
}

// RevealSearchData contains the companies array and pagination info
type RevealSearchData struct {
	Companies []RevealCompany `json:"companies"`
	Total     int             `json:"total"`
	Page      int             `json:"page"`
	Pages     int             `json:"pages"`
}

// RevealCompany represents a company in the search results
type RevealCompany struct {
	Name          string  `json:"name,omitempty"`
	Description   string  `json:"description,omitempty"`
	Country       string  `json:"country,omitempty"`
	State         string  `json:"state,omitempty"`
	City          string  `json:"city,omitempty"`
	StreetAddress string  `json:"street_address,omitempty"`
	PostalCode    string  `json:"postal_code,omitempty"`
	Industry      string  `json:"industry,omitempty"`
	CompanySize   string  `json:"company_size,omitempty"`
	Type          string  `json:"type,omitempty"`
	Founded       string  `json:"founded,omitempty"`
	WebsiteURL    string  `json:"website_url,omitempty"`
	TotalEmails   int     `json:"total_emails,omitempty"`
	Revenue       string  `json:"revenue,omitempty"`
	PhoneNumber   bool    `json:"phone_number,omitempty"`
	LinkedinURL   string  `json:"linkedin_url,omitempty"`
	FacebookURL   string  `json:"facebook_url,omitempty"`
	TwitterURL    string  `json:"twitter_url,omitempty"`
	InstagramURL  *string `json:"instagram_url,omitempty"`
	GithubURL     *string `json:"github_url,omitempty"`
	YoutubeURL    *string `json:"youtube_url,omitempty"`
	PinterestURL  *string `json:"pinterest_url,omitempty"`
	TiktokURL     *string `json:"tiktok_url,omitempty"`
	TotalSimilar  int     `json:"total_similar,omitempty"`
}

// RevealSearchMeta contains metadata about the search results
type RevealSearchMeta struct {
	Total   int                      `json:"total"`
	Page    int                      `json:"page"`
	Pages   int                      `json:"pages"`
	Filters *RevealSearchMetaFilters `json:"filters,omitempty"`
}

// RevealSearchMetaFilters contains the applied filters in the response
type RevealSearchMetaFilters struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}
