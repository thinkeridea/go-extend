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

package helper

import (
	"errors"
	"testing"
)

func TestPanicRecover(t *testing.T) {
	err := PanicRecover(recover())
	if err != nil {
		t.Errorf("panicked %v", err)
	}

	func() {
		defer func() {
			err := PanicRecover(recover())
			if err == nil {
				t.Errorf("no panicked for string test")
			}
		}()

		panic("test panic")
	}()

	func() {
		defer func() {
			err := PanicRecover(recover())
			if err == nil {
				t.Errorf("no panicked for error test")
			}
		}()

		panic(errors.New("test panic"))
	}()

	func() {
		defer func() {
			err := PanicRecover(recover())
			if err == nil {
				t.Errorf("no panicked default type test")
			}
		}()

		panic(12.4)
	}()
}
