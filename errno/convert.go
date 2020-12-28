// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package errno

import (
	"errors"
)

var (
	// InternalServerError 系统错误，非 Errno 实现均被 Parse 解析为系统错误
	InternalServerError = &errno{code: -1, message: "System Error"}

	// OK error 为nil 时会被 Parse 解析为 OK
	OK = &errno{code: 0, message: "ok"}
)

// To 把任意 error 转成 Errno 类型
// 如果 err 实现了 Errno 接口，则返回 err 本身
// 如果 err 为 nil， 返回预定义 OK， 默认 code=0, message=ok
// 如果 err 没有实现 Errno，返回预定义 InternalServerError，并使用 WrapErrno 包装 err，默认 code=-1, message=System Error
func To(err error) Errno {
	if err == nil {
		return OK
	}

	switch e := err.(type) {
	case Errno:
		return e
	}

	return Wrap(InternalServerError).Err(err)
}

// Parse 返回错误码、错误消息以及包装错误
// Parse 首先使用 To(err) 获得一个 Errno 实现
// 包装错误使用  errors.Unwrap 解析获得
func Parse(err error) (int, string, error) {
	w := To(err)
	return w.Code(), w.Message(), errors.Unwrap(w)
}
