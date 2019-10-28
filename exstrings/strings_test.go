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

package exstrings

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
	"unsafe"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestReverseASCII(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseASCII(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestUnsafeReverseASCII(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{fmt.Sprint("Hello, world"), "dlrow ,olleH"},
		{fmt.Sprint(""), ""},
	}
	for _, c := range cases {
		got := UnsafeReverseASCII(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}

		if c.in != c.want {
			t.Errorf("c.in == %q, want %q", c.in, c.want)
		}
	}
}

func TestReplace(t *testing.T) {
	var ReplaceTests = []struct {
		in       string
		old, new string
		n        int
		out      string
	}{
		{"hello", "l", "L", 0, "hello"},
		{"hello", "l", "L", -1, "heLLo"},
		{"hello", "x", "X", -1, "hello"},
		{"", "x", "X", -1, ""},
		{"radar", "r", "<r>", -1, "<r>ada<r>"},
		{"", "", "<>", -1, "<>"},
		{"banana", "a", "<>", -1, "b<>n<>n<>"},
		{"banana", "a", "<>", 1, "b<>nana"},
		{"banana", "a", "<>", 1000, "b<>n<>n<>"},
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
	for _, tt := range ReplaceTests {
		if s := Replace(tt.in, tt.old, tt.new, tt.n); s != tt.out {
			t.Errorf("Replace(%q, %q, %q, %d) = %q, want %q", tt.in, tt.old, tt.new, tt.n, s, tt.out)
		}
	}
}

var RepeatTests = []struct {
	in, out string
	count   int
}{
	{"", "", 0},
	{"", "", 1},
	{"", "", 2},
	{"-", "", 0},
	{"-", "-", 1},
	{"-", "----------", 10},
	{"abc ", "abc abc abc ", 3},
}

func TestRepeat(t *testing.T) {
	for _, tt := range RepeatTests {
		a := Repeat(tt.in, tt.count)
		if !equal("Repeat(s)", a, tt.out, t) {
			t.Errorf("Repeat(%v, %d) = %v; want %v", tt.in, tt.count, a, tt.out)
			continue
		}
	}
}

func repeatPanicRecover(s string, count int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%s", v)
			}
		}
	}()

	Repeat(s, count)

	return
}

// See Issue golang.org/issue/16237
func TestRepeatCatchesOverflow(t *testing.T) {
	tests := [...]struct {
		s      string
		count  int
		errStr string
	}{
		0: {"--", -2147483647, "negative"},
		1: {"", int(^uint(0) >> 1), ""},
		2: {"-", 10, ""},
		3: {"gopher", 0, ""},
		4: {"-", -1, "negative"},
		5: {"--", -102, "negative"},
		6: {string(make([]byte, 255)), int((^uint(0))/255 + 1), "overflow"},
	}

	for i, tt := range tests {
		err := repeatPanicRecover(tt.s, tt.count)
		if tt.errStr == "" {
			if err != nil {
				t.Errorf("#%d panicked %v", i, err)
			}
			continue
		}

		if err == nil || !strings.Contains(err.Error(), tt.errStr) {
			t.Errorf("#%d expected %q got %q", i, tt.errStr, err)
		}
	}
}

func equal(m string, s1, s2 string, t *testing.T) bool {
	if s1 == s2 {
		return true
	}
	e1 := strings.Split(s1, "")
	e2 := strings.Split(s2, "")
	for i, c1 := range e1 {
		if i >= len(e2) {
			break
		}
		r1, _ := utf8.DecodeRuneInString(c1)
		r2, _ := utf8.DecodeRuneInString(e2[i])
		if r1 != r2 {
			t.Errorf("%s diff at %d: U+%04X U+%04X", m, i, r1, r2)
		}
	}
	return false
}

func TestJoin(t *testing.T) {
	for _, v := range []struct {
		in       []string
		out, sep string
	}{
		{[]string{}, "", "-"},
		{[]string{"a"}, "a", "-"},
		{[]string{"a", "b"}, "a-b", "-"},
		{[]string{"a", "b", "c"}, "a-b-c", "-"},
		{[]string{"a", "b", "c", "d"}, "a-b-c-d", "-"},
		{[]string{"a", "b", "c", "d", "e"}, "a-b-c-d-e", "-"},
		{[]string{"a", "b", "c", "d", "e", "f"}, "a-b-c-d-e-f", "-"},
	} {
		a := Join(v.in, v.sep)
		if a != v.out {
			t.Errorf("Join(%v, %s) = %v; want %v", v.in, v.sep, v.out, a)
			continue
		}
	}
}

func TestCopy(t *testing.T) {
	testString := "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"

	var s string
	s = Copy(testString[:119])
	if s != "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。" {
		t.Errorf("Copy(testString[:119]) = %v; want %v", s, "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。")
	}

	a := (*reflect.StringHeader)(unsafe.Pointer(&testString))
	b := (*reflect.StringHeader)(unsafe.Pointer(&s))

	if a.Data == b.Data {
		t.Errorf("testString pointer address == s pointer address")
	}
}

func TestSubString(t *testing.T) {
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
		if out := SubString(tt.in, tt.start, tt.length); out != tt.out {
			t.Errorf("RuneSubstring(%q, %d, %d) = %s, want %s", tt.in, tt.start, tt.length, out, tt.out)
		}
	}
}
