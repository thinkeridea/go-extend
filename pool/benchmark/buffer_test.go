// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package benchmark

import (
	"bytes"
	"container/ring"
	"math/rand"
	"sync"
	"testing"

	"github.com/thinkeridea/go-extend/pool"
)

var bufferData = ring.New(256)

func init() {
	data := bufferData
	for i := 0; i < 256; i++ {
		data.Value = make([]byte, rand.Intn(1<<16))
		data = data.Next()
	}
}

func BenchmarkBufferPool(b *testing.B) {
	buff := pool.NewBuffer(64)
	p := [20]*bytes.Buffer{}
	data := bufferData
	for i := 0; i < b.N; i++ {
		data.Next()
		bf := buff.Get()
		bf.Write(data.Value.([]byte))

		idx := i % 20
		if v := p[idx]; v != nil {
			buff.Put(v)
		}

		p[idx] = bf
	}
}

func BenchmarkBufferSyncPool(b *testing.B) {
	buff := sync.Pool{New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 64))
	}}

	p := [20]*bytes.Buffer{}
	data := bufferData
	for i := 0; i < b.N; i++ {
		data.Next()
		bf := buff.Get().(*bytes.Buffer)
		bf.Write(data.Value.([]byte))

		idx := i % 20
		if v := p[idx]; v != nil {
			buff.Put(v)
		}

		p[idx] = bf
	}
}

// 用来测试在容量没有变化的情况下与原始方式的性能差异
func BenchmarkBufferFixedSizePool(b *testing.B) {
	buff := pool.NewBuffer(64)
	p := [20]*bytes.Buffer{}
	data := make([]byte, 50)
	for i := 0; i < b.N; i++ {
		bf := buff.Get()
		bf.Write(data)

		idx := i % 20
		if v := p[idx]; v != nil {
			buff.Put(v)
		}

		p[idx] = bf
	}
}

func BenchmarkBufferFixedSizeSyncPool(b *testing.B) {
	buff := sync.Pool{New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 64))
	}}

	p := [20]*bytes.Buffer{}
	data := make([]byte, 50)
	for i := 0; i < b.N; i++ {
		bf := buff.Get().(*bytes.Buffer)
		bf.Write(data)

		idx := i % 20
		if v := p[idx]; v != nil {
			buff.Put(v)
		}

		p[idx] = bf
	}
}
