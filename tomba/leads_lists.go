package tomba

import (
	"encoding/json"
)

// ListLeadsLists retrieves all leads lists for the current account.
// See https://docs.tomba.io/api/leads-lists#list-leads-lists
func (conf *Tomba) ListLeadsLists(params Params) (json.RawMessage, error) {
	str, err := conf.TombaCall(LEADS_LISTS_PATH, params, nil, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// GetLeadsList retrieves a specific leads list by its ID.
// See https://docs.tomba.io/api/leads-lists#get-leads-list
func (conf *Tomba) GetLeadsList(id string) (json.RawMessage, error) {
	str, err := conf.TombaCall(LEADS_LISTS_PATH+"/"+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// CreateLeadsList creates a new leads list.
// See https://docs.tomba.io/api/leads-lists#create-leads-list
func (conf *Tomba) CreateLeadsList(params Params) (json.RawMessage, error) {
	method := "POST"
	str, err := conf.TombaCall(LEADS_LISTS_PATH, params, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// UpdateLeadsList updates an existing leads list by its ID.
// See https://docs.tomba.io/api/leads-lists#update-leads-list
func (conf *Tomba) UpdateLeadsList(id string, params Params) (json.RawMessage, error) {
	method := "PUT"
	str, err := conf.TombaCall(LEADS_LISTS_PATH+"/"+id, params, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// DeleteLeadsList deletes a leads list by its ID.
// See https://docs.tomba.io/api/leads-lists#delete-leads-list
func (conf *Tomba) DeleteLeadsList(id string) (json.RawMessage, error) {
	method := "DELETE"
	str, err := conf.TombaCall(LEADS_LISTS_PATH+"/"+id, nil, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}
