// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package pool

import (
	"bytes"
	"math/bits"
	"sort"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/thinkeridea/go-extend/exsync"
)

const (
	bucketSize              = 22
	minSizeBits             = 6
	minSize                 = 1 << minSizeBits                    // 2^6=64
	maxSize                 = 1 << (minSizeBits + bucketSize - 1) // 2^28=256MB
	calibrateCallsThreshold = 10240
)

var (
	buffBucket = [bucketSize]sync.Pool{}
)

// BufferPool bytes.Buffer 的 sync.Pool 接口
// 可以直接 Get *bytes.Buffer 并 Reset Buffer
type BufferPool interface {

	// Get 从 Pool 中获取一个 *bytes.Buffer 实例, 该实例已经被 Reset
	Get() *bytes.Buffer
	// Put 把 *bytes.Buffer 放回 Pool 中
	Put(*bytes.Buffer)
}

type buffer struct {
	index       uint32
	calibrating uint32
	release     uint32
	calls       [bucketSize]uint32
}

func NewBuffer(size int) BufferPool {
	b := &buffer{}
	b.index = uint32(buffBucketIndex(size))
	return b
}

func (p *buffer) Get() *bytes.Buffer {
	idx := atomic.LoadUint32(&p.index)
	v := buffBucket[idx].Get()
	if v != nil {
		b := v.(*bytes.Buffer)
		b.Reset()
		return b
	}

	return bytes.NewBuffer(make([]byte, 0, minSize<<idx))
}

func (p *buffer) Put(b *bytes.Buffer) {
	if b.Cap() <= maxSize {
		buffBucket[buffBucketIndex(b.Cap())].Put(b)
	}

	atomic.AddUint32(&p.calls[buffBucketIndex(bufferLen(b))], 1)
	if atomic.AddUint32(&p.release, 1) > calibrateCallsThreshold {
		p.calibrate()
	}
}

func (p *buffer) calibrate() {
	if !atomic.CompareAndSwapUint32(&p.calibrating, 0, 1) {
		return
	}

	var callSize [bucketSize]uint64
	for i := range callSize {
		callSize[i] = uint64(atomic.SwapUint32(&p.calls[i], 0))<<32 | minSize<<i
	}

	sort.Slice(callSize[:], func(i, j int) bool {
		return callSize[i] > callSize[j]
	})

	atomic.SwapUint32(&p.index, uint32(buffBucketIndex(int(callSize[0]<<32>>32))))
	atomic.StoreUint32(&p.release, 0)
	atomic.SwapUint32(&p.calibrating, 0)
}

func buffBucketIndex(n int) int {
	if n == 0 {
		return 0
	}

	idx := bits.Len32(uint32((n - 1) >> minSizeBits))
	if idx >= bucketSize {
		idx = bucketSize - 1
	}
	return idx
}

var (
	buff64   exsync.OncePointer
	buff128  exsync.OncePointer
	buff512  exsync.OncePointer
	buff1024 exsync.OncePointer
	buff2048 exsync.OncePointer
	buff4096 exsync.OncePointer
	buff8192 exsync.OncePointer
)

type bufferPool struct {
	sync.Pool
}

func newBufferPool(size int) unsafe.Pointer {
	return unsafe.Pointer(&bufferPool{
		Pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, size))
			},
		},
	})
}

// GetBuff64 获取一个初始容量为 64 的 *bytes.Buffer Pool
func GetBuff64() BufferPool {
	return (*bufferPool)(buff64.Do(func() unsafe.Pointer {
		return newBufferPool(64)
	}))
}

// GetBuff128 获取一个初始容量为 128 的 *bytes.Buffer Pool
func GetBuff128() BufferPool {
	return (*bufferPool)(buff128.Do(func() unsafe.Pointer {
		return newBufferPool(128)
	}))
}

// GetBuff512 获取一个初始容量为 512 的 *bytes.Buffer Pool
func GetBuff512() BufferPool {
	return (*bufferPool)(buff512.Do(func() unsafe.Pointer {
		return newBufferPool(512)
	}))
}

// GetBuff1024 获取一个初始容量为 1024 的 *bytes.Buffer Pool
func GetBuff1024() BufferPool {
	return (*bufferPool)(buff1024.Do(func() unsafe.Pointer {
		return newBufferPool(1024)
	}))
}

// GetBuff2048 获取一个初始容量为 2048 的 *bytes.Buffer Pool
func GetBuff2048() BufferPool {
	return (*bufferPool)(buff2048.Do(func() unsafe.Pointer {
		return newBufferPool(2048)
	}))
}

// GetBuff4096 获取一个初始容量为 4096 的 *bytes.Buffer Pool
func GetBuff4096() BufferPool {
	return (*bufferPool)(buff4096.Do(func() unsafe.Pointer {
		return newBufferPool(4096)
	}))
}

// GetBuff8192 获取一个初始容量为 8192 的 *bytes.Buffer Pool
func GetBuff8192() BufferPool {
	return (*bufferPool)(buff8192.Do(func() unsafe.Pointer {
		return newBufferPool(8192)
	}))
}

// Get 从 Pool 中获取一个 *bytes.Buffer 实例, 该实例已经被 Reset
func (p *bufferPool) Get() *bytes.Buffer {
	b := p.Pool.Get().(*bytes.Buffer)
	b.Reset()
	return b
}

// Put 把 *bytes.Buffer 放回 Pool 中
func (p *bufferPool) Put(b *bytes.Buffer) {
	p.Pool.Put(b)
}
