package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// Logs returns the last 1,000 requests you made during the last 3 months.
//
// Supported params keys: page, limit.
//
// See https://docs.tomba.io/api/account#retrieve-api-logs#get-logs
func (conf *Tomba) Logs(params ...Params) (models.Logs, error) {
	logs := models.Logs{}
	var p Params
	if len(params) > 0 {
		p = params[0]
	}
	str, err := conf.TombaCall(LOGS_PATH, p, nil, nil)
	if err != nil {
		return logs, err
	}
	data, err := models.UnmarshalLogs(str)
	if err != nil {
		return logs, err
	}
	return data, nil
}
