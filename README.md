# Helper functions for working with UK phone numbers.

[![Tests](https://github.com/AaronVoIP/phonenumbers/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/AaronVoIP/phonenumbers/actions/workflows/main.yml)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](https://pkg.go.dev/github.com/AaronVoIP/phonenumbers)

Example Usage

```go
package main

import (
	"fmt"

	"github.com/AaronVoIP/phonenumbers"
)

func main() {

	numbers := "02080668850 - 02080668862"

	first, second := phonenumbers.SplitRange(numbers, "-")
	fmt.Println("First number", first)
	fmt.Println("Second number", second)

	fmt.Println("Convert a number to UK e164:", phonenumbers.ConverttoE164("01582123456"))

	fmt.Println("Retrieve the last three trailing digits of a number:", phonenumbers.RetrieveTrailingDigits(first, 3))

	fmt.Println("Count the digits in a number:", phonenumbers.DigitCount(first))

	fmt.Println("Break our original range into individual numbers:", phonenumbers.ListRange(first, second))
	
	//Return map of countries and their relevant prefixes 
	GeoData(nil)
	
	//Return prefix of a specific location
	GeoData("United Kingdom")
	
	//Return the name of country based of prefix provided
	GeoData(44)
	

}
```
