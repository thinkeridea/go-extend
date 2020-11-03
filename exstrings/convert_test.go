// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

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
