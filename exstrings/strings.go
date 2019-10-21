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

import (
	"github.com/thinkeridea/go-extend/exbytes"
)

// Reverse 反转字符串，通过 https://golang.org/doc/code.html#Library 收集
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReverseASCII 反转字符串, 只支持单字节编码，可以提供更快的反转
func ReverseASCII(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return exbytes.ToString(b)
}

// UnsafeReverseASCII 反转字符串, 只支持单字节编码，不支持字面量字符串，
// 原地反转字符串，可以提供更快的性能，但需要注意安全。
func UnsafeReverseASCII(s string) string {
	b := UnsafeToBytes(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return exbytes.ToString(b)
}

// Replace 替换字符串
// 该方法是对标准库 strings.Replace 修改，配合 unsafe 包能有效减少内存分配。
func Replace(s, old, new string, n int) string {
	return exbytes.ToString(UnsafeReplaceToBytes(s, old, new, n))
}

/*
Repeat 返回由字符串s的计数副本组成的新字符串。
该方法是对标准库 strings.Repeat 修改，对于创建大字符串能有效减少内存分配。

如果计数为负或 len(s) * count 溢出将触发panic。
*/
func Repeat(s string, count int) string {
	return exbytes.ToString(RepeatToBytes(s, count))
}

// Join 使用 sep 连接 a 的字符串。
// 该方法是对标准库 strings.Join 修改，配合 unsafe 包能有效减少内存分配。
func Join(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1]
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1] + sep + a[2]
	}

	return exbytes.ToString(JoinToBytes(a, sep))
}

// Copy 拷贝一个字符串，在截取字符串之后，我们得到一个大字符串的引用，这会导致内存泄漏。
// 如果我们引用一个较大字符串的子串，建议进行 copy 以便 GC 可以快速回收大字符串。
// 例如: exstrings.Copy(s[10:50])  这会得到一个子串的拷贝，原字符串不使用可以被 GC 回收。
func Copy(src string) string {
	buf := make([]byte, len(src))
	copy(buf, src)
	return exbytes.ToString(buf)
}
