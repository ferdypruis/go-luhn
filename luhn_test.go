package luhn_test

import (
	"fmt"
	"testing"

	"github.com/ferdypruis/go-luhn"
)

func TestChecksum(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"0", "0"},
		{"1", "8"},
		{"05", "9"},
		{"009", "1"},
		{"99999999999999990", "6"},
		{"99999999999999909", "6"},
		{"7992739871", "3"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := luhn.Checksum(tt.in)
			if err != nil {
				t.Errorf("Checksum() error = %s", err)
			}
			if got != tt.want {
				t.Errorf("Checksum() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestChecksumShortInputError(t *testing.T) {
	_, err := luhn.Checksum("")
	if _, ok := err.(luhn.Error); !ok {
		t.Errorf("Checksum() error = %s, want luhn.Error", err)
	}
}

func TestChecksumNotNumericError(t *testing.T) {
	tests := []string{
		"nein",
		"1mississippi",
		"2three4",
		"?",
		"三",
		"3.14159265359",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			_, err := luhn.Checksum(tt)
			if _, ok := err.(luhn.Error); !ok {
				t.Errorf("Checksum() error = %s, want luhn.Error", err)
			}
		})
	}
}

func TestValid(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"00", true},
		{"18", true},
		{"059", true},
		{"0091", true},
		{"1234", false},
		{"01234", false},
		{"001234", false},
		{"79927398710", false},
		{"79927398711", false},
		{"79927398712", false},
		{"79927398713", true},
		{"79927398714", false},
		{"79927398715", false},
		{"79927398716", false},
		{"79927398717", false},
		{"79927398718", false},
		{"79927398719", false},
		{"999999999999999906", true},
		{"999999999999999096", true},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := luhn.Valid(tt.in)
			if err != nil {
				t.Errorf("Valid() error = %s", err)
			}
			if got != tt.want {
				t.Errorf("Valid() got = %t, want %t", got, tt.want)
			}
		})
	}
}

func TestValidShortInputError(t *testing.T) {
	tests := []string{
		"",
		"0",
		"1",
		"5",
		"9",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			_, err := luhn.Valid(tt)
			if _, ok := err.(luhn.Error); !ok {
				t.Errorf("Valid() error = %s, want luhn.Error", err)
			}
		})
	}
}

func TestValidNotNumericError(t *testing.T) {
	tests := []string{
		`nein`,
		`1mississippi`,
		`2three4`,
		`!?`,
		`测试`,
		`3.14`,
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			_, err := luhn.Valid(tt)
			if _, ok := err.(luhn.Error); !ok {
				t.Errorf("Valid() error = %s, want NotNumericError", err)
			}
		})
	}
}

func TestSign(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"0", "00"},
		{"1", "18"},
		{"05", "059"},
		{"009", "0091"},
		{"7992739871", "79927398713"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := luhn.Sign(tt.in)
			if err != nil {
				t.Errorf("Sign() unexpected error = %v", err)
			}
			if got != tt.want {
				t.Errorf("Sign() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignShortInputError(t *testing.T) {
	_, err := luhn.Sign("")
	if _, ok := err.(luhn.Error); !ok {
		t.Errorf("Sign() error = %s, want luhn.Error", err)
	}
}

func TestSignNotNumericError(t *testing.T) {
	tests := []string{
		`nein`,
		`1mississippi`,
		`2three4`,
		`!?`,
		`测试`,
		`3.14`,
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			_, err := luhn.Sign(tt)
			if _, ok := err.(luhn.Error); !ok {
				t.Errorf("Sign() error = %s, want luhn.Error", err)
			}
		})
	}
}

func TestError(t *testing.T) {
	err := luhn.Error("message")
	want := "luhn: message"
	if got := err.Error(); got != want {
		t.Errorf("Error() got = %s, want %s", got, want)
	}
}

func ExampleValid() {
	number := "79927398713"
	if valid, _ := luhn.Valid(number); valid { // Ignoring error for simplicity
		fmt.Print("The number is valid.")
	} else {
		fmt.Print("The number is not valid.")
	}

	// Output:
	// The number is valid.
}

func ExampleSign() {
	number := "7992739871"
	number, _ = luhn.Sign(number) // Ignoring error for simplicity
	fmt.Printf("Your account number is %s.", number)

	// Output:
	// Your account number is 79927398713.
}

func ExampleChecksum() {
	number := "7992739871"
	checkdigit, _ := luhn.Checksum(number) // Ignoring error for simplicity
	fmt.Printf("The Lühn check digit for %s is %s.", number, checkdigit)

	// Output:
	// The Lühn check digit for 7992739871 is 3.
}
