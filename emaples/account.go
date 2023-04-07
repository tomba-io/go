package main

import (
	"fmt"

	"github.com/tomba-io/go/tomba"
)

func main() {
	tomba := tomba.New("ta_xxxxx", "ts_xxxxx")

	result, err := tomba.Account()
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(result.Data.Email)
	}
}
