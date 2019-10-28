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

package exbytes

import (
	"bytes"

	"github.com/thinkeridea/go-extend/exunicode/exutf8"
)

/*
Replace 思路来源于 bytes.Replace，bytes.Replace 总是返回 s 的副本，
有些场景源数据生命周期非常短，且可以原地替换，如果这么实现可以减少极大的内存分配。

len(old) >= len(new) 会执行原地替换，这回浪费一部分空间，但是会减少内存分配，
建议输入生命周期较短的数据， len(old) < len(new) 会调用 bytes.Replace 并返回一个替换后的副本。

最佳实践是使用 Replace 的结果覆盖源变量，避免再次对源数据引用，导致访问过时的数据，并且数据内容错乱，如下：

	var s []byte
	s = exbytes.Replace(s, []byte(" "), []byte(""), -1)

关于字符串可以结合 exstrings.UnsafeToBytes 来实现，要避免常量字符串和字面量字符串，否者会产生运行时错误。
*/
func Replace(s, old, new []byte, n int) []byte {
	if n == 0 {
		return s
	}

	if len(old) < len(new) {
		return bytes.Replace(s, old, new, n)
	}

	if n < 0 {
		n = len(s)
	}

	var wid, i, j, w int
	for i, j = 0, 0; i < len(s) && j < n; j++ {
		wid = bytes.Index(s[i:], old)
		if wid < 0 {
			break
		}

		w += copy(s[w:], s[i:i+wid])
		w += copy(s[w:], new)
		i += wid + len(old)
	}

	w += copy(s[w:], s[i:])
	return s[0:w]
}

// Reverse 原地反转 []byte
func Reverse(s []byte) {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Sub 是 exutf8.RuneSub 的别名，提供字符数量截取字节数组的方法，针对多字节字符安全高效的截取
// 如果 start 是非负数，返回的字符串将从 string 的 start 位置开始，从 0 开始计算。例如，在字符串 “abcdef” 中，在位置 0 的字符是 “a”，位置 2 的字符串是 “c” 等等。
// 如果 start 是负数，返回的字符串将从 string 结尾处向前数第 start 个字符开始。
// 如果 string 的长度小于 start，将返回空字符串。
//
// 如果提供了正数的 length，返回的字符串将从 start 处开始最多包括 length 个字符（取决于 string 的长度）。
// 如果提供了负数的 length，那么 string 末尾处的 length 个字符将会被省略（若 start 是负数则从字符串尾部算起）。如果 start 不在这段文本中，那么将返回空字符串。
// 如果提供了值为 0 的 length，返回的子字符串将从 start 位置开始直到字符串结尾。
func Sub(p []byte, start, length int) []byte {
	return exutf8.RuneSub(p, start, length)
}
