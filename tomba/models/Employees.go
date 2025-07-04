package models

import "encoding/json"

func UnmarshalEmployees(data []byte) (Employees, error) {
	var r Employees
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Employees) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Employees struct {
	Data []EmployeesData `json:"data"`
}

type EmployeesData struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
}
