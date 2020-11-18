// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package extime

import (
	"testing"
	"time"
)

func TestParseInLocal(t *testing.T) {
	t1, err:=ParseInLocal("2006-01-02 15:04:05", "2020-02-03 04:05:06")
	if err != nil {
		t.Fatal(err)
	}

	t2 := time.Date(2020, 02, 03, 04, 05, 06, 0, time.Local)
	if t1 != t2 {
		t.Fatalf("ParseInLocal(2020-02-03 04:05:06) = %v, want %v", t1, t2)
	}

	if t1.Location() != time.Local {
		t.Fatalf("ParseInLocal(2020-02-03 04:05:06).Location() = _, %v, want _, %v", t1.Location(), time.Local)
	}

	local:= time.Local
	defer func() {
		time.Local = local
	}()

	time.Local, err = time.LoadLocation("Asia/Baghdad")
	if err != nil {
		t.Fatal(err)
	}

	t1, err = ParseInLocal("Jan 02 2006 MST", "Feb 01 2013 AST")
	if err != nil {
		t.Fatal(err)
	}

	_, offset := t1.Zone()
	// A zero offset means that ParseInLocation did not recognize the
	// 'AST' abbreviation as matching the current location (Baghdad,
	// where we'd expect a +03 hrs offset); likely because we're using
	// a recent tzdata release (2017a or newer).
	// If it happens, skip the Baghdad test.
	if offset != 0 {
		t2 = time.Date(2013, time.February, 1, 00, 00, 00, 0, time.Local)
		if t1 != t2 {
			t.Fatalf("ParseInLocal(Feb 01 2013 AST) = %v, want %v", t1, t2)
		}
		if offset != 3*60*60 {
			t.Fatalf("ParseInLocal(Feb 01 2013 AST).Zone = _, %d, want _, %d", offset, 3*60*60)
		}
	}

	time.Local, err = time.LoadLocation("America/Blanc-Sablon")
	if err != nil {
		t.Fatal(err)
	}

	// In this case 'AST' means 'Atlantic Standard Time', and we
	// expect the abbreviation to correctly match the american
	// location.
	t1, err = ParseInLocal("Jan 02 2006 MST", "Feb 01 2013 AST")
	if err != nil {
		t.Fatal(err)
	}
	t2 = time.Date(2013, time.February, 1, 00, 00, 00, 0, time.Local)
	if t1 != t2 {
		t.Fatalf("ParseInLocal(Feb 01 2013 AST) = %v, want %v", t1, t2)
	}
	_, offset = t1.Zone()
	if offset != -4*60*60 {
		t.Fatalf("ParseInLocal(Feb 01 2013 AST).Zone = _, %d, want _, %d", offset, -4*60*60)
	}
}
