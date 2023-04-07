package tomba

// Tomba init
type Tomba struct {
	ApiKey    string
	ApiSecret string
}

// New Constructor
func New(key string, secret string) *Tomba {
	p := new(Tomba)
	p.ApiKey = key
	p.ApiSecret = secret
	return p
}
