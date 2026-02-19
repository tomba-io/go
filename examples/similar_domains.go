package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))
	result, err := client.SimilarDomains("tomba.io")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Total Results: ", len(result.Data))
	for _, domain := range result.Data {
		name := ""
		if domain.Name != nil {
			name = *domain.Name
		}
		fmt.Println("Domain: ", name, " - WebsiteURL: ", domain.WebsiteURL)
	}
}
