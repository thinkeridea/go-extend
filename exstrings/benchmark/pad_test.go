// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package benchmark

import (
	"strings"
	"testing"

	"github.com/thinkeridea/go-extend/exstrings"
)

func BenchmarkLeftPad(b *testing.B) {
	s := strings.Repeat("A", 1000)
	pad := strings.Repeat("B", 10)
	for i := 0; i < b.N; i++ {
		exstrings.LeftPad(s, pad, 100000)
	}
}

func BenchmarkUnsafeLeftPad(b *testing.B) {
	s := strings.Repeat("A", 1000)
	pad := strings.Repeat("B", 10)
	for i := 0; i < b.N; i++ {
		exstrings.UnsafeLeftPad(s, pad, 100000)
	}
}
