package tomba

import (
	"encoding/json"
)

// ListKeys retrieves all API keys for the current account.
// See https://docs.tomba.io/api/keys
func (conf *Tomba) ListKeys() (json.RawMessage, error) {
	str, err := conf.TombaCall(KEYS_PATH, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// GetKey retrieves a specific API key by its ID.
// See https://docs.tomba.io/api/keys#get-key
func (conf *Tomba) GetKey(id string) (json.RawMessage, error) {
	str, err := conf.TombaCall(KEYS_PATH+"/"+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// CreateKey creates a new API key.
// See https://docs.tomba.io/api/keys#create-an-api-key
func (conf *Tomba) CreateKey(params Params) (json.RawMessage, error) {
	method := "POST"
	str, err := conf.TombaCall(KEYS_PATH, params, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// DeleteKey deletes an API key by its ID.
// See https://docs.tomba.io/api/keys#delete-an-api-key
func (conf *Tomba) DeleteKey(id string) (json.RawMessage, error) {
	method := "DELETE"
	str, err := conf.TombaCall(KEYS_PATH+"/"+id, nil, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// ResetKey resets an API key by its ID.
// See https://docs.tomba.io/api/keys#reset-an-api-key
func (conf *Tomba) ResetKey(id string) (json.RawMessage, error) {
	method := "PUT"
	str, err := conf.TombaCall(KEYS_PATH+"/"+id+"/reset", nil, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}
