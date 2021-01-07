// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package helper

import (
	"testing"
)

func TestBool(t *testing.T) {
	var v bool = true

	x := Bool(v)
	y := Bool(true)

	if x == y {
		t.Errorf("Bool(%v) = %v equal Bool(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Bool(%v) = %v, want %v", v, *x, v)
	}
}

func TestUint8(t *testing.T) {
	var v uint8 = 123

	x := Uint8(v)
	y := Uint8(123)

	if x == y {
		t.Errorf("Uint8(%v) = %v equal Uint8(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Uint8(%v) = %v, want %v", v, *x, v)
	}
}

func TestUint16(t *testing.T) {
	var v uint16 = 123

	x := Uint16(v)
	y := Uint16(123)

	if x == y {
		t.Errorf("Uint16(%v) = %v equal Uint16(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Uint16(%v) = %v, want %v", v, *x, v)
	}
}

func TestUint32(t *testing.T) {
	var v uint32 = 123

	x := Uint32(v)
	y := Uint32(123)

	if x == y {
		t.Errorf("Uint32(%v) = %v equal Uint32(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Uint32(%v) = %v, want %v", v, *x, v)
	}
}

func TestUint64(t *testing.T) {
	var v uint64 = 123

	x := Uint64(v)
	y := Uint64(123)

	if x == y {
		t.Errorf("Uint64(%v) = %v equal Uint64(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Uint64(%v) = %v, want %v", v, *x, v)
	}
}

func TestInt8(t *testing.T) {
	var v int8 = 123

	x := Int8(v)
	y := Int8(123)

	if x == y {
		t.Errorf("Int8(%v) = %v equal Int8(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Int8(%v) = %v, want %v", v, *x, v)
	}
}

func TestInt16(t *testing.T) {
	var v int16 = 123

	x := Int16(v)
	y := Int16(123)

	if x == y {
		t.Errorf("Int16(%v) = %v equal Int16(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Int16(%v) = %v, want %v", v, *x, v)
	}
}

func TestInt32(t *testing.T) {
	var v int32 = 123

	x := Int32(v)
	y := Int32(123)

	if x == y {
		t.Errorf("Int32(%v) = %v equal Int32(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Int32(%v) = %v, want %v", v, *x, v)
	}
}

func TestInt64(t *testing.T) {
	var v int64 = 123

	x := Int64(v)
	y := Int64(123)

	if x == y {
		t.Errorf("Int64(%v) = %v equal Int64(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Int64(%v) = %v, want %v", v, *x, v)
	}
}

func TestFloat32(t *testing.T) {
	var v float32 = 123

	x := Float32(v)
	y := Float32(123)

	if x == y {
		t.Errorf("Float32(%v) = %v equal Float32(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Float32(%v) = %v, want %v", v, *x, v)
	}
}

func TestFloat64(t *testing.T) {
	var v float64 = 123

	x := Float64(v)
	y := Float64(123)

	if x == y {
		t.Errorf("Float64(%v) = %v equal Float64(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Float64(%v) = %v, want %v", v, *x, v)
	}
}

func TestComplex64(t *testing.T) {
	var v complex64 = 1.23 + 4.567i

	x := Complex64(v)
	y := Complex64(1.23 + 4.567i)

	if x == y {
		t.Errorf("Complex64(%v) = %v equal Complex64(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Complex64(%v) = %v, want %v", v, *x, v)
	}
}

func TestComplex128(t *testing.T) {
	var v complex128 = 1.23 + 4.567i

	x := Complex128(v)
	y := Complex128(1.23 + 4.567i)

	if x == y {
		t.Errorf("Complex128(%v) = %v equal Complex128(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Complex128(%v) = %v, want %v", v, *x, v)
	}
}

func TestString(t *testing.T) {
	var v string = "string"

	x := String(v)
	y := String("string")

	if x == y {
		t.Errorf("String(%v) = %v equal String(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("String(%v) = %v, want %v", v, *x, v)
	}
}

func TestInt(t *testing.T) {
	var v int = 123

	x := Int(v)
	y := Int(123)

	if x == y {
		t.Errorf("Int(%v) = %v equal Int(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Int(%v) = %v, want %v", v, *x, v)
	}
}

func TestUint(t *testing.T) {
	var v uint = 123

	x := Uint(v)
	y := Uint(123)

	if x == y {
		t.Errorf("Uint(%v) = %v equal Uint(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Uint(%v) = %v, want %v", v, *x, v)
	}
}

func TestByte(t *testing.T) {
	var v byte = 'a'

	x := Byte(v)
	y := Byte('a')

	if x == y {
		t.Errorf("Byte(%v) = %v equal Byte(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Byte(%v) = %v, want %v", v, *x, v)
	}
}

func TestRune(t *testing.T) {
	var v rune = 'a'

	x := Rune(v)
	y := Rune('a')

	if x == y {
		t.Errorf("Rune(%v) = %v equal Rune(%v) = %v", v, x, v, y)
	}

	if *x != v {
		t.Errorf("Rune(%v) = %v, want %v", v, *x, v)
	}
}
