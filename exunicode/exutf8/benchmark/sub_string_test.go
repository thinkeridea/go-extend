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
