package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))
	// Example 1: Find phone by email
	result, err := client.PhoneFinder(tomba.Params{
		"email": "info@stripe.com",
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Phone Finder by email:")
	phoneData, _ := result.GetSingleData()
	fmt.Printf("Phone: %s (Country: %s)\n", phoneData.IntlFormat, phoneData.CountryCode)

	// Example 2: Find phone by domain
	result2, err := client.PhoneFinder(tomba.Params{
		"domain": "tomba.io",
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\nPhone Finder by domain:")
	phoneData2, _ := result2.GetSingleData()
	fmt.Printf("Phone: %s (Country: %s)\n", phoneData2.IntlFormat, phoneData2.CountryCode)

	// Example 3: Find phone by LinkedIn URL
	result3, err := client.PhoneFinder(tomba.Params{
		"linkedin": "https://www.linkedin.com/in/alex-maccaw-ab592978",
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\nPhone Finder by LinkedIn:")
	phoneData3, _ := result3.GetSingleData()
	fmt.Printf("Phone: %s (Country: %s)\n", phoneData3.IntlFormat, phoneData3.CountryCode)

	// Example 4: Get all phone numbers (full=true)
	result4, err := client.PhoneFinder(tomba.Params{
		"domain": "stripe.com",
		"full":   true,
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\nAll phone numbers for domain:")
	allPhones, _ := result4.GetFullData()
	for i, phone := range allPhones {
		fmt.Printf("  %d. %s (%s)\n", i+1, phone.IntlFormat, phone.LineType)
	}
}
