// Copyright (C) 2018  Qi Yin <qiyin@thinkeridea.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package exatomic

import (
	"sync/atomic"
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
	magic64 = 0xdeddeadbeefbeef
)

var test64err = func() (err interface{}) {
	defer func() {
		err = recover()
	}()
	var x int64
	atomic.AddInt64(&x, 1)
	return nil
}()

func TestSwapFloat64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	var x struct {
		before float64
		i      float64
		after  float64
	}
	x.before = magic64
	x.after = magic64
	var j float64
	for delta := float64(0.1); delta+delta > delta; delta += delta {
		k := SwapFloat64(&x.i, delta)
		if x.i != delta || k != j {
			t.Fatalf("delta=%f i=%f j=%f k=%f", delta, x.i, j, k)
		}
		j = delta
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#v _ %#v != %#v _ %#v", x.before, x.after, float64(magic64), float64(magic64))
	}
}

func TestCompareAndSwapFloat64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	var x struct {
		before float64
		i      float64
		after  float64
	}
	x.before = magic64
	x.after = magic64
	for val := float64(0.01); val+val > val; val += val {
		x.i = val
		if !CompareAndSwapFloat64(&x.i, val, val+1) {
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

		if CompareAndSwapFloat64(&x.i, val, val+1) {
			t.Fatalf("should not have swapped %.64f %.64f", val, val+2)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#v val+1=%#v", x.i, val+1)
		}
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#v _ %#v != %#v _ %#v", x.before, x.after, float64(magic64), float64(magic64))
	}
}

func TestAddFloat64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	var x struct {
		before float64
		i      float64
		after  float64
	}
	x.before = magic64
	x.after = magic64
	var j float64
	for delta := float64(1); delta+delta > delta; delta += delta {
		k := AddFloat64(&x.i, delta)
		j += delta
		if x.i != j || k != j {
			t.Fatalf("delta=%f i=%f j=%f k=%f", delta, x.i, j, k)
		}
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#v _ %#v != %#v _ %#v", x.before, x.after, float64(magic64), float64(magic64))
	}
}

func TestLoadFloat64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	var x struct {
		before float64
		i      float64
		after  float64
	}
	x.before = magic64
	x.after = magic64
	for delta := float64(1); delta+delta > delta; delta += delta {
		k := LoadFloat64(&x.i)
		if k != x.i {
			t.Fatalf("delta=%f i=%f k=%f", delta, x.i, k)
		}
		x.i += delta
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#v _ %#v != %#v _ %#v", x.before, x.after, float64(magic64), float64(magic64))
	}
}

func TestStoreFloat64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	var x struct {
		before float64
		i      float64
		after  float64
	}
	x.before = magic64
	x.after = magic64
	v := float64(0)
	for delta := float64(1); delta+delta > delta; delta += delta {
		StoreFloat64(&x.i, v)
		if x.i != v {
			t.Fatalf("delta=%f i=%f v=%f", delta, x.i, v)
		}
		v += delta
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#v _ %#v != %#v _ %#v", x.before, x.after, float64(magic64), float64(magic64))
	}
}
