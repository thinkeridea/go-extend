// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package datalog

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewRecord(t *testing.T) {
	r := NewRecord(3)

	if len(r) != 3 {
		t.Errorf("len(r)=%d, want %d", len(r), 3)
	}

	if cap(r) != 3 {
		t.Errorf("cap(r)=%d, want %d", cap(r), 3)
	}
}

func TestNewRecordPool(t *testing.T) {
	pool := NewRecordPool(3)
	r := pool.Get().(Record)

	if len(r) != 3 {
		t.Errorf("len(r)=%d, want %d", len(r), 3)
	}

	if cap(r) != 3 {
		t.Errorf("cap(r)=%d, want %d", cap(r), 3)
	}
}

func TestRecord_ToBytes(t *testing.T) {
	r := NewRecord(3)
	r[0] = "0"
	r[1] = "1"
	r[2] = "2"

	data := r.ToBytes(FieldSep, NewLine)
	if string(data) != strings.Join([]string(r), FieldSep)+NewLine {
		t.Errorf("r:(%q).ToBytes(%s, %s) = %q, want %q", r, FieldSep, NewLine, data, []byte(strings.Join([]string(r), FieldSep)+NewLine))
	}

	var r0, r2 string
	r0 = "0" + FieldSep
	r2 = "2" + NewLine

	r[0] = r0
	r[1] = "1" + "\x03"
	r[2] = r2

	data = r.ToBytes(FieldSep, NewLine)
	if string(data) != strings.Join([]string{"0 ", "1\x03", "2 "}, FieldSep)+NewLine {
		t.Errorf("r:(%q).ToBytes(%s, %s) = %q, want %q", r, FieldSep, NewLine, data, []byte(strings.Join([]string{"0", "1\x03", "2"}, FieldSep)+NewLine))
	}

	// 引用数据不发生改变
	if r0 != "0"+FieldSep {
		t.Errorf("r0 = %q, want %q", r0, "0"+FieldSep)
	}

	if r2 != "2"+NewLine {
		t.Errorf("r2 = %q, want %q", r2, "2"+NewLine)
	}
}

func TestRecord_UnsafeToBytes(t *testing.T) {
	r := NewRecord(3)
	r[0] = "0"
	r[1] = "1"
	r[2] = "2"

	data := r.UnsafeToBytes(FieldSep, NewLine)
	if string(data) != strings.Join([]string(r), FieldSep)+NewLine {
		t.Errorf("r:(%q).UnsafeToBytes(%s, %s) = %q, want %q", r, FieldSep, NewLine, data, []byte(strings.Join([]string(r), FieldSep)+NewLine))
	}

	var r0, r2 string
	r0 = fmt.Sprint("0" + FieldSep)
	r2 = fmt.Sprint("2" + NewLine)

	r[0] = r0
	r[1] = fmt.Sprint("1" + "\x03")
	r[2] = r2

	data = r.UnsafeToBytes(FieldSep, NewLine)
	if string(data) != strings.Join([]string{"0 ", "1\x03", "2 "}, FieldSep)+NewLine {
		t.Errorf("r:(%q).UnsafeToBytes(%s, %s) = %q, want %q", r, FieldSep, NewLine, data, []byte(strings.Join([]string{"0", "1\x03", "2"}, FieldSep)+NewLine))
	}

	// 引用数据发生改变
	if r0 != "0 " {
		t.Errorf("r0 = %q, want %q", r0, "0 ")
	}

	if r2 != "2 \n" { // 多个字符串替换为一个字符时原字符串长度未发生改变，可能会出现错乱
		t.Errorf("r2 = %q, want %q", r2, "2 \n")
	}
}

func TestRecord_Join(t *testing.T) {
	r := NewRecord(3)
	r[0] = "0"
	r[1] = "1"
	r[2] = "2"

	data := r.Join(FieldSep, NewLine)
	if string(data) != strings.Join([]string(r), FieldSep)+NewLine {
		t.Errorf("r:(%q).Join(%s, %s) = %q, want %q", r, FieldSep, NewLine, data, []byte(strings.Join([]string(r), FieldSep)+NewLine))
	}

	data = Record{}.Join(FieldSep, NewLine)
	if string(data) != NewLine {
		t.Errorf("r([]).Join(%s,%s) = %q, want %q", FieldSep, NewLine, data, []byte(NewLine))
	}
}

