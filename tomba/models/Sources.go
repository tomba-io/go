package models

type Source struct {
	Sources []SourceElement `json:"sources"`
}

type SourceElement struct {
	URI         string `json:"uri"`
	WebsiteURL  string `json:"website_url"`
	ExtractedOn string `json:"extracted_on"`
	LastSeenOn  string `json:"last_seen_on"`
	StillOnPage bool   `json:"still_on_page"`
}
