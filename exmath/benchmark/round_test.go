package benchmark

import (
	"testing"

	"github.com/thinkeridea/go-extend/exmath"
)

func BenchmarkRound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = exmath.Round(169543.34596, 0)
	}
}

func BenchmarkRoundDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = exmath.Round(0.15807659924030304, 5)
	}
}

func BenchmarkRoundInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = exmath.Round(169543.34596, -5)
	}
}
