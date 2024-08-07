package models

import (
	"encoding/json"
	"time"
)

func UnmarshalSource(data []byte) (Source, error) {
	var r Source
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Source) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Source struct {
	Sources []SourceElement `json:"data"`
}

type SourceElement struct {
	URI         string    `json:"uri"`
	ExtractedOn time.Time `json:"extracted_on"`
	LastSeenOn  time.Time `json:"last_seen_on"`
	StillOnPage bool      `json:"still_on_page"`
	WebsiteURL  string    `json:"website_url"`
}
