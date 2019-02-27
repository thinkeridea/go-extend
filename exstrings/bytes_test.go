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

func TestRepeatToBytes(t *testing.T) {
	for _, tt := range RepeatTests {
		a := RepeatToBytes(tt.in, tt.count)
		if !equal("RepeatToBytes(s)", string(a), tt.out, t) {
			t.Errorf("RepeatToBytes(%v, %d) = %v; want %v", tt.in, tt.count, a, tt.out)
			continue
		}
	}
}

func repeatToBytesPanicRecover(s string, count int) (err error) {
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

	RepeatToBytes(s, count)

	return
}

// See Issue golang.org/issue/16237
func TestRepeatToBytesCatchesOverflow(t *testing.T) {
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
		err := repeatToBytesPanicRecover(tt.s, tt.count)
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

func TestJoinToBytes(t *testing.T) {
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
		a := JoinToBytes(v.in, v.sep)
		if string(a) != v.out {
			t.Errorf("Join(%v, %s) = %v; want %v", v.in, v.sep, v.out, string(a))
			continue
		}
	}
}
