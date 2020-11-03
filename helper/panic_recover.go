// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package helper

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/thinkeridea/go-extend/exbytes"
)

const k = 1 << 10

// PanicRecover 帮助把 panic 转为 error返回，并获取堆栈信息打印日志。
// 该方法会把错误信息打印到标准错误输出，并包含一段堆栈信息，帮助我们快速查找程序问题。
func PanicRecover(r interface{}) error {
	loggerStderr := log.New(os.Stderr, "", log.LstdFlags)

	if r != nil {
		buf := make([]byte, 4*k)
		n := runtime.Stack(buf, false)
		loggerStderr.Printf("[Recovery] panic recovered:\n%v\n%s\n", r, exbytes.ToString(buf[:n]))

		var err error
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = fmt.Errorf("%v", r)
		}

		return err
	}

	return nil
}
