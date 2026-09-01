package tomba

import (
	"encoding/json"
	"fmt"
)

// Valid flag types
var validFlagTypes = map[string]bool{
	"email": true, "organization": true, "phone": true,
	"author_url": true, "website": true,
}

// Valid reasons per flag type
var validFlagReasons = map[string]map[string]bool{
	"email":        {"hard_bounce": true, "invalid_email": true, "wrong_person": true, "outdated": true, "other": true},
	"organization": {"wrong_company": true, "outdated": true, "other": true},
	"phone":        {"wrong_phone": true, "outdated": true, "other": true},
	"author_url":   {"broken_url": true, "wrong_person": true, "outdated": true, "other": true},
	"website":      {"broken_url": true, "wrong_company": true, "outdated": true, "other": true},
}

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

// CreateFlag creates a new flag to report incorrect data.
//
// Required params:
//   - flag_type: email, organization, phone, author_url, website
//   - value: the flagged item (email, domain, phone, URL) — 2-255 chars
//   - reason: depends on flag_type (see validFlagReasons)
//
// Optional params:
//   - comment: additional details (max 1000 chars)
//
// See https://docs.tomba.io/api/flag#create-flag
func (conf *Tomba) CreateFlag(params Params) (json.RawMessage, error) {
	// Validate flag_type
	flagType, _ := params["flag_type"].(string)
	if flagType == "" {
		return nil, fmt.Errorf("flag_type is required (email, organization, phone, author_url, website)")
	}
	if !validFlagTypes[flagType] {
		return nil, fmt.Errorf("invalid flag_type %q: must be one of email, organization, phone, author_url, website", flagType)
	}

	// Validate value
	value, _ := params["value"].(string)
	if value == "" {
		return nil, fmt.Errorf("value is required (the item to flag)")
	}
	if len(value) < 2 || len(value) > 255 {
		return nil, fmt.Errorf("value must be 2-255 characters, got %d", len(value))
	}

	// Validate reason
	reason, _ := params["reason"].(string)
	if reason == "" {
		return nil, fmt.Errorf("reason is required")
	}
	if reasons, ok := validFlagReasons[flagType]; ok {
		if !reasons[reason] {
			validList := make([]string, 0, len(reasons))
			for k := range reasons {
				validList = append(validList, k)
			}
			return nil, fmt.Errorf("invalid reason %q for flag_type %q: must be one of %v", reason, flagType, validList)
		}
	}

	// Validate comment length
	if comment, ok := params["comment"].(string); ok && len(comment) > 1000 {
		return nil, fmt.Errorf("comment must be at most 1000 characters, got %d", len(comment))
	}

	method := "POST"
	str, err := conf.TombaCall(FLAG_PATH, params, &method, nil)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(str), nil
}
