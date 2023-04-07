package models

import "encoding/json"

func UnmarshalStatus(data []byte) (Status, error) {
	var r Status
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Status) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Status struct {
	Domain     string `json:"domain"`
	Webmail    bool   `json:"webmail"`
	Disposable bool   `json:"disposable"`
}