func TestRecord_Clean(t *testing.T) {
	r := NewRecord(3)
	r[0] = "0"
	r[1] = "1"
	r[2] = "2"

	r.Clean()

	if r[0] != "" || r[1] != "" || r[2] != "" {
		t.Errorf("r = %q, want %q", r, make([]string, 3))
	}
}

func TestRecord_ArrayJoin(t *testing.T) {
	r := NewRecord(3)
	r[0] = "0"
	r[1] = "1"
	r[2] = "2"

	data := r.ArrayJoin(ArraySep)
	if data != strings.Join(r, ArraySep) {
		t.Errorf("r:(%q).ArrayJoin(%s) = %q, want %q", r, FieldSep, data, strings.Join(r, ArraySep))
	}
}

func TestRecord_ArrayFieldJoin(t *testing.T) {
	r := NewRecord(3)
	r[0] = "0"
	r[1] = "1"
	r[2] = "2"

	data := r.ArrayFieldJoin(ArrayFieldSep, ArraySep)
	if string(data) != strings.Join([]string(r), ArrayFieldSep) {
		t.Errorf("r:(%q).ArrayFieldJoin(%s, %s) = %q, want %q", r, ArrayFieldSep, ArraySep, data, []byte(strings.Join([]string(r), ArrayFieldSep)))
	}

	var r0, r2 string
	r0 = "0" + ArrayFieldSep
	r2 = "2" + ArraySep

	r[0] = r0
	r[1] = "1"
	r[2] = r2

	data = r.ArrayFieldJoin(ArrayFieldSep, ArraySep)
	if string(data) != strings.Join([]string{"0 ", "1", "2 "}, ArrayFieldSep) {
		t.Errorf("r:(%q).ArrayFieldJoin(%s, %s) = %q, want %q", r, ArrayFieldSep, ArraySep, data, []byte(strings.Join([]string{"0 ", "1", "2 "}, ArrayFieldSep)))
	}

	// 引用数据不发生改变
	if r0 != "0"+ArrayFieldSep {
		t.Errorf("r0 = %q, want %q", r0, "0"+ArrayFieldSep)
	}

	if r2 != "2"+ArraySep {
		t.Errorf("r2 = %q, want %q", r2, "2"+ArraySep)
	}
}

func TestRecord_UnsafeArrayFieldJoin(t *testing.T) {
	r := NewRecord(3)
	r[0] = "0"
	r[1] = "1"
	r[2] = "2"

	data := r.UnsafeArrayFieldJoin(ArrayFieldSep, ArraySep)
	if string(data) != strings.Join([]string(r), ArrayFieldSep) {
		t.Errorf("r:(%q).UnsafeArrayFieldJoin(%s, %s) = %q, want %q", r, ArrayFieldSep, ArraySep, data, []byte(strings.Join([]string(r), ArrayFieldSep)))
	}

	var r0, r2 string
	r0 = fmt.Sprint("0" + ArrayFieldSep)
	r2 = fmt.Sprint("2" + ArraySep)

	r[0] = r0
	r[1] = "1"
	r[2] = r2

	data = r.UnsafeArrayFieldJoin(ArrayFieldSep, ArraySep)
	if string(data) != strings.Join([]string{"0 ", "1", "2 "}, ArrayFieldSep) {
		t.Errorf("r:(%q).UnsafeArrayFieldJoin(%s, %s) = %q, want %q", r, ArrayFieldSep, ArraySep, data, []byte(strings.Join([]string{"0 ", "1", "2 "}, ArrayFieldSep)))
	}

	// 引用数据不发生改变
	if r0 != "0 " {
		t.Errorf("r0 = %q, want %q", r0, "0 ")
	}

	if r2 != "2 " {
		t.Errorf("r2 = %q, want %q", r2, "2 ")
	}
}
