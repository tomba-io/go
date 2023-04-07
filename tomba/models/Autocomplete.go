package models

import "encoding/json"

func UnmarshalAutocomplete(data []byte) (Autocomplete, error) {
	var r Autocomplete
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Autocomplete) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Autocomplete struct {
	Data []AutocompleteData `json:"data"`
	Meta AutocompleteMeta   `json:"meta"`
}

type AutocompleteData struct {
	WebsiteURL string `json:"website_url"`
	Name       string `json:"name"`
	EmailCount int64  `json:"email_count"`
	Logo       string `json:"logo"`
}

type AutocompleteMeta struct {
	Limit interface{} `json:"limit"`
	Query string      `json:"query"`
}
