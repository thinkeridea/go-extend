// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package exatomic

import (
	"sync/atomic"
	"unsafe"
)

// SwapFloat64 atomically stores new into *addr and returns the previous *addr value.
func SwapFloat64(addr *float64, new float64) (old float64) {
	v := atomic.SwapUint64((*uint64)(unsafe.Pointer(addr)), *(*uint64)(unsafe.Pointer(&new)))
	return *(*float64)((unsafe.Pointer)(&v))
}

// CompareAndSwapFloat64 executes the compare-and-swap operation for an float64 value.
func CompareAndSwapFloat64(addr *float64, old, new float64) (swapped bool) {
	return atomic.CompareAndSwapUint64((*uint64)(unsafe.Pointer(addr)), *(*uint64)(unsafe.Pointer(&old)), *(*uint64)(unsafe.Pointer(&new)))
}

// AddFloat64 atomically adds delta to *addr and returns the new value.
func AddFloat64(addr *float64, delta float64) (new float64) {
	var cur, next uint64
	var curVal, nextVal float64
	for {
		cur = atomic.LoadUint64((*uint64)(unsafe.Pointer(addr)))
		curVal = *(*float64)((unsafe.Pointer)(&cur))
		nextVal = curVal + delta
		next = *(*uint64)(unsafe.Pointer(&nextVal))
		if atomic.CompareAndSwapUint64((*uint64)(unsafe.Pointer(addr)), cur, next) {
			return nextVal
		}
	}
}

// LoadFloat64 atomically loads *addr.
func LoadFloat64(addr *float64) (val float64) {
	v := atomic.LoadUint64((*uint64)(unsafe.Pointer(addr)))
	return *(*float64)((unsafe.Pointer)(&v))
}

// StoreFloat64 atomically stores val into *addr.
func StoreFloat64(addr *float64, val float64) {
	atomic.StoreUint64((*uint64)(unsafe.Pointer(addr)), *(*uint64)(unsafe.Pointer(&val)))
}
