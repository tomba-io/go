package models

import "encoding/json"

func UnmarshalFinder(data []byte) (Finder, error) {
	var r Finder
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Finder) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Finder struct {
	Data FinderData `json:"data"`
}

type FinderData struct {
	WebsiteURL   string             `json:"website_url"`
	AcceptAll    bool               `json:"accept_all"`
	Email        string             `json:"email"`
	FirstName    string             `json:"first_name"`
	LastName     string             `json:"last_name"`
	Country      string             `json:"country"`
	Gender       string             `json:"gender"`
	PhoneNumber  bool               `json:"phone_number"`
	Position     string             `json:"position"`
	Twitter      interface{}        `json:"twitter"`
	Linkedin     string             `json:"linkedin"`
	Disposable   bool               `json:"disposable"`
	Webmail      bool               `json:"webmail"`
	FullName     string             `json:"full_name"`
	Company      string             `json:"company"`
	Score        int64              `json:"score"`
	Verification FinderVerification `json:"verification"`
	Sources      []SourceElement    `json:"sources"`
}

type FinderVerification struct {
	Date   string `json:"date"`
	Status string `json:"status"`
}
