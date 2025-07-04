package models

import "encoding/json"

func UnmarshalTechnology(data []byte) (Technology, error) {
	var r Technology
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Technology) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Technology struct {
	Domain string           `json:"domain"`
	Data   []TechnologyData `json:"data"`
}

type TechnologyData struct {
	Slug       string               `json:"slug"`
	Name       string               `json:"name"`
	Icon       string               `json:"icon"`
	Website    string               `json:"website"`
	Categories TechnologyCategories `json:"categories"`
}

type TechnologyCategories struct {
	ID   int64  `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}
