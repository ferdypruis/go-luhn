package benchmarks

import (
	"testing"

	luhnmod10 "github.com/luhnmod10/go"
)

func BenchmarkChecksum_luhnmod10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			checksumInt = luhnmod10.CheckDigit(in.s)
		}
	}
}

func BenchmarkValid_luhnmod10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			valid = luhnmod10.Valid(in.s)
		}
	}
}
