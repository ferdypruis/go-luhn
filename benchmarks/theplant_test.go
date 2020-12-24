package benchmarks

import (
	"testing"

	"github.com/theplant/luhn"
)

func BenchmarkChecksum_theplant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			checksumInt = luhn.CalculateLuhn(in.i)
		}
	}
}

func BenchmarkValid_theplant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			valid = luhn.Valid(in.i)
		}
	}
}
