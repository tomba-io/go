package tomba

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// APIError is a non-2xx answer from the Tomba API. Callers can inspect it with
// errors.As:
//
//	var apiErr *tomba.APIError
//	if errors.As(err, &apiErr) && apiErr.HTTPStatus == http.StatusTooManyRequests { ... }
type APIError struct {
	HTTPStatus int    // HTTP status code
	Type       string // errors.type, e.g. "params_invalid", "rate_limit", "unknown_record"
	Message    string // errors.message, human readable
	Code       int    // errors.code
	RetryAfter int    // seconds, from the Retry-After header (0 if absent)
	Body       []byte // raw response body
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("tomba: %s (HTTP %d, %s)", e.Message, e.HTTPStatus, e.Type)
	}
	return fmt.Sprintf("tomba: HTTP %d %s", e.HTTPStatus, http.StatusText(e.HTTPStatus))
}

// newAPIError parses the {"errors": {"type", "message", "code"}} envelope.
func newAPIError(resp *http.Response, body []byte) *APIError {
	e := &APIError{HTTPStatus: resp.StatusCode, Body: body}
	e.RetryAfter, _ = strconv.Atoi(resp.Header.Get("Retry-After"))
	var envelope struct {
		Errors struct {
			Type    string `json:"type"`
			Message string `json:"message"`
			Code    int    `json:"code"`
		} `json:"errors"`
		Error string `json:"error"`
	}
	if json.Unmarshal(body, &envelope) == nil {
		e.Type, e.Message, e.Code = envelope.Errors.Type, envelope.Errors.Message, envelope.Errors.Code
		if e.Message == "" {
			e.Message = envelope.Error
		}
	}
	return e
}
