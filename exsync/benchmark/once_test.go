// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package benchmark

import (
	"sync"
	"testing"
	"unsafe"

	"github.com/thinkeridea/go-extend/exsync"
)

type one int

func (o *one) Increment() {
	*o++
}

func BenchmarkSyncOnce(b *testing.B) {
	var once sync.Once
	f := func() {
		_ = new(one)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(f)
		}
	})
}

func BenchmarkOnce(b *testing.B) {
	var once exsync.Once
	f := func() interface{} { return new(one) }
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = once.Do(f).(*one)
		}
	})
}

func BenchmarkOncePointer(b *testing.B) {
	var once exsync.OncePointer
	f := func() unsafe.Pointer { return unsafe.Pointer(new(one)) }
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = (*one)(once.Do(f))
		}
	})
}
