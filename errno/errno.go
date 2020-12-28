// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package errno

import (
	"fmt"
)

// Errno 包含错误码和错误消息的接口，适用于应用程序错误
type Errno interface {
	error

	// Code 返回错误编码
	Code() int

	// Message 返回错误消息文本
	Message() string
}

// errno 包含错误码和错误消息，适用于应用程序错误
type errno struct {
	code    int
	message string
}

// Code 返回错误编码
func (e *errno) Code() int {
	return e.code
}

// Message 返回错误消息文本
func (e *errno) Message() string {
	return e.message
}

// Error error 接口实现，格式化 Errno 信息
func (e *errno) Error() string {
	return fmt.Sprintf("Error - code: %d, message: %s", e.code, e.message)
}

// New 给定错误码及错误格式化文本，返回一个 Errno
// 每 New 一次返回一个新的 Errno， 即使 code 和 message 相同
func New(code int, message string) Errno {
	return &errno{code: code, message: message}
}

// NewCode 给定错误码，返回一个 Errno
// 每 New 一次返回一个新的 Errno， 即使 code 相同
func NewCode(code int) Errno {
	return &errno{code: code}
}
