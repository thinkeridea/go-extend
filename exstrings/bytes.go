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
	"strings"
	"unicode/utf8"
)

// UnsafeReplaceToBytes 替换字符串，并返回 []byte，减少类型转换
// 如果输入是字面量字符串，返回结果不可以修改
func UnsafeReplaceToBytes(s, old, new string, n int) []byte {
	if old == new || n == 0 {
		return UnsafeToBytes(s) // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return UnsafeToBytes(s) // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		w += copy(t[w:], s[start:j])
		w += copy(t[w:], new)
		start = j + len(old)
	}
	w += copy(t[w:], s[start:])
	return t[0:w]
}

// ReplaceToBytes 替换字符串，并返回 []byte，减少类型转换
// 如果在字符串串中找不到子串，会发生内存拷贝
// 这个方法返回的 []byte 是安全，可以进行修改
func ReplaceToBytes(s, old, new string, n int) []byte {
	if old == new || n == 0 {
		return []byte(s) // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return []byte(s) // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		w += copy(t[w:], s[start:j])
		w += copy(t[w:], new)
		start = j + len(old)
	}
	w += copy(t[w:], s[start:])
	return t[0:w]
}

/*
RepeatToBytes 返回由字符串s的计数副本组成的 []byte。
该方法是对标准库 strings.Repeat 修改，对于创建大字符串能有效减少内存分配。

如果计数为负或 len(s) * count 溢出将触发panic。
*/
func RepeatToBytes(s string, count int) []byte {
	// Since we cannot return an error on overflow,
	// we should panic if the repeat will generate
	// an overflow.
	// See Issue golang.org/issue/16237
	if count < 0 {
		panic("strings: negative Repeat count")
	} else if count > 0 && len(s)*count/count != len(s) {
		panic("strings: Repeat count causes overflow")
	}

	b := make([]byte, len(s)*count)
	bp := copy(b, s)
	for bp < len(b) {
		copy(b[bp:], b[:bp])
		bp *= 2
	}
	return b
}

// JoinToBytes 使用 sep 连接 a 的字符串并返回 []byte
// 该方法是对标准库 strings.Join 修改，配合 unsafe 包能有效减少内存分配。
func JoinToBytes(a []string, sep string) []byte {
	switch len(a) {
	case 0:
		return []byte{}
	case 1:
		return []byte(a[0])
	}

	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return b
}
