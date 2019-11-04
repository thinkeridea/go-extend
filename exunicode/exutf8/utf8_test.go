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

package exutf8

import (
	"testing"
)

func TestRuneIndex(t *testing.T) {
	for _, tt := range []struct {
		in     string
		length int
		out    int
		ok     bool
	}{
		{"abcd", 3, 3, true},
		{"☺☻☹", 2, 6, true},
		{"☺☻☹", 3, 9, true},
		{"1,2,3,4", 5, 5, true},
		{"\xe2\x00", 2, 2, true},
		{"\xe2\x80", 1, 1, true},
		{"\xe2\x80", 2, 2, true},
		{"a\xe2\x80", 5, 3, false},
		{"golang", 5, 5, true},
		{"Go 语言", 4, 6, true},
		{"12345", 0, 0, true},
		{"12345", -1, 0, true},
	} {
		if out, ok := RuneIndex([]byte(tt.in), tt.length); out != tt.out || ok != tt.ok {
			t.Errorf("RuneIndex(%q, %d) = %d, %v, want %d, %v", tt.in, tt.length, out, ok, tt.out, tt.ok)
		}

		if out, ok := RuneIndexInString(tt.in, tt.length); out != tt.out || ok != tt.ok {
			t.Errorf("RuneIndexInString(%q, %d) = %d, %v, want %d, %v", tt.in, tt.length, out, ok, tt.out, tt.ok)
		}
	}
}

func TestRuneSub(t *testing.T) {
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
		if out := RuneSub([]byte(tt.in), tt.start, tt.length); string(out) != tt.out {
			t.Errorf("RuneSub(%q, %d, %d) = %s, want %s", tt.in, tt.start, tt.length, out, tt.out)
		}

		if out := RuneSubString(tt.in, tt.start, tt.length); out != tt.out {
			t.Errorf("RuneSubstring(%q, %d, %d) = %s, want %s", tt.in, tt.start, tt.length, out, tt.out)
		}
	}
}
