package main

import (
	"fmt"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New("", "")
	result, err := client.SimilarDomains("hunter.io")
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
