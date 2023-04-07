package models

import "encoding/json"

func UnmarshalAccount(data []byte) (Account, error) {
	var r Account
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Account) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Account struct {
	Data AccountData `json:"data"`
}

type AccountData struct {
	Available     int64           `json:"available"`
	UserID        int64           `json:"user_id"`
	SecretToken   string          `json:"secret_token"`
	Role          int64           `json:"role"`
	Confirmed     bool            `json:"confirmed"`
	Blocked       bool            `json:"blocked"`
	FirstName     string          `json:"first_name"`
	LastName      string          `json:"last_name"`
	Email         string          `json:"email"`
	Phone         string          `json:"phone"`
	Image         string          `json:"image"`
	Deleted       bool            `json:"deleted"`
	Provider      string          `json:"provider"`
	Timezone      string          `json:"timezone"`
	Newletter     bool            `json:"newletter"`
	Activity      bool            `json:"activity"`
	Title         string          `json:"title"`
	Country       string          `json:"country"`
	CreatedAt     string          `json:"created_at"`
	IP            string          `json:"ip"`
	PaymentStatus bool            `json:"payment_status"`
	Issued        string          `json:"Issued"`
	Expired       string          `json:"expired"`
	Pricing       AccountPricing  `json:"pricing"`
	Time          AccountTime     `json:"time"`
	Used          int64           `json:"used"`
	Teams         AccountTeams    `json:"teams"`
	Requests      AccountRequests `json:"requests"`
}

type AccountPricing struct {
	Name                   string      `json:"name"`
	PricingID              int64       `json:"pricing_id"`
	AvailableSearches      int64       `json:"available_searches"`
	AvailableVerifications int64       `json:"available_verifications"`
	AvailablePhones        int64       `json:"available_phones"`
	AvailableLeads         int64       `json:"available_leads"`
	AvailableList          int64       `json:"available_list"`
	AvailableAttr          int64       `json:"available_attr"`
	AvailableKeys          int64       `json:"available_keys"`
	AvailableTeams         int64       `json:"available_teams"`
	AvailableTechnologies  int64       `json:"available_technologies"`
	AvailableB2B           int64       `json:"available_b2b"`
	AvailableSources       int64       `json:"available_sources"`
	AvailableEmailCount    int64       `json:"available_email_count"`
	Frequency              string      `json:"frequency"`
	PriceMonthly           string      `json:"price_monthly"`
	PriceYearly            string      `json:"price_yearly"`
	UpdateURL              interface{} `json:"update_url"`
	CancelURL              interface{} `json:"cancel_url"`
}

type AccountRequests struct {
	Domains       AccountB2B `json:"domains"`
	Verifications AccountB2B `json:"verifications"`
	Phones        AccountB2B `json:"phones"`
	B2B           AccountB2B `json:"b2b"`
}

type AccountB2B struct {
	Available int64 `json:"available"`
	Used      int64 `json:"used"`
}

type AccountTeams struct {
	TeamID    int64        `json:"team_id"`
	Role      string       `json:"role"`
	Workspace bool         `json:"workspace"`
	Export    bool         `json:"export"`
	Bulks     bool         `json:"bulks"`
	Available int64        `json:"available"`
	Limit     AccountLimit `json:"limit"`
	Owner     AccountOwner `json:"owner"`
}

type AccountLimit struct {
	Domain   interface{} `json:"domain"`
	Finder   interface{} `json:"finder"`
	Verifier interface{} `json:"verifier"`
}

type AccountOwner struct {
	Name                   interface{} `json:"name"`
	AvailableTeams         interface{} `json:"available_teams"`
	AvailableSearches      interface{} `json:"available_searches"`
	AvailableVerifications interface{} `json:"available_verifications"`
	Usage                  interface{} `json:"usage"`
}

type AccountTime struct {
	Date         string `json:"date"`
	TimezoneType int64  `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}
