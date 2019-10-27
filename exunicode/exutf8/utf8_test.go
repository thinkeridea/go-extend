package exutf8

import (
	"testing"
)

func TestRuneIndex(t *testing.T) {
	for _, tt := range []struct {
		in     string
		length int
		out    int
		ok     bool
	}{
		{"abcd", 3, 3, true},
		{"☺☻☹", 2, 6, true},
		{"☺☻☹", 3, 9, true},
		{"1,2,3,4", 5, 5, true},
		{"\xe2\x00", 2, 2, true},
		{"\xe2\x80", 1, 1, true},
		{"\xe2\x80", 2, 2, true},
		{"a\xe2\x80", 5, 3, false},
		{"golang", 5, 5, true},
		{"Go 语言", 4, 6, true},
		{"12345", 0, 0, true},
		{"12345", -1, 0, true},
	} {
		if out, ok := RuneIndex([]byte(tt.in), tt.length); out != tt.out || ok != tt.ok {
			t.Errorf("RuneIndex(%q, %d) = %d, %v, want %d, %v", tt.in, tt.length, out, ok, tt.out, tt.ok)
		}
	}
}

func TestRuneIndexInString(t *testing.T) {
	for _, tt := range []struct {
		in     string
		length int
		out    int
		ok     bool
	}{
		{"abcd", 3, 3, true},
		{"☺☻☹", 2, 6, true},
		{"☺☻☹", 3, 9, true},
		{"1,2,3,4", 5, 5, true},
		{"\xe2\x00", 2, 2, true},
		{"\xe2\x80", 1, 1, true},
		{"\xe2\x80", 2, 2, true},
		{"a\xe2\x80", 5, 3, false},
		{"golang", 5, 5, true},
		{"Go 语言", 4, 6, true},
		{"12345", 0, 0, true},
		{"12345", -1, 0, true},
	} {
		if out, ok := RuneIndexInString(tt.in, tt.length); out != tt.out || ok != tt.ok {
			t.Errorf("RuneIndexInString(%q, %d) = %d, %v, want %d, %v", tt.in, tt.length, out, ok, tt.out, tt.ok)
		}
	}
}
