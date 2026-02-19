package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
)

func main() {

	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))

	// Basic domain search
	result, err := client.DomainSearch(tomba.Params{
		"domain": "tomba.io",
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Basic Search:")
	fmt.Print(result)

	// Advanced domain search with all query parameters
	advancedResult, err := client.DomainSearch(tomba.Params{
		"domain":     "stripe.com",  // Domain name or company name
		"country":    "US",          // Filter by country
		"limit":      50,            // Number of results per page (max 100)
		"page":       1,             // Page number for pagination
		"department": "engineering", // Filter by department (executive, it, finance, management, communication, marketing, sales, legal, hr, support, engineering)
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n\nAdvanced Search with filters:")
	fmt.Printf("Total: %d, Page: %d, Total Pages: %d\n", advancedResult.Meta.Total, advancedResult.Meta.Current, advancedResult.Meta.TotalPages)
	for _, email := range advancedResult.Data.Emails {
		fmt.Printf("- %s (%s, %s)\n", email.Email, email.Department, email.Position)
	}
}
