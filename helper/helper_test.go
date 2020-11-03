// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package helper

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestMust(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("no panicked")
			}
		}()

		err := errors.New("error")
		x := Must(nil, err)
		fmt.Println(x)
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panicked error:%v", r)
			}
		}()

		var expected *os.File
		actual := Must(expected, nil).(*os.File)
		if actual != expected {
			t.Errorf("actual:%v expected:%v", actual, expected)
		}
	}()
}
