# go-luhn
[![Build Status](https://travis-ci.com/ferdypruis/go-luhn.svg?branch=master)](https://travis-ci.com/ferdypruis/go-luhn)
[![GoDoc](https://godoc.org/github.com/ferdypruis/go-luhn?status.svg)](https://godoc.org/github.com/ferdypruis/go-luhn)
[![GolangCI](https://golangci.com/badges/github.com/ferdypruis/go-luhn.svg)](https://golangci.com/r/github.com/ferdypruis/go-luhn)

Create or validate a Lühn (mod 10) check digit in a numeric string in Go.

## Examples
Checksum returns the Lühn check digit for number.
```go
number := "7992739871"

checkdigit, _ := luhn.Checksum(number) // Ignoring error for simplicity
fmt.Println("The Lühn check digit for", number, "is", checkdigit)
```

Sign returns number with its Lühn check digit appended
```go
number := "7992739871"

number, _ = luhn.Sign(number) // Ignoring error for simplicity
fmt.Println("Your account number is", number)
```

Valid returns if number verifies against its appended check digit
```go
number := "79927398713"

if luhn.Valid(number) {
    fmt.Println("The number is valid")
} else {
    fmt.Println("The number is not valid")
}
```
