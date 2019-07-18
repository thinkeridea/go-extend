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

package datalog

import (
	"strings"
	"sync"

	"github.com/thinkeridea/go-extend/exbytes"
	"github.com/thinkeridea/go-extend/exstrings"
)

const (
	// FieldSep 字段分隔符
	FieldSep = "\x01"
	// NewLine 换行符
	NewLine = "\x03\n"

	// ArraySep 数组字段分隔符
	ArraySep = "\x02"
	// ArrayFieldSep 数组分隔符
	ArrayFieldSep = "\x04"
)

// Record 一行日志记录
type Record []string

// NewRecord 创建长度固定的日志记录
func NewRecord(len int) Record {
	return make(Record, len)
}

// NewRecordPool 创建长度固定的日志记录缓存池
func NewRecordPool(len int) *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			return make(Record, len)
		},
	}
}

// ToBytes 使用 sep 连接 Record，并在末尾添加 newline 换行符
// 注意：这个方法会替换 sep 与 newline 为空字符串
func (l Record) ToBytes(sep, newline string) []byte {
	for i := len(l) - 1; i >= 0; i-- {
		// 提前检查是否包含特殊字符，以便跳过字符串替换
		if strings.Index(l[i], sep) < 0 && strings.Index(l[i], newline) < 0 {
			continue
		}

		b := []byte(l[i]) // 这会重新分配内存，避免原地替换导致引用字符串被修改
		b = exbytes.Replace(b, exstrings.UnsafeToBytes(sep), []byte{' '}, -1)
		b = exbytes.Replace(b, exstrings.UnsafeToBytes(newline), []byte{' '}, -1)
		l[i] = exbytes.ToString(b)
	}

	return l.Join(sep, newline)
}

// UnsafeToBytes 使用 sep 连接 Record，并在末尾添加 newline 换行符
// 注意：这个方法会替换 sep 与 newline 为空字符串，替换采用原地替换，这会导致所有引用字符串被修改
// 必须明白其作用，否者将会导致意想不到的结果。但是这会大幅度减少内存分配，提升程序性能
// 我在项目中大量使用，我总是在请求最后记录日志，这样我不会再访问引用的字符串
func (l Record) UnsafeToBytes(sep, newline string) []byte {
	for i := len(l) - 1; i >= 0; i-- {
		b := exstrings.UnsafeToBytes(l[i])
		b = exbytes.Replace(b, exstrings.UnsafeToBytes(sep), []byte{' '}, -1)
		b = exbytes.Replace(b, exstrings.UnsafeToBytes(newline), []byte{' '}, -1)
		l[i] = exbytes.ToString(b)
	}

	return l.Join(sep, newline)
}

// Join 使用 sep 连接 Record， 并在末尾追加 suffix
// 这个类似 strings.Join 方法，但是避免了连接后追加后缀（往往是换行符）导致的内存分配
// 这个方法直接返回需要的 []byte 类型， 可以减少类型转换，降低内存分配导致的性能问题
func (l Record) Join(sep, suffix string) []byte {
	if len(l) == 0 {
		return []byte(suffix)
	}

	n := len(sep) * (len(l) - 1)
	for i := 0; i < len(l); i++ {
		n += len(l[i])
	}

	n += len(suffix)
	b := make([]byte, n)
	bp := copy(b, l[0])
	for i := 1; i < len(l); i++ {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], l[i])
	}
	copy(b[bp:], suffix)
	return b
}

// Clean 清空 Record 中的所有元素，如果使用 sync.Pool 在放回 Pool 之前应该清空 Record，避免内存泄漏
// 该方法没有太多的开销，可以放心的使用，只是为 Record 中的字段赋值为空字符串，空字符串会在编译时处理，没有额外的内存分配
func (l Record) Clean() {
	for i := len(l) - 1; i >= 0; i-- {
		l[i] = ""
	}
}

// ArrayJoin 使用 sep 连接 Record，其结果作为数组字段的值
func (l Record) ArrayJoin(sep string) string {
	return exstrings.Join(l, sep)
}

// ArrayFieldJoin 使用 fieldSep 连接 Record，其结果作为一个数组的单元
// 注意：这个方法会替换 fieldSep 与 arraySep 为空字符串，替换采用原地替换
func (l Record) ArrayFieldJoin(fieldSep, arraySep string) string {
	for i := len(l) - 1; i >= 0; i-- {
		// 提前检查是否包含特殊字符，以便跳过字符串替换
		if strings.Index(l[i], fieldSep) < 0 && strings.Index(l[i], arraySep) < 0 {
			continue
		}

		b := []byte(l[i]) // 这会重新分配内存，避免原地替换导致引用字符串被修改
		b = exbytes.Replace(b, exstrings.UnsafeToBytes(fieldSep), []byte{' '}, -1)
		b = exbytes.Replace(b, exstrings.UnsafeToBytes(arraySep), []byte{' '}, -1)
		l[i] = exbytes.ToString(b)
	}

	return exstrings.Join(l, fieldSep)
}

// UnsafeArrayFieldJoin 使用 fieldSep 连接 Record，其结果作为一个数组的单元
// 注意：这个方法会替换 fieldSep 与 arraySep 为空字符串 ，这会导致所有引用字符串被修改
// 必须明白其作用，否者将会导致意想不到的结果。但是这会大幅度减少内存分配，提升程序性能
// 我在项目中大量使用，我总是在请求最后记录日志，这样我不会再访问引用的字符串
func (l Record) UnsafeArrayFieldJoin(fieldSep, arraySep string) string {
	for i := len(l) - 1; i >= 0; i-- {
		b := exstrings.UnsafeToBytes(l[i])
		b = exbytes.Replace(b, exstrings.UnsafeToBytes(fieldSep), []byte{' '}, -1)
		b = exbytes.Replace(b, exstrings.UnsafeToBytes(arraySep), []byte{' '}, -1)
		l[i] = exbytes.ToString(b)
	}

	return exstrings.Join(l, fieldSep)
}
