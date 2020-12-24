package benchmarks

import (
	"testing"

	"github.com/jancajthaml-go/luhn"
)

func BenchmarkChecksum_jancajthaml(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			checksumInt, _ = luhn.Digit(in.s)
		}
	}
}

func BenchmarkValid_jancajthaml(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			valid = luhn.Validate(in.s)
		}
	}
}
