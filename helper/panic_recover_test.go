// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

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
