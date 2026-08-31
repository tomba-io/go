package tomba

import (
	"github.com/tomba-io/go/tomba/models"
)

// EmailFinder generates or retrieves the most likely email address from a domain name,
// a first name, and a last name.
//
// Supported params keys: domain, company, full_name, first_name, last_name, enrich_mobile, webhook_url.
//
// See https://docs.tomba.io/api/finder#email-finder
func (conf *Tomba) EmailFinder(params Params) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(FINDER_PATH, params, nil, nil)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder(str)
	if err != nil {
		return finder, err
	}
	return data, nil
}

// AuthorFinder generates or retrieves the most likely email address from a blog post URL.
//
// Supported params keys: url, webhook_url.
//
// See https://docs.tomba.io/api/finder#author-finder
func (conf *Tomba) AuthorFinder(params Params) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(AUTHOR_PATH, params, nil, nil)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder(str)
	if err != nil {
		return finder, err
	}
	return data, nil
}

// LinkedinFinder generates or retrieves the most likely email address from a LinkedIn URL.
//
// Supported params keys: url, enrich_mobile, full, webhook_url.
//
// See https://docs.tomba.io/api/finder#linkedin-finder
func (conf *Tomba) LinkedinFinder(params Params) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(LINKEDIN_PATH, params, nil, nil)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder(str)
	if err != nil {
		return finder, err
	}
	return data, nil
}

// Enrichment looks up person and company data based on an email address.
// For example, you can retrieve a person's name, location, and social handles from an email.
//
// Supported params keys: email, webhook_url.
//
// See https://docs.tomba.io/api/enrichment#enrich
func (conf *Tomba) Enrichment(params Params) (models.Finder, error) {
	finder := models.Finder{}
	str, err := conf.TombaCall(ENRICHMENT_PATH, params, nil, nil)
	if err != nil {
		return finder, err
	}
	data, err := models.UnmarshalFinder(str)
	if err != nil {
		return finder, err
	}
	return data, nil
}
