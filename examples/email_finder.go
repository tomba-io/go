package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))

	// Example 1: Using full_name
	result, err := client.EmailFinder(tomba.Params{
		"domain":    "asana.com",
		"full_name": "moskoz dustin",
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Using full_name:")
	fmt.Print(result)

	// Example 2: Using first_name + last_name
	result2, err := client.EmailFinder(tomba.Params{
		"domain":     "stripe.com",
		"first_name": "Patrick",
		"last_name":  "Collison",
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n\nUsing first_name + last_name:")
	fmt.Print(result2)

	// Example 3: With enrich_mobile to get phone number
	result3, err := client.EmailFinder(tomba.Params{
		"domain":        "tomba.io",
		"full_name":     "Mohamed Ben Rebia",
		"enrich_mobile": true, // Set to true to get the phone number associated with the email
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n\nWith enrich_mobile:")
	fmt.Print(result3)
}
