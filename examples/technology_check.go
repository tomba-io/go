package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))
	result, err := client.TechnologyCheck("tomba.io")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Domain:", result.Domain)
	fmt.Println("Total Technologies:", len(result.Data))
	for _, tech := range result.Data {
		fmt.Printf("- %s (%s) - Category: %s\n", tech.Name, tech.Website, tech.Categories.Name)
	}
}
