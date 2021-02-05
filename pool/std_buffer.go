// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package pool

import (
	"bytes"
	"unsafe"
)

// stdBuffer 是 bytes.Buffer 的结构引用，这需要对各个go的版本进行测试，保证该结构的准确性
type stdBuffer struct {
	b   []byte
	off int
}

// bufferLen 获取buf的实际使用长度， 默认使用 cap 作为降级方案
func bufferLen(b *bytes.Buffer) int {
	x := (*stdBuffer)(unsafe.Pointer(b))
	if x.off+b.Len() == len(x.b) {
		return len(x.b)
	}

	return b.Cap()
}
