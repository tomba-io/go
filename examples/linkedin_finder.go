package main

import (
	"fmt"

	"github.com/tomba-io/go/tomba"
)

func main() {
	tomba := tomba.New("ta_xxxxx", "ts_xxxxx")

	result, err := tomba.LinkedinFinder("https://www.linkedin.com/in/alex-maccaw-ab592978")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Print(result)
}
