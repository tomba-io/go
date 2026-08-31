package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// Account returns information about the current account.
// See https://docs.tomba.io/api/account#get-account
func (conf *Tomba) Account() (models.Account, error) {
	account := models.Account{}
	str, err := conf.TombaCall(ACCOUNT_PATH, nil, nil, nil)
	if err != nil {
		return account, err
	}
	data, err := models.UnmarshalAccount(str)
	if err != nil {
		return account, err
	}
	return data, nil
}
