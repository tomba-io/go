package tomba

import (
	"encoding/json"
)

// ListLeads retrieves all leads for the current account.
//
// Supported params keys: page, limit, domain.
//
// See https://docs.tomba.io/api/leads
func (conf *Tomba) ListLeads(params Params) (json.RawMessage, error) {
	str, err := conf.TombaCall(LEADS_PATH, params, nil, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// GetLead retrieves a specific lead by its ID.
// See https://docs.tomba.io/api/leads#retrieve-a-single-lead
func (conf *Tomba) GetLead(id string) (json.RawMessage, error) {
	str, err := conf.TombaCall(LEADS_PATH+"/"+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// CreateLead creates a new lead.
// See https://docs.tomba.io/api/leads#create-a-lead
func (conf *Tomba) CreateLead(params Params) (json.RawMessage, error) {
	method := "POST"
	str, err := conf.TombaCall(LEADS_PATH, params, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// UpdateLead updates an existing lead by its ID.
// See https://docs.tomba.io/api/leads#update-a-lead
func (conf *Tomba) UpdateLead(id string, params Params) (json.RawMessage, error) {
	method := "PUT"
	str, err := conf.TombaCall(LEADS_PATH+"/"+id, params, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// DeleteLead deletes a lead by its ID.
// See https://docs.tomba.io/api/leads#delete-a-lead
func (conf *Tomba) DeleteLead(id string) (json.RawMessage, error) {
	method := "DELETE"
	str, err := conf.TombaCall(LEADS_PATH+"/"+id, nil, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}
