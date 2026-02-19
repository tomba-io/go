package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))

	// Basic phone validation
	result, err := client.PhoneValidator(tomba.Params{
		"phone": "+14155552671",
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Phone Validation Result:")
	fmt.Printf("Valid: %v\n", result.Data.Valid)
	fmt.Printf("Local Format: %s\n", result.Data.LocalFormat)
	fmt.Printf("International Format: %s\n", result.Data.IntlFormat)
	fmt.Printf("E164 Format: %s\n", result.Data.E164Format)
	fmt.Printf("Country Code: %s\n", result.Data.CountryCode)
	fmt.Printf("Line Type: %s\n", result.Data.LineType)
	fmt.Printf("Carrier: %s\n", result.Data.Carrier)

	// Phone validation with country code
	resultWithCountry, err := client.PhoneValidator(tomba.Params{
		"phone":        "4155552671",
		"country_code": "US",
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\nPhone Validation with Country Code:")
	fmt.Printf("Valid: %v\n", resultWithCountry.Data.Valid)
	fmt.Printf("International Format: %s\n", resultWithCountry.Data.IntlFormat)
}
