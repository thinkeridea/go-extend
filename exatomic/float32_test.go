// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package exatomic

import (
	"testing"
)

// Tests of correct behavior, without contention.
// (Does the function work as advertised?)
//
// Test that the Add functions add correctly.
// Test that the CompareAndSwap functions actually
// do the comparison and the swap correctly.
//
// The loop over power-of-two values is meant to
// ensure that the operations apply to the full word size.
// The struct fields x.before and x.after check that the
// operations do not extend past the full word size.

const (
	magic32 = 0xdedbeef
)

func TestSwapFloat32(t *testing.T) {
	var x struct {
		before float32
		i      float32
		after  float32
	}
	x.before = magic32
	x.after = magic32
	var j float32
	for delta := float32(0.1); delta+delta > delta; delta += delta {
		k := SwapFloat32(&x.i, delta)
		if x.i != delta || k != j {
			t.Fatalf("delta=%f i=%f j=%f k=%f", delta, x.i, j, k)
		}
		j = delta
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#v _ %#v != %#v _ %#v", x.before, x.after, float32(magic32), float32(magic32))
	}
}

func TestCompareAndSwapFloat32(t *testing.T) {
	var x struct {
		before float32
		i      float32
		after  float32
	}
	x.before = magic32
	x.after = magic32
	for val := float32(0.01); val+val > val; val += val {
		x.i = val
		if !CompareAndSwapFloat32(&x.i, val, val+1) {
			t.Fatalf("should have swapped %#v %#v", val, val+1)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#v val+1=%#v", x.i, val+1)
		}

		x.i = val + 1

		// float 在运算时会丢失精度，导致数据修改前后没有变化
		if x.i == val {
			continue
		}

		if CompareAndSwapFloat32(&x.i, val, val+1) {
			t.Fatalf("should not have swapped %.32f %.32f", val, val+2)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#v val+1=%#v", x.i, val+1)
		}
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#v _ %#v != %#v _ %#v", x.before, x.after, float32(magic32), float32(magic32))
	}
}

func TestAddFloat32(t *testing.T) {
	var x struct {
		before float32
		i      float32
		after  float32
	}
	x.before = magic32
	x.after = magic32
	var j float32
	for delta := float32(1); delta+delta > delta; delta += delta {
		k := AddFloat32(&x.i, delta)
		j += delta
		if x.i != j || k != j {
			t.Fatalf("delta=%f i=%f j=%f k=%f", delta, x.i, j, k)
		}
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#v _ %#v != %#v _ %#v", x.before, x.after, float32(magic32), float32(magic32))
	}
}

func TestLoadFloat32(t *testing.T) {
	var x struct {
		before float32
		i      float32
		after  float32
	}
	x.before = magic32
	x.after = magic32
	for delta := float32(1); delta+delta > delta; delta += delta {
		k := LoadFloat32(&x.i)
		if k != x.i {
			t.Fatalf("delta=%f i=%f k=%f", delta, x.i, k)
		}
		x.i += delta
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#v _ %#v != %#v _ %#v", x.before, x.after, float32(magic32), float32(magic32))
	}
}

func TestStoreFloat32(t *testing.T) {
	var x struct {
		before float32
		i      float32
		after  float32
	}
	x.before = magic32
	x.after = magic32
	v := float32(0)
	for delta := float32(1); delta+delta > delta; delta += delta {
		StoreFloat32(&x.i, v)
		if x.i != v {
			t.Fatalf("delta=%f i=%f v=%f", delta, x.i, v)
		}
		v += delta
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#v _ %#v != %#v _ %#v", x.before, x.after, float32(magic32), float32(magic32))
	}
}
