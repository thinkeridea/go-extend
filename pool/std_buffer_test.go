// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package pool

import (
	"bytes"
	"reflect"
	"testing"
	"unsafe"
)

func TestStdBuffer(t *testing.T) {
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	buf := bytes.NewBuffer(make([]byte, 0, 20))
	buf.Write(data)

	r := make([]byte, 5)
	buf.Read(r)

	x := (*stdBuffer)(unsafe.Pointer(buf))
	if !reflect.DeepEqual(x.b, data) {
		t.Errorf("Buffer.buf(%v) != (%v)", x.b, data)
	}

	if x.off != 5 {
		t.Errorf("Buffer.off(%d) != 5", x.off)
	}
}

func TestBufferLen(t *testing.T) {
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	buf := bytes.NewBuffer(make([]byte, 0, 20))
	buf.Write(data)

	r := make([]byte, 5)
	buf.Read(r)

	n := bufferLen(buf)
	if n != 10 {
		t.Errorf("bufferLen(%d) != 10", n)
	}
}
