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
