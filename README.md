# go-luhn
[![Build Status](https://travis-ci.com/ferdypruis/go-luhn.svg?branch=master)](https://travis-ci.com/ferdypruis/go-luhn)
[![GoDoc](https://godoc.org/github.com/ferdypruis/go-luhn?status.svg)](https://godoc.org/github.com/ferdypruis/go-luhn)
[![GolangCI](https://golangci.com/badges/github.com/ferdypruis/go-luhn.svg)](https://golangci.com/r/github.com/ferdypruis/go-luhn)

Create or validate a Lühn (mod 10) check digit in a numeric string in Go.

## Usage
```go
package main

import (
	"fmt"
	"github.com/ferdypruis/go-luhn"
)

func main() {
	// Obtain the check digit for a numeric string
	number := "7992739871"
	checkdigit, _ := luhn.Checksum(number) // Ignoring error for simplicity
	fmt.Printf("The Lühn check digit for %s is %s.\n", number, checkdigit)

	// Add the Lühn check digit to a numeric string
	number, _ = luhn.Sign(number) // Ignoring error for simplicity
	fmt.Printf("Your account number is %s.\n", number)

	// Verify a number against its included check digit
	if luhn.Valid(number) {
		fmt.Print("The number is valid.\n")
	} else {
		fmt.Print("The number is not valid.\n")
	}
}
```
