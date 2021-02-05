// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package pool

import (
	"bytes"
	"testing"
)

func TestBuffer_Get(t *testing.T) {
	p := NewBuffer(64)
	b := p.Get()
	if b.Cap() != 64 {
		t.Errorf("b.Cap():%d != 64", b.Cap())
	}

	if b.Len() != 0 {
		t.Errorf("b.String():%s != xx", b.String())
	}

	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	// 开启 race 时有一定概率导致 Put 被丢弃
	for i := 0; i < 10; i++ {
		b = p.Get()
		if b.Len() != 0 {
			t.Errorf("b.String():%s != xx", b.String())
		}

		if b.Cap() != 64 {
			t.Errorf("b.Cap():%d != 64", b.Cap())
		}
		p.Put(b)
	}
}

func TestBuffer_Put(t *testing.T) {
	p := NewBuffer(64)
	b := p.Get()
	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p.Put(b)

	bufp := p.(*buffer)
	if bufp.release != 1 {
		t.Errorf("release:%d != 1", bufp.release)
	}

	if bufp.calls[0] != 1 {
		t.Errorf("calls[0]:%d != 1", bufp.calls[0])
	}

	var bb *bytes.Buffer
	// 开启 race 时有一定概率导致 Put 被丢弃
	pp := buffBucket[0]
	var n uint32 = 1
	for i := 0; i < 10; i++ {
		v := pp.Get()
		if v == nil {
			p.Put(b)
			n++
			continue
		}

		bb = v.(*bytes.Buffer)
		if bb.String() == "xx" {
			break
		}

		p.Put(b)
		n++
	}

	if bb.String() != "xx" {
		t.Errorf("b1.String():%s != xx", bb.String())
	}

	if bufp.release != n {
		t.Errorf("release:%d != %d", bufp.release, n)
	}

	if bufp.calls[0] != n {
		t.Errorf("calls[0]:%d != %d", bufp.calls[0], n)
	}

	p.Put(bytes.NewBuffer(make([]byte, 1024)))
	if bufp.release != n+1 {
		t.Errorf("release:%d != %d", bufp.release, n+1)
	}

	if bufp.calls[4] != 1 {
		t.Errorf("calls[4]:%d != 1", bufp.calls[0])
	}

	for i := 0; i < bucketSize; i++ {
		if i != 0 && i != 4 && bufp.calls[i] != 0 {
			t.Errorf("calls[%d]:%d != 0", i, bufp.calls[0])
		}
	}
}

func TestBuffer_Calibrate(t *testing.T) {
	p := NewBuffer(64)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			p.Put(bytes.NewBuffer(make([]byte, minSize<<i)))
		}
	}

	bufp := p.(*buffer)
	bufp.release = calibrateCallsThreshold - 10
	for i := 0; i < 11; i++ {
		p.Put(bytes.NewBuffer(make([]byte, minSize<<5)))
	}

	if bufp.index != 5 {
		t.Errorf("index:%d != 5", bufp.index)
	}

	if bufp.release != 0 {
		t.Errorf("release:%d != 0", bufp.release)
	}

	if bufp.calibrating != 0 {
		t.Errorf("calibrating:%d != 0", bufp.calibrating)
	}

	for i := 0; i < bucketSize; i++ {
		if bufp.calls[i] != 0 {
			t.Errorf("calls[%d]:%d != 0", i, bufp.calls[0])
		}
	}

	bufp.calibrating = 1
	bufp.release = calibrateCallsThreshold - 10
	for i := 0; i < 11; i++ {
		p.Put(bytes.NewBuffer(make([]byte, minSize)))
	}

	if bufp.index != 5 {
		t.Errorf("index:%d != 5", bufp.index)
	}

	if bufp.release != 10241 {
		t.Errorf("release:%d != 10241", bufp.release)
	}

	if bufp.calibrating != 1 {
		t.Errorf("calibrating:%d != 1", bufp.calibrating)
	}

	if bufp.calls[0] != 11 {
		t.Errorf("calls[0]:%d != 11", bufp.calls[0])
	}
}

