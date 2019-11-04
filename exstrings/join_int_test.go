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

import "testing"

func TestJoinInts(t *testing.T) {
	cases := []struct {
		in       []int
		sep, out string
	}{
		{[]int{1, 2, 3, 4, 5}, "", "12345"},
		{[]int{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
		{[]int{1, 2, 3, 4, 5}, "}{", "1}{2}{3}{4}{5"},
		{[]int{1, 2}, ",", "1,2"},
		{[]int{1}, ",", "1"},
		{[]int{}, ",", ""},
	}
	for _, c := range cases {
		got := JoinInts(c.in, c.sep)
		if got != c.out {
			t.Errorf("JoinInts(%v, %q) == %q, out %q", c.in, c.sep, got, c.out)
		}
	}
}

func TestJoinInt8s(t *testing.T) {
	cases := []struct {
		in       []int8
		sep, out string
	}{
		{[]int8{1, 2, 3, 4, 5}, "", "12345"},
		{[]int8{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
		{[]int8{1, 2, 3, 4, 5}, "}{", "1}{2}{3}{4}{5"},
		{[]int8{1, 2}, ",", "1,2"},
		{[]int8{1}, ",", "1"},
		{[]int8{}, ",", ""},
	}
	for _, c := range cases {
		got := JoinInt8s(c.in, c.sep)
		if got != c.out {
			t.Errorf("JoinInt8s(%v, %q) == %q, out %q", c.in, c.sep, got, c.out)
		}
	}
}

func TestJoinInt16s(t *testing.T) {
	cases := []struct {
		in       []int16
		sep, out string
	}{
		{[]int16{1, 2, 3, 4, 5}, "", "12345"},
		{[]int16{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
		{[]int16{1, 2, 3, 4, 5}, "}{", "1}{2}{3}{4}{5"},
		{[]int16{1, 2}, ",", "1,2"},
		{[]int16{1}, ",", "1"},
		{[]int16{}, ",", ""},
	}
	for _, c := range cases {
		got := JoinInt16s(c.in, c.sep)
		if got != c.out {
			t.Errorf("JoinInt16s(%v, %q) == %q, out %q", c.in, c.sep, got, c.out)
		}
	}
}

func TestJoinInt32s(t *testing.T) {
	cases := []struct {
		in       []int32
		sep, out string
	}{
		{[]int32{1, 2, 3, 4, 5}, "", "12345"},
		{[]int32{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
		{[]int32{1, 2, 3, 4, 5}, "}{", "1}{2}{3}{4}{5"},
		{[]int32{1, 2}, ",", "1,2"},
		{[]int32{1}, ",", "1"},
		{[]int32{}, ",", ""},
	}
	for _, c := range cases {
		got := JoinInt32s(c.in, c.sep)
		if got != c.out {
			t.Errorf("JoinInt32s(%v, %q) == %q, out %q", c.in, c.sep, got, c.out)
		}
	}
}

func TestJoinInt64s(t *testing.T) {
	cases := []struct {
		in       []int64
		sep, out string
	}{
		{[]int64{1, 2, 3, 4, 5}, "", "12345"},
		{[]int64{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
		{[]int64{1, 2, 3, 4, 5}, "}{", "1}{2}{3}{4}{5"},
		{[]int64{1, 2}, ",", "1,2"},
		{[]int64{1}, ",", "1"},
		{[]int64{}, ",", ""},
	}
	for _, c := range cases {
		got := JoinInt64s(c.in, c.sep)
		if got != c.out {
			t.Errorf("JoinInt64s(%v, %q) == %q, out %q", c.in, c.sep, got, c.out)
		}
	}
}

func TestJoinUints(t *testing.T) {
	cases := []struct {
		in       []uint
		sep, out string
	}{
		{[]uint{1, 2, 3, 4, 5}, "", "12345"},
		{[]uint{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
		{[]uint{1, 2, 3, 4, 5}, "}{", "1}{2}{3}{4}{5"},
		{[]uint{1, 2}, ",", "1,2"},
		{[]uint{1}, ",", "1"},
		{[]uint{}, ",", ""},
	}
	for _, c := range cases {
		got := JoinUints(c.in, c.sep)
		if got != c.out {
			t.Errorf("JoinUints(%v, %q) == %q, out %q", c.in, c.sep, got, c.out)
		}
	}
}

func TestJoinUint8s(t *testing.T) {
	cases := []struct {
		in       []uint8
		sep, out string
	}{
		{[]uint8{1, 2, 3, 4, 5}, "", "12345"},
		{[]uint8{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
		{[]uint8{1, 2, 3, 4, 5}, "}{", "1}{2}{3}{4}{5"},
		{[]uint8{1, 2}, ",", "1,2"},
		{[]uint8{1}, ",", "1"},
		{[]uint8{}, ",", ""},
	}
	for _, c := range cases {
		got := JoinUint8s(c.in, c.sep)
		if got != c.out {
			t.Errorf("JoinUint8s(%v, %q) == %q, out %q", c.in, c.sep, got, c.out)
		}
	}
}

func TestJoinUint16s(t *testing.T) {
	cases := []struct {
		in       []uint16
		sep, out string
	}{
		{[]uint16{1, 2, 3, 4, 5}, "", "12345"},
		{[]uint16{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
		{[]uint16{1, 2, 3, 4, 5}, "}{", "1}{2}{3}{4}{5"},
		{[]uint16{1, 2}, ",", "1,2"},
		{[]uint16{1}, ",", "1"},
		{[]uint16{}, ",", ""},
	}
	for _, c := range cases {
		got := JoinUint16s(c.in, c.sep)
		if got != c.out {
			t.Errorf("JoinUint16s(%v, %q) == %q, out %q", c.in, c.sep, got, c.out)
		}
	}
}

func TestJoinUint32s(t *testing.T) {
	cases := []struct {
		in       []uint32
		sep, out string
	}{
		{[]uint32{1, 2, 3, 4, 5}, "", "12345"},
		{[]uint32{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
		{[]uint32{1, 2, 3, 4, 5}, "}{", "1}{2}{3}{4}{5"},
		{[]uint32{1, 2}, ",", "1,2"},
		{[]uint32{1}, ",", "1"},
		{[]uint32{}, ",", ""},
	}
	for _, c := range cases {
		got := JoinUint32s(c.in, c.sep)
		if got != c.out {
			t.Errorf("JoinUint32s(%v, %q) == %q, out %q", c.in, c.sep, got, c.out)
		}
	}
}

func TestJoinUint64s(t *testing.T) {
	cases := []struct {
		in       []uint64
		sep, out string
	}{
		{[]uint64{1, 2, 3, 4, 5}, "", "12345"},
		{[]uint64{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
		{[]uint64{1, 2, 3, 4, 5}, "}{", "1}{2}{3}{4}{5"},
		{[]uint64{1, 2}, ",", "1,2"},
		{[]uint64{1}, ",", "1"},
		{[]uint64{}, ",", ""},
	}
	for _, c := range cases {
		got := JoinUint64s(c.in, c.sep)
		if got != c.out {
			t.Errorf("JoinUint64s(%v, %q) == %q, out %q", c.in, c.sep, got, c.out)
		}
	}
}
