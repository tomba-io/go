package models

import "encoding/json"

func UnmarshalLogs(data []byte) (Logs, error) {
	var r Logs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Logs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Logs struct {
	Data []LogsData `json:"data"`
}

type LogsData struct {
	Type       string `json:"type"`
	URI        string `json:"uri"`
	Cost       bool   `json:"cost"`
	IPAddress  string `json:"ip_address"`
	CreatedAt  string `json:"created_at"`
	ID         int64  `json:"id"`
	UserAgent  string `json:"user_agent"`
	Source     string `json:"source"`
	HTTPMethod string `json:"http_method"`
	Country    string `json:"country"`
}
