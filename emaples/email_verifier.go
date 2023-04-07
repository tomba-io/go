package main

import (
	"fmt"

	"github.com/tomba-io/go/tomba"
)

func main() {
	tomba := tomba.New("ta_xxxxx", "ts_xxxxx")

	result, err := tomba.EmailVerifier("b.mohamed@tomba.io")
	if err == nil {
		fmt.Println(result)
	}
}
