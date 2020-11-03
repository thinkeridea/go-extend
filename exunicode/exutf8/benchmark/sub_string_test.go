// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package benchmark

import (
	"testing"
	"unicode/utf8"

	"github.com/thinkeridea/go-extend/exunicode/exutf8"
)

var benchmarkSubString = "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"
var benchmarkSubStringLength = 20

func SubStrRunes(s string, length int) string {
	if utf8.RuneCountInString(s) > length {
		rs := []rune(s)
		return string(rs[:length])
	}

	return s
}

func SubStrRange(s string, length int) string {
	var n, i int
	for i = range s {
		if n == length {
			break
		}

		n++
	}

	return s[:i]
}

func SubStrDecodeRuneInString(s string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[n:])
		n += size
	}

	return s[:n]
}

func SubStrRuneIndexInString(s string, length int) string {
	n, _ := exutf8.RuneIndexInString(s, length)
	return s[:n]
}

func SubStrRuneSubString(s string, length int) string {
	return exutf8.RuneSubString(s, 0, length)
}

func BenchmarkSubStrRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubStrRunes(benchmarkSubString, benchmarkSubStringLength)
	}
}

func BenchmarkSubStrRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubStrRange(benchmarkSubString, benchmarkSubStringLength)
	}
}

func BenchmarkSubStrDecodeRuneInString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubStrDecodeRuneInString(benchmarkSubString, benchmarkSubStringLength)
	}
}

func BenchmarkSubStrRuneIndexInString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubStrRuneIndexInString(benchmarkSubString, benchmarkSubStringLength)
	}
}

func BenchmarkSubStrRuneSubString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubStrRuneSubString(benchmarkSubString, benchmarkSubStringLength)
	}
}
