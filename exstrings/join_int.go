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
	"strconv"

	"github.com/thinkeridea/go-extend/pool"
)

var buffPool = pool.GetBuff64()

// JoinInts 使用 sep 连接 []int 并返回连接的字符串
func JoinInts(i []int, sep string) string {
	buf := buffPool.Get()
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	}

	buffPool.Put(buf)
	return buf.String()
}

// JoinInt8s 使用 sep 连接 []int8 并返回连接的字符串
func JoinInt8s(i []int8, sep string) string {
	buf := buffPool.Get()
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	}

	buffPool.Put(buf)
	return buf.String()
}

// JoinInt16s 使用 sep 连接 []int16 并返回连接的字符串
func JoinInt16s(i []int16, sep string) string {
	buf := buffPool.Get()
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	}

	buffPool.Put(buf)
	return buf.String()
}

// JoinInt32s 使用 sep 连接 []int32 并返回连接的字符串
func JoinInt32s(i []int32, sep string) string {
	buf := buffPool.Get()
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	}

	buffPool.Put(buf)
	return buf.String()
}

// JoinInt64s 使用 sep 连接 []int64 并返回连接的字符串
func JoinInt64s(i []int64, sep string) string {
	buf := buffPool.Get()
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatInt(v, 10))
	}

	buffPool.Put(buf)
	return buf.String()
}

// JoinUints 使用 sep 连接 []uint 并返回连接的字符串
func JoinUints(i []uint, sep string) string {
	buf := buffPool.Get()
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatUint(uint64(v), 10))
	}

	buffPool.Put(buf)
	return buf.String()
}

// JoinUint8s 使用 sep 连接 []uint8 并返回连接的字符串
func JoinUint8s(i []uint8, sep string) string {
	buf := buffPool.Get()
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatUint(uint64(v), 10))
	}

	buffPool.Put(buf)
	return buf.String()
}

// JoinUint16s 使用 sep 连接 []uint16 并返回连接的字符串
func JoinUint16s(i []uint16, sep string) string {
	buf := buffPool.Get()
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatUint(uint64(v), 10))
	}

	buffPool.Put(buf)
	return buf.String()
}

// JoinUint32s 使用 sep 连接 []uint32 并返回连接的字符串
func JoinUint32s(i []uint32, sep string) string {
	buf := buffPool.Get()
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatUint(uint64(v), 10))
	}

	buffPool.Put(buf)
	return buf.String()
}

// JoinUint64s 使用 sep 连接 []uint64 并返回连接的字符串
func JoinUint64s(i []uint64, sep string) string {
	buf := buffPool.Get()
	for _, v := range i {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(strconv.FormatUint(v, 10))
	}

	buffPool.Put(buf)
	return buf.String()
}
