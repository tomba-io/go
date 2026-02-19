package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))

	// Basic LinkedIn Finder
	result, err := client.LinkedinFinder(tomba.Params{
		"url": "https://www.linkedin.com/in/alex-maccaw-ab592978",
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Basic LinkedIn Finder:")
	fmt.Print(result)

	// LinkedIn Finder with mobile phone lookup
	result2, err := client.LinkedinFinder(tomba.Params{
		"url":           "https://www.linkedin.com/in/alex-maccaw-ab592978",
		"enrich_mobile": true, // Set to true to get the phone number associated with the email
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n\nLinkedIn Finder with enrich_mobile:")
	fmt.Print(result2)
}
