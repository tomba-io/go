package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
	"github.com/tomba-io/go/tomba/models"
)

func main() {
	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))

	// Example 1: Simple natural language query
	result, err := client.SearchCompanies(&models.RevealSearchRequest{
		Query: "Real Estate in Europe",
		Page:  1,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Simple Query Results:")
	if result.Meta != nil {
		fmt.Printf("Total: %d, Page: %d, Pages: %d\n", result.Meta.Total, result.Meta.Page, result.Meta.Pages)
	}
	for _, company := range result.Data.Companies {
		fmt.Printf("- %s (%s, %s)\n", company.Name, company.Industry, company.Country)
	}

	// Example 2: Using structured filters
	resultFiltered, err := client.SearchCompanies(&models.RevealSearchRequest{
		Page: 1,
		Filters: &models.RevealFilters{
			LocationCountry: &models.SearchFilter{
				Include: []string{"US", "UK"},
			},
			Industry: &models.SearchFilter{
				Include: []string{"Technology"},
			},
			Size: &models.SearchFilter{
				Include: []string{"101-500", "501-1000"},
			},
		},
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\nFiltered Results:")
	fmt.Printf("Total: %d\n", resultFiltered.Data.Total)
	for _, company := range resultFiltered.Data.Companies {
		fmt.Printf("- %s (%s, %s) - %s employees\n", company.Name, company.Industry, company.Country, company.CompanySize)
	}
}
