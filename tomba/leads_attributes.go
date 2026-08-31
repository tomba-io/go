package tomba

import (
	"encoding/json"
)

// ListAttributes retrieves all custom attributes.
// See https://docs.tomba.io/api/attributes#list-attributes
func (conf *Tomba) ListAttributes(params Params) (json.RawMessage, error) {
	str, err := conf.TombaCall(ATTRIBUTES_PATH, params, nil, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// GetAttribute retrieves a specific attribute by its ID.
// See https://docs.tomba.io/api/attributes#get-attribute
func (conf *Tomba) GetAttribute(id string) (json.RawMessage, error) {
	str, err := conf.TombaCall(ATTRIBUTES_PATH+"/"+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// CreateAttribute creates a new custom attribute.
// See https://docs.tomba.io/api/attributes#create-attribute
func (conf *Tomba) CreateAttribute(params Params) (json.RawMessage, error) {
	method := "POST"
	str, err := conf.TombaCall(ATTRIBUTES_PATH, params, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// UpdateAttribute updates an existing attribute by its ID.
// See https://docs.tomba.io/api/attributes#update-attribute
func (conf *Tomba) UpdateAttribute(id string, params Params) (json.RawMessage, error) {
	method := "PUT"
	str, err := conf.TombaCall(ATTRIBUTES_PATH+"/"+id, params, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// DeleteAttribute deletes a custom attribute by its ID.
// See https://docs.tomba.io/api/attributes#delete-attribute
func (conf *Tomba) DeleteAttribute(id string) (json.RawMessage, error) {
	method := "DELETE"
	str, err := conf.TombaCall(ATTRIBUTES_PATH+"/"+id, nil, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}
