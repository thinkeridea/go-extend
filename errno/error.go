package errno

import (
	"errors"
	"fmt"
	"reflect"
)

type wrapError struct {
	old *wrapError
	new error
	msg string
}

// Wrap 包装一个错误，这个辅助 wrap 实现错误包装使用的
// 当被包装的 err 为 nil 时会返回自身
// 当首次被包装且 err 为 Errno 实现，这不设置错误消息
// 使用 `old err, new err` 格式存储错误信息
func (w *wrapError) Wrap(err error) *wrapError {
	if err == nil {
		return w
	}

	// 第一层错误为 Errno，不设置msg
	if w == nil {
		if _, ok := err.(Errno); ok {
			return &wrapError{new: err}
		}
	}

	var msg string
	if w == nil || w.msg == "" {
		msg = fmt.Sprintf("%v", err)
	} else {
		msg = fmt.Sprintf("%s, %v", w.msg, err)
	}

	return &wrapError{old: w, new: err, msg: msg}
}

// Error 响应一个被格式化的错误文本，这是 error 接口的实现
func (w *wrapError) Error() string {
	if w == nil {
		return ""
	}
	return w.msg
}

// Is 这是 errors.Is 接口实现，可以通过 errors.Is 函数判断是否包含某个具体的错误
// 更详细的细节请浏览 errors.Is 文档
func (w *wrapError) Is(target error) bool {
	if w == target {
		return true
	}
	for w != nil {
		if errors.Is(w.new, target) {
			return true
		}

		w = w.old
	}

	return false
}

// As 这是 errors.As 接口实现，可以通过 errors.As 函数判断是否包含某个类型的错误，并把发现的第一个错误赋值给 target
// 更详细的细节请浏览 errors.As 文档
func (w *wrapError) As(target interface{}) bool {
	for w != nil {
		if errors.As(w.new, target) {
			return true
		}

		w = w.old
	}

	return false
}

// Unwrap 返回最后被包装的错误
// 这并不是 errors 中的 Unwrap 接口实现，而且用来获取最后绑定的错误
// 如果需要判断是否包含某个错误请使用 errors.Is
// 如果判断是否包含的某种类型的错误，请使用 errors.As
func (w *wrapError) Unwrap() error {
	if w == nil {
		return nil
	}
	return w.new
}

var errorType = reflect.TypeOf((*error)(nil)).Elem()
