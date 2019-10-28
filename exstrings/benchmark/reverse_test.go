package benchmark

import (
	"testing"
	"unicode/utf8"

	"github.com/thinkeridea/go-extend/exbytes"
	"github.com/thinkeridea/go-extend/exstrings"
)

func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseRange(s string) string {
	n := len(s)
	buf := make([]byte, n)
	var start, end int
	for _, r := range s {
		l := utf8.RuneLen(r)
		n -= l
		end = start + l
		copy(buf[n:], s[start:end])
		start = end
	}

	return exbytes.ToString(buf)
}

func ReverseUTF8DecodeRuneInString(s string) string {
	n := len(s)
	buf := make([]byte, n)
	var start, size, end int
	for i := 0; i < len(s[start:]); {
		_, size = utf8.DecodeRuneInString(s[start:])
		n -= size
		end = start + size
		copy(buf[n:], s[start:end])
		start = end
	}

	return exbytes.ToString(buf)
}

func BenchmarkReverseRunes(b *testing.B) {
	s := "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"
	for i := 0; i < b.N; i++ {
		ReverseRunes(s)
	}
}

func BenchmarkReverseRange(b *testing.B) {
	s := "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"
	for i := 0; i < b.N; i++ {
		ReverseRange(s)
	}
}

func BenchmarkReverseUTF8DecodeRuneInString(b *testing.B) {
	s := "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"
	for i := 0; i < b.N; i++ {
		ReverseUTF8DecodeRuneInString(s)
	}
}

func BenchmarkExstringsReverse(b *testing.B) {
	s := "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"
	for i := 0; i < b.N; i++ {
		exstrings.Reverse(s)
	}
}
