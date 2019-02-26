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
	"reflect"
	"testing"
)

func TestReplaceToBytes(t *testing.T) {
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
		if s := ReplaceToBytes(tt.in, tt.old, tt.new, tt.n); !reflect.DeepEqual(s, []byte(tt.out)) {
			t.Errorf("Replace(%q, %q, %q, %d) = %q, want %q", tt.in, tt.old, tt.new, tt.n, string(s), tt.out)
		}
	}
}

func TestUnsafeReplaceToBytes(t *testing.T) {
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
		if s := UnsafeReplaceToBytes(tt.in, tt.old, tt.new, tt.n); !reflect.DeepEqual(s, []byte(tt.out)) {
			t.Errorf("Replace(%q, %q, %q, %d) = %q, want %q", tt.in, tt.old, tt.new, tt.n, string(s), tt.out)
		}
	}
}

