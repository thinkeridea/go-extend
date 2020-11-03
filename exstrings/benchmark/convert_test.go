// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package benchmark

import (
	"testing"

	"github.com/thinkeridea/go-extend/exstrings"
)

func BenchmarkStandardLibraryStringToBytes(b *testing.B) {
	s := "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"
	for i := 0; i < b.N; i++ {
		p := []byte(s)
		_ = p

	}
}

func BenchmarkExstringsStringToBytes(b *testing.B) {
	s := "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"
	for i := 0; i < b.N; i++ {
		p := exstrings.Bytes(s)
		_ = p
	}
}
