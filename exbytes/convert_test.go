// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package exbytes

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestToString(t *testing.T) {
	b := []byte("hello word")
	s := ToString(b)

	bptr := (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data
	sptr := (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	if bptr != sptr {
		t.Fatalf("bptr=%x sptr=%x", bptr, sptr)
	}

	if string(b) != s {
		t.Fatalf("string(b)=%s s=%s", string(b), s)
	}

	b[0] = 'a'
	if string(b) != s {
		t.Fatalf("string(b)=%s s=%s", string(b), s)
	}
}