func TestBuffBucketIndexIndex(t *testing.T) {
	var n int
	for i := 0; i < minSizeBits+bucketSize+5; i++ {
		n = 0
		if i > minSizeBits {
			n = i - minSizeBits
		}

		if n >= bucketSize {
			n = bucketSize - 1
		}

		if idx := buffBucketIndex(1 << i); idx != n {
			t.Errorf("index(%d) :%d = %d", 1<<i, idx, n)
		}

		if i > minSizeBits {
			if idx := buffBucketIndex(1<<(i-1) + 1); idx != n {
				t.Errorf("index(%d) :%d != %d", 1<<(i-1)+1, idx, n)
			}
		}
	}
}

func TestGetBuff64(t *testing.T) {
	p1 := GetBuff64()
	p2 := GetBuff64()
	if p1 != p2 {
		t.Errorf("p1:%p != p2:%p", p1, p2)
	}

	b := p1.Get()
	if c := b.Cap(); c != 64 {
		t.Errorf("b.Cap:%d != 64", c)
	}

	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p1.Put(b)
}

func TestGetBuff128(t *testing.T) {
	p1 := GetBuff128()
	p2 := GetBuff128()
	if p1 != p2 {
		t.Errorf("p1:%p != p2:%p", p1, p2)
	}

	b := p1.Get()
	if c := b.Cap(); c != 128 {
		t.Errorf("b.Cap:%d != 128", c)
	}

	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p1.Put(b)
}

func TestGetBuff512(t *testing.T) {
	p1 := GetBuff512()
	p2 := GetBuff512()
	if p1 != p2 {
		t.Errorf("p1:%p != p2:%p", p1, p2)
	}

	b := p1.Get()
	if c := b.Cap(); c != 512 {
		t.Errorf("b.Cap:%d != 512", c)
	}

	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p1.Put(b)
}

func TestGetBuff1024(t *testing.T) {
	p1 := GetBuff1024()
	p2 := GetBuff1024()
	if p1 != p2 {
		t.Errorf("p1:%p != p2:%p", p1, p2)
	}

	b := p1.Get()
	if c := b.Cap(); c != 1024 {
		t.Errorf("b.Cap:%d != 1024", c)
	}

	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p1.Put(b)
}

func TestGetBuff2048(t *testing.T) {
	p1 := GetBuff2048()
	p2 := GetBuff2048()
	if p1 != p2 {
		t.Errorf("p1:%p != p2:%p", p1, p2)
	}

	b := p1.Get()
	if c := b.Cap(); c != 2048 {
		t.Errorf("b.Cap:%d != 2048", c)
	}

	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p1.Put(b)
}

func TestGetBuff4096(t *testing.T) {
	p1 := GetBuff4096()
	p2 := GetBuff4096()
	if p1 != p2 {
		t.Errorf("p1:%p != p2:%p", p1, p2)
	}

	b := p1.Get()
	if c := b.Cap(); c != 4096 {
		t.Errorf("b.Cap:%d != 4096", c)
	}

	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p1.Put(b)
}

func TestGetBuff8192(t *testing.T) {
	p1 := GetBuff8192()
	p2 := GetBuff8192()
	if p1 != p2 {
		t.Errorf("p1:%p != p2:%p", p1, p2)
	}

	b := p1.Get()
	if c := b.Cap(); c != 8192 {
		t.Errorf("b.Cap:%d != 8192", c)
	}

	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p1.Put(b)
}

func TestPool_Get(t *testing.T) {
	p := GetBuff64()

	b := p.Get()
	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p.Put(b)

	b = p.Get()
	if b.Len() != 0 {
		t.Errorf("b.Len():%d != 0", b.Len())
	}
}

func TestPool_Put(t *testing.T) {
	p := (*bufferPool)(newBufferPool(64))
	b := p.Get()
	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p.Put(b)

	var bb *bytes.Buffer
	// 开启 race 时有一定概率导致 Put 被丢弃
	pp := p.Pool
	for i := 0; i < 10; i++ {
		bb = pp.Get().(*bytes.Buffer)
		if bb.String() == "xx" {
			break
		}

		p.Put(b)
	}

	if bb.String() != "xx" {
		t.Errorf("b1.String():%s != xx", bb.String())
	}
}
