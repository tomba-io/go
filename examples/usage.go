package main

import (
	"fmt"
	"os"

	"github.com/tomba-io/go/tomba"
)

func main() {
	client := tomba.New(os.Getenv("TOMBA_API_KEY"), os.Getenv("TOMBA_SECRET_KEY"))

	result, err := client.Usage()
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Print(result)
}
