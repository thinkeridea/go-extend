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

func (w *wrapError) Error() string {
	if w == nil {
		return ""
	}
	return w.msg
}

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

func (w *wrapError) As(target interface{}) bool {
	for w != nil {
		if errors.As(w.new, target) {
			return true
		}

		w = w.old
	}

	return false
}

func (w *wrapError) Unwrap() error {
	if w == nil {
		return nil
	}
	return w.new
}

var errorType = reflect.TypeOf((*error)(nil)).Elem()
