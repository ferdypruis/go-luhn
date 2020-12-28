// Package luhn implements the Luhn algorithm, or Luhn formula, also known as the "modulus 10" or "mod 10" algorithm.
package luhn

import (
	"fmt"
)

// mod10 is the pre-calculated value of 10 minus the modulos-10 of every digit
const mod10 = "0987654321"

// doubles is the pre-calculated double of every digit
// When the result of doubling is greater than 9 the digits are added together; 16 => 1 + 6 => 7
var doubles = [10]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// Checksum returns the check digit for the number.
func Checksum(number string) (string, error) {
	l := len(number)
	if l == 0 {
		return "", Error("input is too short")
	}

	var sum int

	for i := l - 1; i >= 0; i -= 2 {
		n := int(number[i] - '0')
		if n < 0 || n > 9 {
			return "", Error(fmt.Sprintf("character %q at position %d is not a digit", number[i], i))
		}

		sum += doubles[n]
	}

	for i := l - 2; i >= 0; i -= 2 {
		n := int(number[i] - '0')
		if n < 0 || n > 9 {
			return "", Error(fmt.Sprintf("character %q at position %d is not a digit", number[i], i))
		}

		sum += n
	}

	unit := sum % 10
	return mod10[unit:unit+1], nil
}

// Valid returns whether number verifies against its appended check digit.
func Valid(number string) bool {
	l := len(number)
	if l < 2 {
		return false
	}

	var sum = int(number[l-1] - '0')

	for i := l - 2; i >= 0; i -= 2 {
		n := int(number[i] - '0')
		if n < 0 || n > 9 {
			return false
		}

		sum += doubles[n]
	}

	for i := l - 3; i >= 0; i -= 2 {
		n := int(number[i] - '0')
		if n < 0 || n > 9 {
			return false
		}

		sum += n
	}

	return sum % 10 == 0
}

// Sign returns the partial number with its Lühn check digit appended, creating a full number.
func Sign(partial string) (string, error) {
	checkdigit, err := Checksum(partial)
	if err != nil {
		return "", err
	}

	return partial + checkdigit, nil
}

type Error string
func (e Error) Error() string { return "luhn: " + string(e) }
