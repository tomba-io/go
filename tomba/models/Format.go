package models

import "encoding/json"

func UnmarshalFormat(data []byte) (Format, error) {
	var r Format
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Format) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Format struct {
	Data []FormatData `json:"data"`
}

type FormatData struct {
	Format     string `json:"format"`
	Percentage int64  `json:"percentage"`
}
