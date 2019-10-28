// Copyright (C) 2018  Qi Yin <qiyin@thinkeridea.com>
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

package exbytes

import (
	"reflect"
	"testing"
)

type ReplaceTest struct {
	in       string
	old, new string
	n        int
	out      string
}

var ReplaceTests = []ReplaceTest{
	{"hello", "l", "L", 0, "hello"},
	{"hello", "l", "L", -1, "heLLo"},
	{"hello", "x", "X", -1, "hello"},
	{"", "x", "X", -1, ""},
	{"radarad", "rad", "<r>", -1, "<r>a<r>"},
	{"", "", "<>", -1, "<>"},
	{"banana", "a", "<", -1, "b<n<n<"},
	{"banana", "a", ">", 1, "b>nana"},
	{"banana", "a", "<", 1000, "b<n<n<"},
	{"banana", "an", "<>", -1, "b<><>a"},
	{"banana", "ana", "<>", -1, "b<>na"},
	{"banana", "", "<>", -1, "<>b<>a<>n<>a<>n<>a<>"},
	{"banana", "", "<>", 10, "<>b<>a<>n<>a<>n<>a<>"},
	{"banana", "", "<>", 6, "<>b<>a<>n<>a<>n<>a"},
	{"banana", "", "<>", 5, "<>b<>a<>n<>a<>na"},
	{"banana", "", "<>", 1, "<>banana"},
	{"banana", "a", "a", -1, "banana"},
	{"banana", "a", "a", 1, "banana"},
	{"☺☻☹", "", "<>", -1, "<>☺<>☻<>☹<>"},
}

func TestReplace(t *testing.T) {
	for _, tt := range ReplaceTests {
		in := append([]byte(tt.in), "<spare>"...)
		in = in[:len(tt.in)]
		out := Replace(in, []byte(tt.old), []byte(tt.new), tt.n)
		if s := string(out); s != tt.out {
			t.Errorf("Replace(%q, %q, %q, %d) = %q, want %q", tt.in, tt.old, tt.new, tt.n, s, tt.out)
		}
	}
}

func TestReverse(t *testing.T) {
	s := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Reverse(s)

	if !reflect.DeepEqual(s, []byte{9, 8, 7, 6, 5, 4, 3, 2, 1}) {
		t.Errorf("Reverse(%q) = %q, want %q", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}, s, []byte{9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}

func TestSub(t *testing.T) {
	s := "Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言"
	for _, tt := range []struct {
		start, length int
		in, out       string
	}{
		{-1, 0, s, "言"},
		{-2, 0, s, "语言"},
		{-3, 1, s, "程"},
		{-3, -1, s, "程语"},
		{0, -1, s, "Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语"},
		{2, -1, s, "（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语"},
		{52, 52, s, ""},
		{3, 5, s, "又称Gol"},
		{1, 0, s, "o（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言"},
		{0, 10000, "", ""},
		{-10000, 10000, s, ""},
		{0, -10000, s, ""},
	} {
		if out := Sub([]byte(tt.in), tt.start, tt.length); string(out) != tt.out {
			t.Errorf("RuneSub(%q, %d, %d) = %s, want %s", tt.in, tt.start, tt.length, out, tt.out)
		}
	}
}
