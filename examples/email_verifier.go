package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))

	// Basic email verification
	result, err := client.EmailVerifier(tomba.Params{
		"email": "b.mohamed@tomba.io",
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Basic Email Verification:")
	fmt.Print(result)

	// Email verification with mobile phone lookup
	result2, err := client.EmailVerifier(tomba.Params{
		"email":         "b.mohamed@tomba.io",
		"enrich_mobile": true, // Set to true to get the phone number associated with the email
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n\nEmail Verification with enrich_mobile:")
	fmt.Print(result2)
}
