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
	"strings"
	"testing"
	"unsafe"
)

func TestUnsafeToBytes(t *testing.T) {
	s := "hello word"
	b := UnsafeToBytes("hello word")

	bptr := (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data
	sptr := (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	if bptr != sptr {
		t.Fatalf("bptr=%x sptr=%x", bptr, sptr)
	}

	if string(b) != s {
		t.Fatalf("string(b)=%s s=%s", string(b), s)
	}

	s = strings.Repeat("A", 3)
	b = UnsafeToBytes(s)
	b[0] = 'A'
	b[1] = 'B'
	b[2] = 'C'

	bptr = (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data
	sptr = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	if bptr != sptr {
		t.Fatalf("bptr=%x sptr=%x", bptr, sptr)
	}

	if string(b) != s {
		t.Fatalf("string(b)=%s s=%s", string(b), s)
	}
}

func TestBytes(t *testing.T) {
	for _, in := range []string{
		"abcd",
		"☺☻☹",
		"☺☻☹",
	} {
		if p := Bytes(in); !reflect.DeepEqual(p, []byte(in)) {
			t.Fatalf("Bytes(%s)=%q want %q", in, p, []byte(in))
		}
	}
}

func BenchmarkUnsafeToBytes(b *testing.B) {
	str := strings.Repeat("abc", 128)
	for i := 0; i < b.N; i++ {
		UnsafeToBytes(str)
	}
}
