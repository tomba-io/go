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
	Data interface{} `json:"data,omitempty"` // Can be FinderData or []FinderData
}

type FinderData struct {
	WebsiteURL   string              `json:"website_url,omitempty"`
	AcceptAll    bool                `json:"accept_all,omitempty"`
	Email        string              `json:"email,omitempty"`
	FirstName    string              `json:"first_name,omitempty"`
	LastName     string              `json:"last_name,omitempty"`
	FullName     string              `json:"full_name,omitempty"`
	Country      *string             `json:"country,omitempty"`
	Gender       string              `json:"gender,omitempty"`
	PhoneNumber  bool                `json:"phone_number,omitempty"`
	Position     string              `json:"position,omitempty"`
	Twitter      *string             `json:"twitter,omitempty"`
	Linkedin     *string             `json:"linkedin,omitempty"`
	Disposable   bool                `json:"disposable,omitempty"`
	Webmail      bool                `json:"webmail,omitempty"`
	Company      string              `json:"company,omitempty"`
	Score        int64               `json:"score,omitempty"`
	Verification *FinderVerification `json:"verification,omitempty"`
	Sources      []SourceElement     `json:"sources,omitempty"`
	Info         *AuthorInfo         `json:"info,omitempty"`
}

type FinderVerification struct {
	Date   string `json:"date,omitempty"`
	Status string `json:"status,omitempty"`
}

type AuthorInfo struct {
	URL           string   `json:"url,omitempty"`
	Title         string   `json:"title,omitempty"`
	OgImage       string   `json:"og_image,omitempty"`
	Description   string   `json:"description,omitempty"`
	FullName      string   `json:"full_name,omitempty"`
	FirstName     string   `json:"first_name,omitempty"`
	LastName      string   `json:"last_name,omitempty"`
	ArticleDomain string   `json:"article_domain,omitempty"`
	SameAsDomain  string   `json:"same_as_domain,omitempty"`
	Linkedin      string   `json:"linkedin,omitempty"`
	Twitter       string   `json:"twitter,omitempty"`
	Email         string   `json:"email,omitempty"`
	Image         string   `json:"image,omitempty"`
	Gravatar      bool     `json:"gravatar,omitempty"`
	Emails        []string `json:"emails,omitempty"`
	AuthorScore   int      `json:"author_score,omitempty"`
}
