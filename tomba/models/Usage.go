package models

import "encoding/json"

func UnmarshalUsage(data []byte) (Usage, error) {
	var r Usage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Usage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Usage struct {
	Data  []UsageData `json:"data"`
	Total UsageTotal  `json:"total"`
}

type UsageData struct {
	ID              int64  `json:"id"`
	UserID          int64  `json:"user_id"`
	Search          int64  `json:"search"`
	Verifier        int64  `json:"verifier"`
	Export          int64  `json:"export"`
	Sources         int64  `json:"sources"`
	EmailCount      int64  `json:"email_count"`
	SourceWebsite   int64  `json:"source_website"`
	SourceBulk      int64  `json:"source_bulk"`
	SourceExtension int64  `json:"source_extension"`
	SourceAPI       int64  `json:"source_api"`
	SourceSheets    int64  `json:"source_sheets"`
	CreatedAt       string `json:"created_at"`
}

type UsageTotal struct {
	Search          int64 `json:"search"`
	Verifier        int64 `json:"verifier"`
	Export          int64 `json:"export"`
	Sources         int64 `json:"sources"`
	SourceWebsite   int64 `json:"source_website"`
	SourceExtension int64 `json:"source_extension"`
	SourceBulk      int64 `json:"source_bulk"`
	SourceAPI       int64 `json:"source_api"`
	SourceSheets    int64 `json:"source_sheets"`
}
