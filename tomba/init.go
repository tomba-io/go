package tomba

import (
	"net/http"
	"strconv"
)

// RateLimit holds parsed rate limit headers from an API response.
type RateLimit struct {
	SecondLimit     int    `json:"second_limit"`
	MinuteLimit     int    `json:"minute_limit"`
	DailyLimit      int    `json:"daily_limit"`
	MinuteRemaining int    `json:"minute_remaining"`
	DailyRemaining  int    `json:"daily_remaining"`
	MinuteReset     int    `json:"minute_reset"`
	DailyReset      int    `json:"daily_reset"`
	RetryAfter      int    `json:"retry_after"`
	Policy          string `json:"policy"`
	RateLimit       string `json:"rate_limit"`
}

// ParseRateLimit extracts rate limit information from HTTP response headers.
func ParseRateLimit(resp *http.Response) RateLimit {
	getInt := func(key string) int {
		val, err := strconv.Atoi(resp.Header.Get(key))
		if err != nil {
			return 0
		}
		return val
	}

	return RateLimit{
		SecondLimit:     getInt("x-second-rate-limit"),
		MinuteLimit:     getInt("x-minute-rate-limit"),
		DailyLimit:      getInt("x-daily-rate-limit"),
		MinuteRemaining: getInt("x-minute-request-left"),
		DailyRemaining:  getInt("x-daily-request-left"),
		MinuteReset:     getInt("x-minute-reset-seconds"),
		DailyReset:      getInt("x-daily-reset-seconds"),
		RetryAfter:      getInt("Retry-After"),
		Policy:          resp.Header.Get("RateLimit-Policy"),
		RateLimit:       resp.Header.Get("RateLimit"),
	}
}

// Tomba holds the API credentials used to authenticate with the Tomba API.
type Tomba struct {
	ApiKey        string
	ApiSecret     string
	LastRateLimit RateLimit
	// BaseURL overrides DEFAULT_BASE_URL (e.g. a staging API); see WithBaseURL.
	BaseURL string
}

// WithBaseURL points the client at another API base URL
// (default "https://api.tomba.io/v1").
func (conf *Tomba) WithBaseURL(baseURL string) *Tomba {
	for len(baseURL) > 0 && baseURL[len(baseURL)-1] == '/' {
		baseURL = baseURL[:len(baseURL)-1]
	}
	conf.BaseURL = baseURL
	return conf
}

func (conf *Tomba) baseURL() string {
	if conf.BaseURL != "" {
		return conf.BaseURL
	}
	return DEFAULT_BASE_URL
}

// New creates a new Tomba client with the given API key and secret.
// See https://docs.tomba.io/api/authentication
func New(key string, secret string) *Tomba {
	p := new(Tomba)
	p.ApiKey = key
	p.ApiSecret = secret
	return p
}
