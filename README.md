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

	numbers :=

	//Split a string of numbers
	first, second := phonenumbers.SplitRange("02080668850 - 02080668862", "-")
	fmt.Println("First number", first)
	fmt.Println("Second number", second)

	//Convert a number to UK e164 (441582123456):
	phonenumbers.ConverttoE164("01582123456")

	//Retrieve the last X trailing digits of a number
	phonenumbers.RetrieveTrailingDigits(02080668850, 3)

	//Count the digits in a number
	phonenumbers.DigitCount("02080668850")
	
	//return all the numbers in the middle of this range
	phonenumbers.ListRange(02080668850, 02080668862)
	
	//Return map of countries and their relevant prefixes 
	GeoData(nil)
	
	//Return prefix of a specific location
	GeoData("United Kingdom")
	
	//Return the name of country based of prefix provided
	GeoData(44)
	

}
```
