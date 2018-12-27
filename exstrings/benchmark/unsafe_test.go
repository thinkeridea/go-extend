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

package benchmark

import (
	"strings"
	"testing"

	"github.com/thinkeridea/go-extend/exstrings"
)

func BenchmarkUnsafeRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exstrings.UnsafeRepeat("ABC", 100000)
	}
}

func BenchmarkStandardLibraryRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("ABC", 100000)
	}
}
