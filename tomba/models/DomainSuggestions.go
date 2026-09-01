package models

import "encoding/json"

// UnmarshalDomainSuggestions unmarshals a DomainSuggestions response from JSON bytes.
func UnmarshalDomainSuggestions(data []byte) (DomainSuggestions, error) {
	var r DomainSuggestions
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal returns the JSON encoding of DomainSuggestions.
func (r *DomainSuggestions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// DomainSuggestions represents the response from the /domain-suggestions endpoint.
type DomainSuggestions struct {
	Data []DomainSuggestion    `json:"data"`
	Meta DomainSuggestionsMeta `json:"meta"`
}

// DomainSuggestion represents a single domain suggestion.
type DomainSuggestion struct {
	Name       string `json:"name"`
	Domain     string `json:"domain"`
	EmailCount int    `json:"email_count"`
}

// DomainSuggestionsMeta contains metadata about the suggestions query.
type DomainSuggestionsMeta struct {
	Query      string `json:"query"`
	Limit      int    `json:"limit"`
	TotalFound int    `json:"total_found"`
}
