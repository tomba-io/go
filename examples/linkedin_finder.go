package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
	"github.com/tomba-io/go/tomba/models"
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

	// LinkedIn Finder with full=true (returns array of FinderData)
	result3, err := client.LinkedinFinder(tomba.Params{
		"url":  "https://www.linkedin.com/in/davidebarros",
		"full": true,
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n\nLinkedIn Finder with full=true:")

	// Parse as Finder and extract []FinderData
	var finder models.Finder
	jsonData, _ := json.Marshal(result3)
	json.Unmarshal(jsonData, &finder)

	// Convert Data to []FinderData
	if dataArray, ok := finder.Data.([]interface{}); ok {
		var finderDataList []models.FinderData
		for _, item := range dataArray {
			itemJSON, _ := json.Marshal(item)
			var fd models.FinderData
			json.Unmarshal(itemJSON, &fd)
			finderDataList = append(finderDataList, fd)
		}

		// Print the 2 emails
		fmt.Println("Emails found:")
		for i, fd := range finderDataList {
			fmt.Printf("  %d: %s\n", i+1, fd.Email)
		}
	}
}
