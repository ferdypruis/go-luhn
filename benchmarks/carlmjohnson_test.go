package benchmarks

import (
	"testing"

	"github.com/carlmjohnson/luhn"
)

func BenchmarkChecksum_carlmjohnson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			checksumInt = luhn.CheckDigitInt(in.i)
		}
	}
}

func BenchmarkValid_carlmjohnson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			valid = luhn.IsValid(in.s)
		}
	}
}
