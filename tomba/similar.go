package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// SimilarDomains retrieves a list of similar domains based on a given domain name.
// See https://docs.tomba.io/api/~endpoints#similar
func (conf *Tomba) SimilarDomains(domain string) (models.Similar, error) {
	similarDomains := models.Similar{}
	str, err := conf.TombaCall(SIMILAR_PATH, Params{"domain": domain}, nil, nil)
	if err != nil {
		return similarDomains, err
	}
	data, err := models.UnmarshalSimilar(str)
	if err != nil {
		return similarDomains, err
	}
	return data, nil
}
