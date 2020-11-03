// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package pool

import (
	"bytes"
	"sync"
	"testing"
)

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
	p := (*pool)(newBufferPool(64))
	b := p.Get()
	b.WriteString("xx")
	if b.String() != "xx" {
		t.Errorf("b.String():%s != xx", b.String())
	}

	p.Put(b)

	var bb *bytes.Buffer
	// 开启 race 时有一定概率导致 Put 被丢弃
	pp := (*sync.Pool)(p)
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
