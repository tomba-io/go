package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))

	// Basic enrichment
	result, err := client.Enrichment(tomba.Params{
		"email": "b.mohamed@tomba.io",
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Basic Enrichment:")
	fmt.Print(result)

	// Enrichment with mobile phone lookup
	result2, err := client.Enrichment(tomba.Params{
		"email":         "b.mohamed@tomba.io",
		"enrich_mobile": true, // Set to true to get the phone number associated with the email
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n\nEnrichment with enrich_mobile:")
	fmt.Print(result2)
}
