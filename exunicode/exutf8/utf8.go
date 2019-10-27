package exutf8

import (
	"unicode/utf8"
)

const (
	// t1 = 0x00 // 0000 0000
	// tx = 0x80 // 1000 0000
	// t2 = 0xC0 // 1100 0000
	// t3 = 0xE0 // 1110 0000
	// t4 = 0xF0 // 1111 0000
	// t5 = 0xF8 // 1111 1000
	//
	// maskx = 0x3F // 0011 1111
	// mask2 = 0x1F // 0001 1111
	// mask3 = 0x0F // 0000 1111
	// mask4 = 0x07 // 0000 0111
	//
	// rune1Max = 1<<7 - 1
	// rune2Max = 1<<11 - 1
	// rune3Max = 1<<16 - 1

	// The default lowest and highest continuation byte.
	locb = 0x80 // 1000 0000
	hicb = 0xBF // 1011 1111

	// These names of these constants are chosen to give nice alignment in the
	// table below. The first nibble is an index into acceptRanges or F for
	// special one-byte cases. The second nibble is the Rune length or the
	// Status for the special one-byte case.
	xx = 0xF1 // invalid: size 1
	as = 0xF0 // ASCII: size 1
	s1 = 0x02 // accept 0, size 2
	s2 = 0x13 // accept 1, size 3
	s3 = 0x03 // accept 0, size 3
	s4 = 0x23 // accept 2, size 3
	s5 = 0x34 // accept 3, size 4
	s6 = 0x04 // accept 0, size 4
	s7 = 0x44 // accept 4, size 4
)

// first is information about the first byte in a UTF-8 sequence.
var first = [256]uint8{
	//   1   2   3   4   5   6   7   8   9   A   B   C   D   E   F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x00-0x0F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x10-0x1F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x20-0x2F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x30-0x3F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x40-0x4F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x50-0x5F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x60-0x6F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x70-0x7F
	//   1   2   3   4   5   6   7   8   9   A   B   C   D   E   F
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0x80-0x8F
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0x90-0x9F
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0xA0-0xAF
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0xB0-0xBF
	xx, xx, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, // 0xC0-0xCF
	s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, // 0xD0-0xDF
	s2, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s4, s3, s3, // 0xE0-0xEF
	s5, s6, s6, s6, s7, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0xF0-0xFF
}

// acceptRange gives the range of valid values for the second byte in a UTF-8
// sequence.
type acceptRange struct {
	lo uint8 // lowest value for second byte.
	hi uint8 // highest value for second byte.
}

// acceptRanges has size 16 to avoid bounds checks in the code that uses it.
var acceptRanges = [16]acceptRange{
	0: {locb, hicb},
	1: {0xA0, hicb},
	2: {locb, 0x9F},
	3: {0x90, hicb},
	4: {locb, 0x8F},
}

// RuneIndex 返回 p 中第 n 个字符的位置索引，可以通过索引位置截取 []byte
// 如果 n 超过 p 中字符的数量，则第二个参数返回 false
// 错误的或短的编码被当做宽度为一个字节的单一字符。
func RuneIndex(p []byte, n int) (int, bool) {
	if n <= 0 {
		return 0, true
	}

	var i int
	for i < len(p) {
		if n <= 0 {
			break
		}

		n--
		if p[i] < utf8.RuneSelf {
			// ASCII fast path
			i++
			continue
		}

		x := first[p[i]]
		if x == xx {
			i++ // invalid.
			continue
		}

		size := int(x & 7)
		if i+size > len(p) {
			i++ // Short or invalid.
			continue
		}
		accept := acceptRanges[x>>4]
		if c := p[i+1]; c < accept.lo || accept.hi < c {
			size = 1
		} else if size == 2 {
		} else if c := p[i+2]; c < locb || hicb < c {
			size = 1
		} else if size == 3 {
		} else if c := p[i+3]; c < locb || hicb < c {
			size = 1
		}
		i += size

		if n <= 0 {
			break
		}
	}

	return i, n <= 0
}

// RuneIndexInString 返回 s 中第 n 个字符的位置索引，可以通过索引位置截取字符串
// 如果 n 超过 s 中字符的数量，则第二个参数返回 false
// 错误的或短的编码被当做宽度为一个字节的单一字符。
func RuneIndexInString(s string, n int) (int, bool) {
	if n <= 0 {
		return 0, true
	}

	var i int
	for i < len(s) {
		if n <= 0 {
			break
		}

		n--
		if s[i] < utf8.RuneSelf {
			// ASCII fast path
			i++
			continue
		}

		x := first[s[i]]
		if x == xx {
			i++ // invalid.
			continue
		}

		size := int(x & 7)
		if i+size > len(s) {
			i++ // Short or invalid.
			continue
		}
		accept := acceptRanges[x>>4]
		if c := s[i+1]; c < accept.lo || accept.hi < c {
			size = 1
		} else if size == 2 {
		} else if c := s[i+2]; c < locb || hicb < c {
			size = 1
		} else if size == 3 {
		} else if c := s[i+3]; c < locb || hicb < c {
			size = 1
		}
		i += size

		if n <= 0 {
			break
		}
	}

	return i, n <= 0
}
