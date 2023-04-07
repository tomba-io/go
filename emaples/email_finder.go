package main

import (
	"fmt"

	"github.com/tomba-io/go/tomba"
)

func main() {
	tomba := tomba.New("ta_xxxxx", "ts_xxxxx")

	result, err := tomba.EmailFinder("stripe.com", "fname", "lname")
	if err == nil {
		fmt.Println(result)
	}
}
