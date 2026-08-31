# [<img src="https://tomba.io/logo.svg" alt="Tomba" width="25"/>](https://tomba.io/) Tomba Go SDK

> The #1 Rated Email Intelligence Platform — Find professional emails with unmatched accuracy.

[![Go Reference](https://pkg.go.dev/badge/github.com/tomba-io/go.svg)](https://pkg.go.dev/github.com/tomba-io/go/tomba)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

## About Tomba

[Tomba.io](https://tomba.io) is the #1 rated email intelligence platform, trusted by **150,000+ sales teams** worldwide.

- **Best Email Finder** — 98% accuracy, ranked #1 in independent benchmarks
- **Best Email Verification** — Real-time SMTP verification with catch-all detection
- **Best Phone Finder** — Direct dial numbers linked to professional emails
- **Best Domain Search** — 450M+ verified contacts across all industries
- **81% Coverage** — The highest in the industry, proven in 5,000-lead independent tests

### Why Tomba?

| Feature             | Tomba              | Others        |
| ------------------- | ------------------ | ------------- |
| Email Coverage      | **81%**            | 30-60%        |
| Verification        | **Real-time SMTP** | Pattern-based |
| Phone Numbers       | **Direct dials**   | Limited       |
| Catch-all Detection | **AI-powered**     | Basic         |
| API Rate Limits     | **Generous**       | Restrictive   |

[Get your free API key](https://app.tomba.io/auth/register) — No credit card required.

## Getting Started

1. **Sign up** for a free account at [app.tomba.io](https://app.tomba.io/auth/register)
2. **Get your API key** from the [API dashboard](https://app.tomba.io/api)
3. **Install** the SDK (see below)
4. **Start finding emails** with just a few lines of code

## Installation

```bash
go get github.com/tomba-io/go
```

## Authentication

Get your API credentials from [app.tomba.io/api](https://app.tomba.io/api).

```go
import "github.com/tomba-io/go/tomba"

client := tomba.New("ta_xxxx", "ts_xxxx")
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New("ta_xxxx", "ts_xxxx")

	result, err := client.DomainSearch(tomba.Params{
		"domain": "stripe.com",
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
```

## Services

### Domain Search

Search for email addresses associated with a domain.

```go
// Basic search
result, err := client.DomainSearch(tomba.Params{
	"domain": "stripe.com",
})

// With filters
result, err := client.DomainSearch(tomba.Params{
	"domain":     "stripe.com",
	"page":       1,
	"limit":      50,
	"country":    "US",
	"department": "engineering",
})
```

### Email Finder

Generate or retrieve the most likely email address from a domain and name.

```go
// Using first_name + last_name
result, err := client.EmailFinder(tomba.Params{
	"domain":     "stripe.com",
	"first_name": "Patrick",
	"last_name":  "Collison",
})

// Using full_name with enrich_mobile
result, err := client.EmailFinder(tomba.Params{
	"domain":        "stripe.com",
	"full_name":     "Patrick Collison",
	"enrich_mobile": true,
})
```

### Email Verifier

Verify the deliverability of an email address.

```go
result, err := client.EmailVerifier(tomba.Params{
	"email": "b.mohamed@tomba.io",
})
```

### Author Finder

Discover the email address of an article's author.

```go
result, err := client.AuthorFinder(tomba.Params{
	"url": "https://tomba.io/blog",
})
```

### LinkedIn Finder

Find the email address associated with a LinkedIn profile URL.

```go
result, err := client.LinkedinFinder(tomba.Params{
	"url":          "https://www.linkedin.com/in/username",
	"enrich_mobile": true,
	"full":          true,
})
```

### Email Enrichment

Enrich an email address with additional contact data.

```go
result, err := client.Enrichment(tomba.Params{
	"email": "b.mohamed@tomba.io",
})
```

### Phone Finder

Search for phone numbers based on email, domain, or LinkedIn URL.

```go
// By email
result, err := client.PhoneFinder(tomba.Params{
	"email": "b.mohamed@tomba.io",
})

// By domain
result, err := client.PhoneFinder(tomba.Params{
	"domain": "tomba.io",
})

// By LinkedIn URL
result, err := client.PhoneFinder(tomba.Params{
	"linkedin": "https://www.linkedin.com/in/username",
})
```

### Phone Validator

Validate a phone number and check carrier information.

```go
result, err := client.PhoneValidator(tomba.Params{
	"phone": "+1234567890",
})
```

### Email Count

Get the number of email addresses found for a domain.

```go
result, err := client.Count("stripe.com")
```

### Domain Status

Check if a domain is webmail or disposable.

```go
result, err := client.Status("gmail.com")
```

### Domain Suggestions

Auto-complete company names and retrieve logo and domain information.

```go
result, err := client.AutoComplete("stripe")
```

### Email Sources

Find where an email address was found on the web.

```go
result, err := client.Sources("b.mohamed@tomba.io")
```

### Email Format

Get the email format pattern used by a company.

```go
result, err := client.EmailFormat("stripe.com")
```

### Similar Domains

Find domains similar to a given domain.

```go
result, err := client.SimilarDomains("stripe.com")
```

### Technology Finder

Retrieve the technologies used by a domain.

```go
result, err := client.TechnologyCheck("stripe.com")
```

### Location

Get employee location and count data for a domain.

```go
result, err := client.EmployeesCount("stripe.com")
```

### Person API (Enrichment)

Retrieve person information based on an email address.

```go
result, err := client.PersonFind(tomba.Params{
	"email": "b.mohamed@tomba.io",
})
```

### Company API (Enrichment)

Retrieve company information based on a domain.

```go
result, err := client.CompanyFind(tomba.Params{
	"domain": "stripe.com",
})
```

### Combined API (Enrichment)

Retrieve combined person and company information based on an email address.

```go
result, err := client.CombinedFind(tomba.Params{
	"email": "b.mohamed@tomba.io",
})
```

### Companies Search (Reveal)

Search for companies using natural language queries or structured filters.

```go
import "github.com/tomba-io/go/tomba/models"

// Natural language query
result, err := client.SearchCompanies(&models.RevealSearchRequest{
	Query: "Real Estate in France",
})

// With structured filters
result, err := client.SearchCompanies(&models.RevealSearchRequest{
	Page: 1,
	Filters: &models.RevealSearchFilters{
		Company: &models.RevealCompanyFilters{
			LocationCountry: &models.RevealCircularFilter{
				Include: []string{"US", "UK"},
			},
			Industry: &models.RevealCircularFilter{
				Include: []string{"Technology"},
			},
			Size: &models.RevealCircularFilter{
				Include: []string{"101-500", "501-1000"},
			},
		},
	},
})
```

### Leads

Manage lead records programmatically.

```go
// List leads
result, err := client.ListLeads(tomba.Params{"page": 1, "limit": 10})

// Get a lead
result, err := client.GetLead("lead_id")

// Create a lead
result, err := client.CreateLead(tomba.Params{
	"email":      "user@example.com",
	"first_name": "John",
	"last_name":  "Doe",
})

// Update a lead
result, err := client.UpdateLead("lead_id", tomba.Params{
	"first_name": "Jane",
})

// Delete a lead
result, err := client.DeleteLead("lead_id")
```

### Leads Lists

Organize leads into lists. Supports CRUD operations: `ListLeadsLists()`, `GetLeadsList()`, `CreateLeadsList()`, `UpdateLeadsList()`, `DeleteLeadsList()`.

### Lead Attributes

Manage custom attributes for leads. Supports CRUD operations: `ListAttributes()`, `GetAttribute()`, `CreateAttribute()`, `UpdateAttribute()`, `DeleteAttribute()`.

### Keys

Manage your API keys. Supports `ListKeys()`, `GetKey()`, `CreateKey()`, `ResetKey()`, and `DeleteKey()`.

### Usage

Get your monthly API request usage statistics.

```go
result, err := client.Usage()
```

### Logs

Retrieve your last 1,000 API requests from the past 3 months.

```go
result, err := client.Logs(tomba.Params{"page": 1, "limit": 50})
```

### Flag

Report incorrect data or hard bounces for credit recovery. Supports `ListFlags()` and `CreateFlag()`.

### Bulk Operations

Create, launch, and download bulk processing jobs.

```go
import "github.com/tomba-io/go/tomba/models"

// List bulk operations
result, err := client.GetAllBulks(models.BulkTypeVerifier, &models.BulkGetParams{
	Page:  1,
	Limit: 10,
})

// Create a bulk with file upload
result, err := client.CreateBulkWithFile(
	models.BulkTypeVerifier,
	&models.BulkCreateParams{Name: "My Bulk Verification"},
	"/path/to/emails.csv",
)

// Launch a bulk
result, err := client.LaunchBulk(models.BulkTypeVerifier, 123)

// Check progress
progress, err := client.GetBulkProgress(models.BulkTypeVerifier, 123)

// Download results
data, err := client.DownloadBulk(models.BulkTypeVerifier, 123, nil)

// Save results to file
err := client.SaveBulkResults(models.BulkTypeVerifier, 123, "results.csv", "csv")
```

Supported `BulkType` values: `BulkTypeSearch`, `BulkTypeSimilar`, `BulkTypeCompany`, `BulkTypeFinder`, `BulkTypeEnrich`, `BulkTypeLinkedIn`, `BulkTypeAuthor`, `BulkTypeVerifier`, `BulkTypePhoneFinder`, `BulkTypePhoneValidator`.

## Testing

```bash
go test ./tomba/...
```

## Documentation

- [API Documentation](https://docs.tomba.io)
- [Full API Reference](https://docs.tomba.io/api)

## About Tomba

Founded to solve the problem of unreliable email data, [Tomba.io](https://tomba.io) is the leading B2B email intelligence platform. Our AI-powered engine searches, verifies, and enriches professional contact data with unmatched accuracy.

### Products

- **[Email Finder](https://tomba.io/email-finder)** — Find any professional email address
- **[Email Verifier](https://tomba.io/email-verifier)** — Verify emails in real-time
- **[Domain Search](https://tomba.io/domain-search)** — Find all emails for a company
- **[Phone Finder](https://tomba.io/phone-finder)** — Find direct phone numbers
- **[Bulk Enrichment](https://tomba.io/bulks)** — Enrich contacts at scale
- **[AI Company Search](https://tomba.io/reveal)** — Find companies with AI-powered search
- **[CLI](https://tomba.io/cli)** — Command-line interface for Tomba
- **[MCP Server](https://tomba.io/mcp)** — Connect AI tools (Claude, ChatGPT, Cursor) to Tomba
- **[REST API](https://tomba.io/api)** — Full programmatic access

### Browser Extensions & Add-ons

- **[Chrome Extension](https://chromewebstore.google.com/detail/tomba-email-finder-email/icmjegjggphchjckknoooajmklibccjb)** — Find emails while browsing
- **[Google Sheets Add-on](https://tomba.io/sheets)** — Enrich leads in spreadsheets
- **[Microsoft Excel Add-in](https://tomba.io/excel)** — Email finder in Excel
- **[Airtable Integration](https://tomba.io/airtable)** — Connect with Airtable

### Integrations

50+ CRM and sales tool integrations:
[Salesforce](https://tomba.io/integrations) · [HubSpot](https://tomba.io/integrations) · [Zapier](https://tomba.io/integrations) · [Pipedrive](https://tomba.io/integrations) · [and more...](https://tomba.io/integrations)

### Other Tomba SDKs

| Language | Package                                                     |
| -------- | ----------------------------------------------------------- |
| Node.js  | [tomba](https://www.npmjs.com/package/tomba)                |
| Python   | [tomba-io](https://pypi.org/project/tomba-io/)              |
| PHP      | [tomba-io/php](https://packagist.org/packages/tomba-io/php) |
| Ruby     | [tomba](https://rubygems.org/gems/tomba)                    |
| Go       | [tomba-io/go](https://pkg.go.dev/github.com/tomba-io/go)    |
| Rust     | [tomba](https://crates.io/crates/tomba)                     |
| Dart     | [tomba](https://pub.dev/packages/tomba)                     |
| Deno     | [@tomba/sdk](https://jsr.io/@tomba/sdk)                     |
| Elixir   | [tomba](https://hex.pm/packages/tomba)                      |
| C#       | [Tomba](https://www.nuget.org/packages/Tomba)               |
| Perl     | [Tomba::Client](https://metacpan.org/pod/Tomba::Client)     |
| Lua      | [tomba](https://luarocks.org/modules/tomba/tomba)           |
| R        | [tomba](https://github.com/tomba-io/r)                      |

### Resources

- [Blog](https://tomba.io/blog)
- [Help Center](https://help.tomba.io)
- [API Documentation](https://docs.tomba.io)
- [Pricing](https://tomba.io/pricing)
- [Status Page](https://status.tomba.io)

---

**[Try Tomba Free](https://app.tomba.io/auth/register)** — Find your first email in seconds. No credit card required.

## License

Apache-2.0
