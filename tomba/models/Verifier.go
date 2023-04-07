package models

import "encoding/json"

func UnmarshalVerifier(data []byte) (Verifier, error) {
	var r Verifier
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Verifier) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Verifier struct {
	Data VerifierData `json:"data"`
}

type VerifierData struct {
	Email   VerifierEmail `json:"email"`
	Sources []interface{} `json:"sources"`
}

type VerifierEmail struct {
	MXRecords  bool          `json:"mx_records"`
	SMTPServer bool          `json:"smtp_server"`
	SMTPCheck  bool          `json:"smtp_check"`
	AcceptAll  bool          `json:"accept_all"`
	Block      bool          `json:"block"`
	Email      string        `json:"email"`
	Gibberish  bool          `json:"gibberish"`
	Disposable bool          `json:"disposable"`
	Webmail    bool          `json:"webmail"`
	Result     string        `json:"result"`
	Score      int64         `json:"score"`
	Regex      bool          `json:"regex"`
	Whois      VerifierWhois `json:"whois"`
	Status     string        `json:"status"`
}

type VerifierWhois struct {
	RegistrarName string `json:"registrar_name"`
	CreatedDate   string `json:"created_date"`
	ReferralURL   string `json:"referral_url"`
}
