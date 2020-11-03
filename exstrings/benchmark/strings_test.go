// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package benchmark

import (
	"strings"
	"testing"

	"github.com/thinkeridea/go-extend/exstrings"
)

func BenchmarkReplace(b *testing.B) {
	s := "acccbbbaacaabbcbbaaccbaaacaabaccacabbcaacbbccccbbbccaccbcaac"
	for i := 0; i < b.N; i++ {
		exstrings.Replace(s, "cc", "d", -1)
		exstrings.Replace(s, "aa", "d", -1)
		exstrings.Replace(s, "bb", "d", -1)
		exstrings.Replace(s, "ac", "d", -1)
		exstrings.Replace(s, "ca", "d", -1)
		exstrings.Replace(s, "bc", "d", -1)
		exstrings.Replace(s, "ba", "d", -1)
		exstrings.Replace(s, "acc", "d", -1)
		exstrings.Replace(s, "ccb", "d", -1)
		exstrings.Replace(s, "cbb", "d", -1)
		exstrings.Replace(s, "caa", "d", -1)
		exstrings.Replace(s, "bbc", "d", -1)
		exstrings.Replace(s, "aca", "d", -1)
		exstrings.Replace(s, "ccc", "d", -1)
		exstrings.Replace(s, "ab", "d", -1)
		exstrings.Replace(s, "dd", "d", -1)
	}
}

func BenchmarkReplaceToBytes(b *testing.B) {
	s := "acccbbbaacaabbcbbaaccbaaacaabaccacabbcaacbbccccbbbccaccbcaac"
	for i := 0; i < b.N; i++ {
		exstrings.ReplaceToBytes(s, "cc", "d", -1)
		exstrings.ReplaceToBytes(s, "aa", "d", -1)
		exstrings.ReplaceToBytes(s, "bb", "d", -1)
		exstrings.ReplaceToBytes(s, "ac", "d", -1)
		exstrings.ReplaceToBytes(s, "ca", "d", -1)
		exstrings.ReplaceToBytes(s, "bc", "d", -1)
		exstrings.ReplaceToBytes(s, "ba", "d", -1)
		exstrings.ReplaceToBytes(s, "acc", "d", -1)
		exstrings.ReplaceToBytes(s, "ccb", "d", -1)
		exstrings.ReplaceToBytes(s, "cbb", "d", -1)
		exstrings.ReplaceToBytes(s, "caa", "d", -1)
		exstrings.ReplaceToBytes(s, "bbc", "d", -1)
		exstrings.ReplaceToBytes(s, "aca", "d", -1)
		exstrings.ReplaceToBytes(s, "ccc", "d", -1)
		exstrings.ReplaceToBytes(s, "ab", "d", -1)
		exstrings.ReplaceToBytes(s, "dd", "d", -1)
	}
}

func BenchmarkUnsafeReplaceToBytes(b *testing.B) {
	s := "acccbbbaacaabbcbbaaccbaaacaabaccacabbcaacbbccccbbbccaccbcaac"
	for i := 0; i < b.N; i++ {
		exstrings.UnsafeReplaceToBytes(s, "cc", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "aa", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "bb", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "ac", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "ca", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "bc", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "ba", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "acc", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "ccb", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "cbb", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "caa", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "bbc", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "aca", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "ccc", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "ab", "d", -1)
		exstrings.UnsafeReplaceToBytes(s, "dd", "d", -1)
	}
}

func BenchmarkStandardLibraryReplace(b *testing.B) {
	s := "acccbbbaacaabbcbbaaccbaaacaabaccacabbcaacbbccccbbbccaccbcaac"
	for i := 0; i < b.N; i++ {
		strings.Replace(s, "cc", "d", -1)
		strings.Replace(s, "aa", "d", -1)
		strings.Replace(s, "bb", "d", -1)
		strings.Replace(s, "ac", "d", -1)
		strings.Replace(s, "ca", "d", -1)
		strings.Replace(s, "bc", "d", -1)
		strings.Replace(s, "ba", "d", -1)
		strings.Replace(s, "acc", "d", -1)
		strings.Replace(s, "ccb", "d", -1)
		strings.Replace(s, "cbb", "d", -1)
		strings.Replace(s, "caa", "d", -1)
		strings.Replace(s, "bbc", "d", -1)
		strings.Replace(s, "aca", "d", -1)
		strings.Replace(s, "ccc", "d", -1)
		strings.Replace(s, "ab", "d", -1)
		strings.Replace(s, "dd", "d", -1)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exstrings.Repeat("ABC", 100000)
	}
}

func BenchmarkRepeatToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exstrings.RepeatToBytes("ABC", 100000)
	}
}

func BenchmarkStandardLibraryRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("ABC", 100000)
	}
}

func BenchmarkJoin(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := 0; i < b.N; i++ {
		exstrings.Join(s, "-")
	}
}

func BenchmarkJoinToBytes(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := 0; i < b.N; i++ {
		exstrings.JoinToBytes(s, "-")
	}
}

func BenchmarkStandardLibraryJoin(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := 0; i < b.N; i++ {
		strings.Join(s, "-")
	}
}
