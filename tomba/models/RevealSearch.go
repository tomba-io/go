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

// RevealSearchRequest represents the request body for Search Companies.
// See https://docs.tomba.io/api/reveal#search-companies
type RevealSearchRequest struct {
	// Natural language query - AI assistant will select appropriate filters.
	// Use this only on the first request, then use filters for subsequent requests.
	Query string `json:"query,omitempty"`
	// Structured search filters to narrow company results.
	Filters *RevealFilters `json:"filters,omitempty"`
	// Page number for pagination (1-1000).
	Page int `json:"page,omitempty"`
	// Optional webhook URL to receive results asynchronously.
	WebhookURL string `json:"webhook_url,omitempty"`
}

// RevealFilters contains structured filters for company search.
type RevealFilters struct {
	Company         *SearchFilter `json:"company,omitempty"`
	LocationCountry *SearchFilter `json:"location_country,omitempty"`
	LocationCity    *SearchFilter `json:"location_city,omitempty"`
	LocationState   *SearchFilter `json:"location_state,omitempty"`
	Industry        *SearchFilter `json:"industry,omitempty"`
	Size            *SearchFilter `json:"size,omitempty"`
	Type            *SearchFilter `json:"type,omitempty"`
	Keywords        *SearchFilter `json:"keywords,omitempty"`
	Founded         *SearchFilter `json:"founded,omitempty"`
	Technologies    *SearchFilter `json:"technologies,omitempty"`
	Similar         *SearchFilter `json:"similar,omitempty"`
	Revenue         *SearchFilter `json:"revenue,omitempty"`
	SIC             *SearchFilter `json:"sic,omitempty"`
	NAICS           *SearchFilter `json:"naics,omitempty"`
}

// SearchFilter represents include/exclude filter options.
type SearchFilter struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

// RevealSearchResponse represents the response from Search Companies.
type RevealSearchResponse struct {
	Success bool             `json:"success"`
	Data    RevealSearchData `json:"data"`
	Meta    *RevealMeta      `json:"meta,omitempty"`
	Message string           `json:"message,omitempty"`
}

// RevealSearchData contains the companies array and pagination info.
type RevealSearchData struct {
	Companies []Company `json:"companies"`
	Total     int       `json:"total"`
	Page      int       `json:"page"`
	Limit     int       `json:"limit"`
	Pages     int       `json:"pages"`
}

// Company represents a company in the search results.
type Company struct {
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	Country       string `json:"country,omitempty"`
	State         string `json:"state,omitempty"`
	City          string `json:"city,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	Industry      string `json:"industry,omitempty"`
	CompanySize   string `json:"company_size,omitempty"`
	Type          string `json:"type,omitempty"`
	Founded       string `json:"founded,omitempty"`
	WebsiteURL    string `json:"website_url,omitempty"`
	TotalEmails   int    `json:"total_emails,omitempty"`
	Revenue       string `json:"revenue,omitempty"`
	PhoneNumber   bool   `json:"phone_number,omitempty"`
	LinkedinURL   string `json:"linkedin_url,omitempty"`
	FacebookURL   string `json:"facebook_url,omitempty"`
	TwitterURL    string `json:"twitter_url,omitempty"`
	InstagramURL  string `json:"instagram_url,omitempty"`
	GithubURL     string `json:"github_url,omitempty"`
	YoutubeURL    string `json:"youtube_url,omitempty"`
	PinterestURL  string `json:"pinterest_url,omitempty"`
	TiktokURL     string `json:"tiktok_url,omitempty"`
	TotalSimilar  int    `json:"total_similar,omitempty"`
}

// RevealMeta contains metadata about the search results.
type RevealMeta struct {
	Total   int                              `json:"total"`
	Page    int                              `json:"page"`
	Limit   int                              `json:"limit"`
	Pages   int                              `json:"pages"`
	Filters map[string]RevealMetaFilterEntry `json:"filters,omitempty"`
}

// RevealMetaFilterEntry represents a single filter's include/exclude values.
type RevealMetaFilterEntry struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}
