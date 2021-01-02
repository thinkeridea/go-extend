package benchmark

import (
	"testing"

	"github.com/thinkeridea/go-extend/exmath"
)

func BenchmarkRound(b *testing.B) {
	f := 0.15807659924030304
	for i := 0; i < b.N; i++ {
		_ = exmath.Round(f, 5)
	}
}
