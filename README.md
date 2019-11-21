# go-luhn
[![Build Status](https://travis-ci.com/ferdypruis/go-luhn.svg?branch=master)](https://travis-ci.com/ferdypruis/go-luhn)
[![GoDoc](https://godoc.org/github.com/ferdypruis/go-luhn?status.svg)](https://godoc.org/github.com/ferdypruis/go-luhn)
[![GolangCI](https://golangci.com/badges/github.com/ferdypruis/go-luhn.svg)](https://golangci.com/r/github.com/ferdypruis/go-luhn)

Create or validate a Lühn (mod 10) check digit in a numeric string in Go.

## Usage
```go
package main

import (
	"github.com/ferdypruis/go-luhn"
)

func main() {
    // Verify a number against its included check digit
    valid, err := luhn.Valid("79927398713") // valid = true
    
    // Obtain the check digit for a numeric string
    checkdigit, err := luhn.Checksum("7992739871") // checkdigit = "3"
    	
    // Add the Lühn check digit to a numeric string
    number, err := luhn.Sign("7992739871") // number = "79927398713"
}
```
