# Tomba Go Client Examples

This folder contains example code demonstrating how to use the Tomba Go client library.

## Setup

Before running any example, make sure to replace `ta_xxxxx` and `ts_xxxxx` with your actual Tomba API key and secret.

```go
client := tomba.New("your_api_key", "your_api_secret")
```

Get your API keys at [https://app.tomba.io/api-keys](https://app.tomba.io/api-keys)

## Setup Environment Variables

### Initial Setup

```bash
cp tomba.env.example tomba.env
```

### Environment Variable

Update the development environment with your TOMBA_API_KEY and TOMBA_SECRET_KEY

```bash
echo "export TOMBA_API_KEY=ta_xxxxx" >> tomba.env
echo "export TOMBA_SECRET_KEY=ts_xxxxx" >> tomba.env
```

### Load Environment Variables

```bash
source tomba.env
```

## Examples

### Account

Get information about your Tomba account.

```bash
go run account.go
```

### Domain Search

Search for email addresses associated with a domain.

**Basic Usage:**

```go
client.DomainSearch(tomba.Params{
    "domain": "tomba.io",
})
```

**Advanced Usage with Query Parameters:**

```go
client.DomainSearch(tomba.Params{
    "domain":     "stripe.com",   // Domain name or company name
    "country":    "US",           // Filter by country
    "limit":      50,             // Number of results per page (max 100)
    "page":       1,              // Page number for pagination
    "department": "engineering",  // Filter by department
})
```

Available department values: `executive`, `it`, `finance`, `management`, `communication`, `marketing`, `sales`, `legal`, `hr`, `support`, `engineering`

```bash
go run domain_search.go
```

### Email Finder

Find the email address of a person.

**Using full_name:**

```go
client.EmailFinder(tomba.Params{
    "domain":    "asana.com",
    "full_name": "dustin moskovitz",
})
```

**Using first_name + last_name:**

```go
client.EmailFinder(tomba.Params{
    "domain":     "stripe.com",
    "first_name": "Patrick",
    "last_name":  "Collison",
})
```

**With enrich_mobile (get phone number):**

```go
client.EmailFinder(tomba.Params{
    "domain":        "tomba.io",
    "full_name":     "Mohamed Ben Rebia",
    "enrich_mobile": true,
})
```

```bash
go run email_finder.go
```

### Email Verifier

Verify the deliverability of an email address.

**Basic Usage:**

```go
client.EmailVerifier(tomba.Params{
    "email": "b.mohamed@tomba.io",
})
```

**With enrich_mobile:**

```go
client.EmailVerifier(tomba.Params{
    "email":         "b.mohamed@tomba.io",
    "enrich_mobile": true,
})
```

```bash
go run email_verifier.go
```

### Enrichment

Look up person and company data based on an email.

**Basic Usage:**

```go
client.Enrichment(tomba.Params{
    "email": "b.mohamed@tomba.io",
})
```

**With enrich_mobile:**

```go
client.Enrichment(tomba.Params{
    "email":         "b.mohamed@tomba.io",
    "enrich_mobile": true,
})
```

```bash
go run enrichment.go
```

### LinkedIn Finder

Find email address from a LinkedIn profile URL.

**Basic Usage:**

```go
client.LinkedinFinder(tomba.Params{
    "url": "https://www.linkedin.com/in/alex-maccaw-ab592978",
})
```

**With enrich_mobile:**

```go
client.LinkedinFinder(tomba.Params{
    "url":           "https://www.linkedin.com/in/alex-maccaw-ab592978",
    "enrich_mobile": true,
})
```

```bash
go run linkedin_finder.go
```

### Author Finder

Find the email address of the author of a blog post.

```go
client.AuthorFinder("https://clearbit.com/blog/company-name-to-domain-api")
```

```bash
go run author_finder.go
```

### Search Companies

Search for companies using natural language queries or structured filters.

**Natural Language Query:**

```go
client.SearchCompanies(&models.RevealSearchRequest{
    Query: "Real Estate in France",
})
```

**Using Structured Filters:**

```go
client.SearchCompanies(&models.RevealSearchRequest{
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

Available filter options:

- `LocationCountry` - Country codes (e.g., "US", "UK", "FR")
- `LocationCity` - City names
- `LocationState` - State/province names
- `Industry` - Industry sectors
- `Size` - Company size ranges (e.g., "1-10", "11-50", "51-100", "101-500", "501-1000", "1001-5000", "5001+")
- `Type` - Company type (e.g., "Private", "Public")
- `Keywords` - Keywords to search for
- `Founded` - Year founded
- `Technologies` - Technologies used by the company
- `Similar` - Similar company domains
- `Revenue` - Revenue ranges
- `SIC` - SIC codes
- `NAICS` - NAICS codes

```bash
go run search_companies.go
```

### Other Examples

- **count.go** - Get the total number of emails for a domain
- **status.go** - Check if a domain is webmail or disposable
- **source.go** - Get email sources
- **similar_domains.go** - Find similar domains
- **logs.go** - View your API request logs
- **usage.go** - Check your monthly API usage

### Bulk Operations

See the `bulks/` folder for bulk operation examples.

## Documentation

For full API documentation, visit [https://docs.tomba.io](https://docs.tomba.io)
