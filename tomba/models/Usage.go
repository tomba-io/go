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
	Data  []UsageTotal `json:"data"`
	Total UsageTotal   `json:"total"`
}

type UsageTotal struct {
	Usage        int64   `json:"usage"`
	CreatedAt    *string `json:"created_at,omitempty"`
	Domain       int64   `json:"domain"`
	Finder       int64   `json:"finder"`
	Verifier     int64   `json:"verifier"`
	Technologies int64   `json:"technologies"`
	Website      int64   `json:"website"`
	Bulk         int64   `json:"bulk"`
	Extension    int64   `json:"extension"`
	API          int64   `json:"api"`
	Mobile       int64   `json:"mobile"`
	Desktop      int64   `json:"desktop"`
	Sheets       int64   `json:"sheets"`
}
