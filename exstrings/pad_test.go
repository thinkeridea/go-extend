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
	"strings"
	"testing"
)

func TestPad(t *testing.T) {
	s := "hello world"

	for _, v := range []struct {
		pad, expected string
		count, flag   int
	}{
		{"-", s, 10, PadLeft},
		{"-", s, 10, PadRight},
		{"-", s, 10, PadBoth},

		{"-", "--" + s, 13, PadLeft},
		{"-", s + "--", 13, PadRight},
		{"-", "-" + s + "-", 13, PadBoth},
		{"-", "-" + s + "--", 14, PadBoth},

		{"AB", "AB" + s, 13, PadLeft},
		{"AB", "ABA" + s, 14, PadLeft},
		{"AB", s + "AB", 13, PadRight},
		{"AB", s + "ABA", 14, PadRight},
		{"AB", "A" + s + "A", 13, PadBoth},
		{"AB", "AB" + s + "AB", 15, PadBoth},
		{"AB", "AB" + s + "ABA", 16, PadBoth},
	} {
		if actual := Pad(s, v.pad, v.count, v.flag); actual != v.expected {
			t.Fatalf("s=%s pad=%s count=%d, flag=%d actual:%s!=expected:%s", s, v.pad, v.count, v.flag, actual, v.expected)
		}
	}
}

func TestLeftPad(t *testing.T) {
	s := "hello world"

	for _, v := range []struct {
		pad, expected string
		count         int
	}{
		{"-", s, 10},
		{"-", "--" + s, 13},
		{"AB", "AB" + s, 13},
		{"AB", "ABA" + s, 14},
	} {
		if actual := LeftPad(s, v.pad, v.count); actual != v.expected {
			t.Fatalf("s=%s pad=%s count=%d, actual:%s!=expected:%s", s, v.pad, v.count, actual, v.expected)
		}
	}
}

func TestRightPad(t *testing.T) {
	s := "hello world"

	for _, v := range []struct {
		pad, expected string
		count         int
	}{
		{"-", s, 10},
		{"-", s + "--", 13},
		{"AB", s + "AB", 13},
		{"AB", s + "ABA", 14},
	} {
		if actual := RightPad(s, v.pad, v.count); actual != v.expected {
			t.Fatalf("s=%s pad=%s count=%d, actual:%s!=expected:%s", s, v.pad, v.count, actual, v.expected)
		}
	}
}

func TestBothPad(t *testing.T) {
	s := "hello world"

	for _, v := range []struct {
		pad, expected string
		count         int
	}{
		{"-", s, 10},
		{"-", "-" + s + "-", 13},
		{"-", "-" + s + "--", 14},
		{"AB", "A" + s + "A", 13},
		{"AB", "AB" + s + "AB", 15},
		{"AB", "AB" + s + "ABA", 16},
	} {
		if actual := BothPad(s, v.pad, v.count); actual != v.expected {
			t.Fatalf("s=%s pad=%s count=%d, actual:%s!=expected:%s", s, v.pad, v.count, actual, v.expected)
		}
	}
}

func TestUnsafePad(t *testing.T) {
	s := "hello world"

	for _, v := range []struct {
		pad, expected string
		count, flag   int
	}{
		{"-", s, 10, PadLeft},
		{"-", s, 10, PadRight},
		{"-", s, 10, PadBoth},

		{"-", "--" + s, 13, PadLeft},
		{"-", s + "--", 13, PadRight},
		{"-", "-" + s + "-", 13, PadBoth},
		{"-", "-" + s + "--", 14, PadBoth},

		{"AB", "AB" + s, 13, PadLeft},
		{"AB", "ABA" + s, 14, PadLeft},
		{"AB", s + "AB", 13, PadRight},
		{"AB", s + "ABA", 14, PadRight},
		{"AB", "A" + s + "A", 13, PadBoth},
		{"AB", "AB" + s + "AB", 15, PadBoth},
		{"AB", "AB" + s + "ABA", 16, PadBoth},
	} {
		if actual := UnsafePad(s, v.pad, v.count, v.flag); actual != v.expected {
			t.Fatalf("s=%s pad=%s count=%d, flag=%d actual:%s!=expected:%s", s, v.pad, v.count, v.flag, actual, v.expected)
		}
	}
}

func TestUnsafeLeftPad(t *testing.T) {
	s := "hello world"

	for _, v := range []struct {
		pad, expected string
		count         int
	}{
		{"-", s, 10},
		{"-", "--" + s, 13},
		{"AB", "AB" + s, 13},
		{"AB", "ABA" + s, 14},
	} {
		if actual := UnsafeLeftPad(s, v.pad, v.count); actual != v.expected {
			t.Fatalf("s=%s pad=%s count=%d, actual:%s!=expected:%s", s, v.pad, v.count, actual, v.expected)
		}
	}
}

func TestUnsafeRightPad(t *testing.T) {
	s := "hello world"

	for _, v := range []struct {
		pad, expected string
		count         int
	}{
		{"-", s, 10},
		{"-", s + "--", 13},
		{"AB", s + "AB", 13},
		{"AB", s + "ABA", 14},
	} {
		if actual := UnsafeRightPad(s, v.pad, v.count); actual != v.expected {
			t.Fatalf("s=%s pad=%s count=%d, actual:%s!=expected:%s", s, v.pad, v.count, actual, v.expected)
		}
	}
}

func TestUnsafeBothPad(t *testing.T) {
	s := "hello world"

	for _, v := range []struct {
		pad, expected string
		count         int
	}{
		{"-", s, 10},
		{"-", "-" + s + "-", 13},
		{"-", "-" + s + "--", 14},
		{"AB", "A" + s + "A", 13},
		{"AB", "AB" + s + "AB", 15},
		{"AB", "AB" + s + "ABA", 16},
	} {
		if actual := UnsafeBothPad(s, v.pad, v.count); actual != v.expected {
			t.Fatalf("s=%s pad=%s count=%d, actual:%s!=expected:%s", s, v.pad, v.count, actual, v.expected)
		}
	}
}

func BenchmarkLeftPad(b *testing.B) {
	s := strings.Repeat("A", 1000)
	pad := strings.Repeat("B", 10)
	for i := 0; i < b.N; i++ {
		LeftPad(s, pad, 100000)
	}
}

func BenchmarkUnsafeLeftPad(b *testing.B) {
	s := strings.Repeat("A", 1000)
	pad := strings.Repeat("B", 10)
	for i := 0; i < b.N; i++ {
		UnsafeLeftPad(s, pad, 100000)
	}
}
