// Package luhn implements the Luhn algorithm, or Luhn formula, also known as the "modulus 10" or "mod 10" algorithm.
package luhn

type Error string

func (e Error) Error() string { return "luhn: " + string(e) }

// mod10 is the pre-calculated value of 10 minus the modulos-10 of every digit
const mod10 = "0987654321"

// doubles is the pre-calculated double of every digit
// When the result of doubling is greater than 9 the digits are added together; 16 => 1 + 6 => 7
var doubles = [10]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// Checksum returns the Lühn check digit for number.
func Checksum(number string) (string, error) {
	if number == "" {
		return "", Error("input is too short")
	}

	// determine even or odd length, so we know when to double
	double := 1&len(number) == 1

	var sum int
	for _, c := range []byte(number) {
		// assert character is numeric
		if c < '0' || c > '9' {
			return "", Error("input contains a non-number; " + string(c))
		}

		// numeric value from character
		n := int(c - '0')

		// double every other digit
		if double {
			n = doubles[n]
		}
		double = !double

		sum += n
	}

	unit := sum % 10
	return mod10[unit : unit+1], nil
}

// Valid returns if number verifies against its appended check digit
func Valid(number string) bool {
	l := len(number) - 1
	if l < 1 {
		return false
	}

	number, check := number[:l], number[l:]
	if digit, err := Checksum(number); err == nil {
		return check == digit
	}

	return false
}

// Sign returns number with its Lühn check digit appended
func Sign(number string) (string, error) {
	digit, err := Checksum(number)
	if err != nil {
		return "", err
	}

	return number + digit, nil
}
