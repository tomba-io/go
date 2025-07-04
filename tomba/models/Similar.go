package models

import "encoding/json"

func UnmarshalSimilar(data []byte) (Similar, error) {
	var r Similar
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Similar) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Similar struct {
	Data []SimilarData `json:"data"`
}

type SimilarData struct {
	WebsiteURL string  `json:"website_url"`
	Name       *string `json:"name"`
	Industries *string `json:"industries"`
}
