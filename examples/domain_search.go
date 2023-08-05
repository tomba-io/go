package main

import (
	"fmt"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New("ta_xxxxx", "ts_xxxxx")
	result, err := client.DomainSearch(tomba.Params{
		"domain": "tomba.io",
	})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Print(result)
}
