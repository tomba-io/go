package models

import "encoding/json"

func UnmarshalCount(data []byte) (Count, error) {
	var r Count
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Count) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Count struct {
	Data CountData `json:"data"`
}

type CountData struct {
	Total          int64           `json:"total"`
	PersonalEmails int64           `json:"personal_emails"`
	GenericEmails  int64           `json:"generic_emails"`
	Department     CountDepartment `json:"department"`
	Seniority      CountSeniority  `json:"seniority"`
}

type CountDepartment struct {
	Engineering   int64 `json:"engineering"`
	Finance       int64 `json:"finance"`
	Hr            int64 `json:"hr"`
	It            int64 `json:"it"`
	Marketing     int64 `json:"marketing"`
	Operations    int64 `json:"operations"`
	Management    int64 `json:"management"`
	Sales         int64 `json:"sales"`
	Legal         int64 `json:"legal"`
	Support       int64 `json:"support"`
	Communication int64 `json:"communication"`
	Executive     int64 `json:"executive"`
}

type CountSeniority struct {
	Junior    int64 `json:"junior"`
	Senior    int64 `json:"senior"`
	Executive int64 `json:"executive"`
}
