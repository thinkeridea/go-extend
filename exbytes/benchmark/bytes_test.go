// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package benchmark

import (
	"bytes"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/thinkeridea/go-extend/exbytes"
)

var replaces string
var replaceb []byte

func init() {
	replaces = strings.Repeat("A BC", 100)
	replaceb = bytes.Repeat([]byte("A BC"), 100)
}

func UnsafeStringsReplace(s, old, new string, n int) string {
	if old == new || n == 0 {
		return s // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return s // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		w += copy(t[w:], s[start:j])
		w += copy(t[w:], new)
		start = j + len(old)
	}
	w += copy(t[w:], s[start:])
	return exbytes.ToString(t[0:w])
}

func BenchmarkReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exbytes.Replace([]byte(replaces), []byte(" "), []byte(""), -1)
	}
}

func BenchmarkBytesReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes.Replace([]byte(replaces), []byte(" "), []byte(""), -1)
	}
}

func BenchmarkStringsReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Replace(string(replaceb), " ", "", -1)
	}
}

func BenchmarkUnsafeStringsReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnsafeStringsReplace(string(replaceb), " ", "", -1)
	}
}
