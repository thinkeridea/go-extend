// Copyright (C) 2019  Qi Yin <qiyin@thinkeridea.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
