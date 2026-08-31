package tomba

import (
	"encoding/json"
)

// ListFlags retrieves all flags for the current account.
//
// Supported params keys: page, limit.
//
// See https://docs.tomba.io/api/flag#list-flags
func (conf *Tomba) ListFlags(params Params) (json.RawMessage, error) {
	str, err := conf.TombaCall(FLAG_PATH, params, nil, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}

// CreateFlag creates a new flag.
// See https://docs.tomba.io/api/flag#create-flag
func (conf *Tomba) CreateFlag(params Params) (json.RawMessage, error) {
	method := "POST"
	str, err := conf.TombaCall(FLAG_PATH, params, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}
