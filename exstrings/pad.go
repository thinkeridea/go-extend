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

package exstrings

import "github.com/thinkeridea/go-extend/exbytes"

const (
	// PadLeft 在左侧填充字符串为指定长度。
	PadLeft = iota

	// PadBoth 在两边填充字符串为指定长度，如果补充长度是奇数，右边的字符会更多一些。
	PadBoth

	// PadRight 在右侧填充字符串为指定长度。
	PadRight
)

// repeat 重复字符串到指定长度, []byte必须有充足的容量。
func repeat(b []byte, pad string, padLen int) {
	bp := copy(b[:padLen], pad)
	for bp < padLen {
		copy(b[bp:padLen], b[:bp])
		bp *= 2
	}
}

/*
Pad 使用另一个字符串填充字符串为指定长度。

该函数返回 s 被从左端、右端或者同时两端被填充到制定长度后的结果。
填充方向由 falg 控制，可选值：PadLeft、PadBoth、PadRight。

在两边填充字符串为指定长度，如果补充长度是奇数，右边的字符会更多一些。
*/
func Pad(s, pad string, c, falg int) string {
	padLen := c - len(s)
	if padLen <= 0 {
		return s
	}

	b := make([]byte, c)
	switch falg {
	case PadLeft:
		repeat(b, pad, padLen)
		copy(b[padLen:], s)
	case PadRight:
		l := copy(b, s)
		repeat(b[l:], pad, padLen)
	case PadBoth:
		repeat(b, pad, padLen/2)
		l := copy(b[padLen/2:], s)
		l += padLen / 2
		repeat(b[l:], pad, c-l)
	}

	return exbytes.ToString(b)
}

// LeftPad 使用另一个字符串从左端填充字符串为指定长度。
func LeftPad(s, pad string, c int) string {
	padLen := c - len(s)
	if padLen <= 0 {
		return s
	}

	b := make([]byte, c)
	repeat(b, pad, padLen)
	copy(b[padLen:], s)

	return exbytes.ToString(b)
}

// RightPad 使用另一个字符串从右端填充字符串为指定长度。
func RightPad(s, pad string, c int) string {
	padLen := c - len(s)
	if padLen <= 0 {
		return s
	}

	b := make([]byte, c)
	l := copy(b, s)
	repeat(b[l:], pad, padLen)

	return exbytes.ToString(b)
}

// BothPad 使用另一个字符串从两端填充字符串为指定长度，
// 如果补充长度是奇数，右边的字符会更多一些。
func BothPad(s, pad string, c int) string {
	padLen := c - len(s)
	if padLen <= 0 {
		return s
	}

	b := make([]byte, c)
	repeat(b, pad, padLen/2)
	l := copy(b[padLen/2:], s)
	l += padLen / 2
	repeat(b[l:], pad, c-l)

	return exbytes.ToString(b)
}

/*
UnsafePad 使用另一个字符串填充字符串为指定长度。

该函数使用 unsafe 包转换数据类型，降低内存分配。

该函数返回 s 被从左端、右端或者同时两端被填充到制定长度后的结果。
填充方向由 falg 控制，可选值：PadLeft、PadBoth、PadRight。

在两边填充字符串为指定长度，如果补充长度是奇数，右边的字符会更多一些。

Deprecated: 不在使用 Unsafe 前缀
*/
func UnsafePad(s, pad string, c, falg int) string {
	return Pad(s, pad, c, falg)
}

// UnsafeLeftPad 使用另一个字符串从左端填充字符串为指定长度。
//
// 该函数使用 unsafe 包转换数据类型，降低内存分配。
// Deprecated: 不在使用 Unsafe 前缀
func UnsafeLeftPad(s, pad string, c int) string {
	return LeftPad(s, pad, c)
}

// UnsafeRightPad 使用另一个字符串从右端填充字符串为指定长度。
//
// 该函数使用 unsafe 包转换数据类型，降低内存分配。
// Deprecated: 不在使用 Unsafe 前缀
func UnsafeRightPad(s, pad string, c int) string {
	return RightPad(s, pad, c)
}

// UnsafeBothPad 使用另一个字符串从两端填充字符串为指定长度，
// 如果补充长度是奇数，右边的字符会更多一些。
//
// 该函数使用 unsafe 包转换数据类型，降低内存分配。
// Deprecated: 不在使用 Unsafe 前缀
func UnsafeBothPad(s, pad string, c int) string {
	return BothPad(s, pad, c)
}
