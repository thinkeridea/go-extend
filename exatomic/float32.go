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

// SwapFloat32 atomically stores new into *addr and returns the previous *addr value.
func SwapFloat32(addr *float32, new float32) (old float32) {
	v := atomic.SwapUint32((*uint32)(unsafe.Pointer(addr)), *(*uint32)(unsafe.Pointer(&new)))
	return *(*float32)((unsafe.Pointer)(&v))
}

// CompareAndSwapFloat32 executes the compare-and-swap operation for an float32 value.
func CompareAndSwapFloat32(addr *float32, old, new float32) (swapped bool) {
	return atomic.CompareAndSwapUint32((*uint32)(unsafe.Pointer(addr)), *(*uint32)(unsafe.Pointer(&old)), *(*uint32)(unsafe.Pointer(&new)))
}

// AddFloat32 atomically adds delta to *addr and returns the new value.
func AddFloat32(addr *float32, delta float32) (new float32) {
	var cur, next uint32
	var curVal, nextVal float32
	for {
		cur = atomic.LoadUint32((*uint32)(unsafe.Pointer(addr)))
		curVal = *(*float32)((unsafe.Pointer)(&cur))
		nextVal = curVal + delta
		next = *(*uint32)(unsafe.Pointer(&nextVal))
		if atomic.CompareAndSwapUint32((*uint32)(unsafe.Pointer(addr)), cur, next) {
			return nextVal
		}
	}
}

// LoadFloat32 atomically loads *addr.
func LoadFloat32(addr *float32) (val float32) {
	v := atomic.LoadUint32((*uint32)(unsafe.Pointer(addr)))
	return *(*float32)((unsafe.Pointer)(&v))
}

// StoreFloat32 atomically stores val into *addr.
func StoreFloat32(addr *float32, val float32) {
	atomic.StoreUint32((*uint32)(unsafe.Pointer(addr)), *(*uint32)(unsafe.Pointer(&val)))
}
