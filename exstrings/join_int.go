// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package exstrings

import (
	"strconv"

	"github.com/thinkeridea/go-extend/pool"
)

var buffPool = pool.GetBuff64()

// JoinInts 使用 sep 连接 []int 并返回连接的字符串
func JoinInts(v []int, sep string) string {
	return joinInts(len(v), sep, func(i int) string {
		return strconv.FormatInt(int64(v[i]), 10)
	})
}

// JoinInt8s 使用 sep 连接 []int8 并返回连接的字符串
func JoinInt8s(v []int8, sep string) string {
	return joinInts(len(v), sep, func(i int) string {
		return strconv.FormatInt(int64(v[i]), 10)
	})
}

// JoinInt16s 使用 sep 连接 []int16 并返回连接的字符串
func JoinInt16s(v []int16, sep string) string {
	return joinInts(len(v), sep, func(i int) string {
		return strconv.FormatInt(int64(v[i]), 10)
	})
}

// JoinInt32s 使用 sep 连接 []int32 并返回连接的字符串
func JoinInt32s(v []int32, sep string) string {
	return joinInts(len(v), sep, func(i int) string {
		return strconv.FormatInt(int64(v[i]), 10)
	})
}

// JoinInt64s 使用 sep 连接 []int64 并返回连接的字符串
func JoinInt64s(v []int64, sep string) string {
	return joinInts(len(v), sep, func(i int) string {
		return strconv.FormatInt(v[i], 10)
	})
}

// JoinUints 使用 sep 连接 []uint 并返回连接的字符串
func JoinUints(v []uint, sep string) string {
	return joinInts(len(v), sep, func(i int) string {
		return strconv.FormatUint(uint64(v[i]), 10)
	})
}

// JoinUint8s 使用 sep 连接 []uint8 并返回连接的字符串
func JoinUint8s(v []uint8, sep string) string {
	return joinInts(len(v), sep, func(i int) string {
		return strconv.FormatUint(uint64(v[i]), 10)
	})
}

// JoinUint16s 使用 sep 连接 []uint16 并返回连接的字符串
func JoinUint16s(v []uint16, sep string) string {
	return joinInts(len(v), sep, func(i int) string {
		return strconv.FormatUint(uint64(v[i]), 10)
	})
}

// JoinUint32s 使用 sep 连接 []uint32 并返回连接的字符串
func JoinUint32s(v []uint32, sep string) string {
	return joinInts(len(v), sep, func(i int) string {
		return strconv.FormatUint(uint64(v[i]), 10)
	})
}

// JoinUint64s 使用 sep 连接 []uint64 并返回连接的字符串
func JoinUint64s(v []uint64, sep string) string {
	return joinInts(len(v), sep, func(i int) string {
		return strconv.FormatUint(v[i], 10)
	})
}

func joinInts(n int, sep string, f func(i int) string) string {
	buf := buffPool.Get()
	defer buffPool.Put(buf)
	for i := 0; i < n; i++ {
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(f(i))
	}

	return buf.String()
}
