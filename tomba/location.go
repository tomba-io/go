package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// EmployeesCount retrieves the number of employees in each country for a given domain.
// See https://docs.tomba.io/api/finder#employees
func (conf *Tomba) EmployeesCount(domain string) (models.Employees, error) {
	employees := models.Employees{}
	str, err := conf.TombaCall(EMPLOYEES_PATH, Params{"domain": domain}, nil, nil)
	if err != nil {
		return employees, err
	}
	data, err := models.UnmarshalEmployees(str)
	if err != nil {
		return employees, err
	}
	return data, nil
}
