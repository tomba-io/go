package main

import (
	"fmt"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New("ta_xxxxx", "ts_xxxxx")
	result, err := client.EmailFinder(tomba.Params{"domain": "asana.com", "full_name": "moskoz dustin"})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Print(result)
}
