package benchmarks

import (
	"testing"

	"github.com/ferdypruis/go-luhn"
)

func BenchmarkChecksum_ferdypruis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			checksum, _ = luhn.Checksum(in.s)
		}
	}
}

func BenchmarkValid_ferdypruis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			valid = luhn.Valid(in.s)
		}
	}
}


