// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package pool

import (
	"bytes"
	"sync"
)

var (
	buff64   *sync.Pool
	buff128  *sync.Pool
	buff512  *sync.Pool
	buff1024 *sync.Pool
	buff2048 *sync.Pool
	buff4096 *sync.Pool
	buff8192 *sync.Pool

	buff64One   sync.Once
	buff128One  sync.Once
	buff512One  sync.Once
	buff1024One sync.Once
	buff2048One sync.Once
	buff4096One sync.Once
	buff8192One sync.Once
)

type pool sync.Pool

// BufferPool bytes.Buffer 的 sync.Pool 接口
// 可以直接 Get *bytes.Buffer 并 Reset Buffer
type BufferPool interface {

	// Get 从 Pool 中获取一个 *bytes.Buffer 实例, 该实例已经被 Reset
	Get() *bytes.Buffer
	// Put 把 *bytes.Buffer 放回 Pool 中
	Put(*bytes.Buffer)
}

func newBufferPool(size int) *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, size))
		},
	}
}

// GetBuff64 获取一个初始容量为 64 的 *bytes.Buffer Pool
func GetBuff64() BufferPool {
	buff64One.Do(func() {
		buff64 = newBufferPool(64)
	})

	return (*pool)(buff64)
}

// GetBuff128 获取一个初始容量为 128 的 *bytes.Buffer Pool
func GetBuff128() BufferPool {
	buff128One.Do(func() {
		buff128 = newBufferPool(128)
	})

	return (*pool)(buff128)
}

// GetBuff512 获取一个初始容量为 512 的 *bytes.Buffer Pool
func GetBuff512() BufferPool {
	buff512One.Do(func() {
		buff512 = newBufferPool(512)
	})

	return (*pool)(buff512)
}

// GetBuff1024 获取一个初始容量为 1024 的 *bytes.Buffer Pool
func GetBuff1024() BufferPool {
	buff1024One.Do(func() {
		buff1024 = newBufferPool(1024)
	})

	return (*pool)(buff1024)
}

// GetBuff2048 获取一个初始容量为 2048 的 *bytes.Buffer Pool
func GetBuff2048() BufferPool {
	buff2048One.Do(func() {
		buff2048 = newBufferPool(2048)
	})

	return (*pool)(buff2048)
}

// GetBuff4096 获取一个初始容量为 4096 的 *bytes.Buffer Pool
func GetBuff4096() BufferPool {
	buff4096One.Do(func() {
		buff4096 = newBufferPool(4096)
	})

	return (*pool)(buff4096)
}

// GetBuff8192 获取一个初始容量为 8192 的 *bytes.Buffer Pool
func GetBuff8192() BufferPool {
	buff8192One.Do(func() {
		buff8192 = newBufferPool(8192)
	})

	return (*pool)(buff8192)
}

// Get 从 Pool 中获取一个 *bytes.Buffer 实例, 该实例已经被 Reset
func (p *pool) Get() *bytes.Buffer {
	b := (*sync.Pool)(p).Get().(*bytes.Buffer)
	b.Reset()
	return b
}

// Put 把 *bytes.Buffer 放回 Pool 中
func (p *pool) Put(b *bytes.Buffer) {
	(*sync.Pool)(p).Put(b)
}
