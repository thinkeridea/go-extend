// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package errno

import (
	"fmt"
)

// WrapErrno 对 Errno 进行扩展
// 支持错误信息、描述信息、错误消息格式化
type WrapErrno interface {
	Errno

	// WrapErr 包装一个 error，返回新的 WrapErrno
	// 使用 fmt %w 标识实现多次包装
	WrapErr(err error) WrapErrno

	// WrapComment 包装一个评论，返回新的 WrapErrno
	// 这会覆盖旧 WrapErrno 中的评论信息，但不会修改旧 WrapErrno
	WrapComment(comment string) WrapErrno

	// WrapFormat 包装错误消息文本格式化参数，返回新的 WrapErrno
	// 在定义 Errno 错误消息文本时，可以使用 fmt 格式化标识，通过 WrapFormat 可以包装 fmt 格式化参数
	// 这会覆盖旧 WrapErrno 中的格式化参数，但不会修改旧 WrapErrno
	WrapFormat(a ...interface{}) WrapErrno

	// Err 使用 fmt %w 标识来包装现有错误，返回调用者自身
	// 这不会产生一个新的 WrapErrno， 但会修改包装的 error 数据，并使旧的 error 数据不丢失
	Err(err error) WrapErrno

	// Comment 设置评论信息，返回调用者自身
	// 这不会产生一个新的 WrapErrno，但是会覆盖评论信息
	Comment(comment string) WrapErrno

	// Format 设置错误消息文本格式化参数，返回调用者自身
	// 这不会产生一个新的 WrapErrno，但是会覆盖错误消息文本格式化参数
	Format(a ...interface{}) WrapErrno

	// Unwrap 返回 WrapErrno 中包装的错误对象，他并不一定是最后一个被包装错误对象，这取决具体实现
	// 这并不是 errors 中的 Unwrap 接口实现，而且用来获取 WrapErrno 中绑定的错误
	// 如果需要判断是否包含某个错误请使用 errors.Is
	// 如果判断是否包含的某种类型的错误，请使用 errors.As
	Unwrap() error
}

// WrapErrno 接口的实现
type wrap struct {
	errno
	err     *wrapError
	comment string
	format  []interface{}
}

// init 解析 Errno 并拷贝数据到 WrapErrno，返回调用者自身
func (w *wrap) init(err Errno) WrapErrno {
	if e, ok := err.(*wrap); ok {
		w.code = e.code
		w.message = e.message
		w.err = e.err
		w.comment = e.comment
		w.format = e.format
	} else {
		w.code = err.Code()
		w.message = err.Message()
	}

	// 第一层错误为 Errno
	if w.err == nil {
		w.err = w.err.Wrap(err)
	}

	return w
}

// WrapErr 包装一个 error，返回新的 WrapErrno
// 使用 fmt %w 标识实现多次包装
func (w *wrap) WrapErr(err error) WrapErrno {
	return (&wrap{}).init(w).Err(err)
}

// WrapComment 包装一个评论，返回新的 WrapErrno
// 这会覆盖旧 WrapErrno 中的评论信息，但不会修改旧 WrapErrno
func (w *wrap) WrapComment(comment string) WrapErrno {
	return (&wrap{}).init(w).Comment(comment)
}

// WrapFormat 包装错误消息文本格式化参数，返回新的 WrapErrno
// 在定义 Errno 错误消息文本时，可以使用 fmt 格式化标识，通过 WrapFormat 可以包装 fmt 格式化参数
// 这会覆盖旧 WrapErrno 中的格式化参数，但不会修改旧 WrapErrno
func (w *wrap) WrapFormat(a ...interface{}) WrapErrno {
	return (&wrap{}).init(w).Format(a...)
}

// Err 使用 fmt %w 标识来包装现有错误，返回调用者自身
// 这不会产生一个新的 WrapErrno， 但会修改包装的 error 数据，并使旧的 error 数据不丢失
func (w *wrap) Err(err error) WrapErrno {
	w.err = w.err.Wrap(err)
	return w
}

// Comment 设置评论信息，返回调用者自身
// 这不会产生一个新的 WrapErrno，但是会覆盖评论信息
func (w *wrap) Comment(comment string) WrapErrno {
	w.comment = comment
	return w
}

// Format 设置错误消息文本格式化参数，返回调用者自身
// 这不会产生一个新的 WrapErrno，但是会覆盖错误消息文本格式化参数
func (w *wrap) Format(a ...interface{}) WrapErrno {
	w.format = a
	return w
}

// Error error 接口实现，格式化 WrapErrno 信息
func (w *wrap) Error() string {
	message := w.message
	if w.format != nil {
		message = fmt.Sprintf(message, w.format...)
	}

	format := fmt.Sprintf("Error - code: %d, message: %s", w.code, message)
	if w.comment != "" {
		format += ", comment: " + w.comment
	}

	if w.err != nil {
		msg := w.err.Error()
		if msg != "" {
			format += ", error: " + msg
		}
	}

	return format
}

// Message 返回错误消息文本及评论信息
// 如果设置了 format 参数信息，将对错误消息文本进行 fmt 格式化
// 如果设置了评论信息返回格式化文本 `message: comment`
func (w *wrap) Message() string {
	message := w.message
	if w.format != nil {
		message = fmt.Sprintf(message, w.format...)
	}

	if w.comment == "" {
		return message
	}

	if w.message == "" {
		return w.comment
	}

	return fmt.Sprintf("%s: %s", message, w.comment)
}

// Unwrap 返回 WrapErrno 中包装的错误对象
// 这并不是 errors 中的 Unwrap 接口实现，而且用来获取 WrapErrno 中绑定的错误
// 如果需要判断是否包含某个错误请使用 errors.Is
// 如果判断是否包含的某种类型的错误，请使用 errors.As
func (w *wrap) Unwrap() error {
	if w == nil {
		return nil
	}
	return w.err
}

func (w *wrap) Is(target error) bool {
	if w == nil {
		return false
	}
	return w.err.Is(target)
}

func (w *wrap) As(target interface{}) bool {
	if w == nil {
		return false
	}
	return w.err.As(target)
}

// Wrap 转换 Errno 为 WrapErrno，返回一个新的 WrapErrno
// 即使 Errno 为 WrapErrno 也会得到一个新的 WrapErrno
func Wrap(e Errno) WrapErrno {
	return (&wrap{}).init(e)
}

// WrapErr 包装一个 error，返回一个新的 WrapErrno
// 使用 fmt %w 标识实现多次包装
func WrapErr(e Errno, err error) WrapErrno {
	return Wrap(e).Err(err)
}

// WrapComment 包装一个评论，返回新的 WrapErrno
// 这会覆盖旧 WrapErrno 中的评论信息，但不会修改旧 WrapErrno
func WrapComment(e Errno, comment string) WrapErrno {
	return Wrap(e).Comment(comment)
}

// WrapFormat 包装错误消息文本格式化参数，返回新的 WrapErrno
// 在定义 Errno 错误消息文本时，可以使用 fmt 格式化标识，通过 WrapFormat 可以包装 fmt 格式化参数
// 这会覆盖旧 WrapErrno 中的格式化参数，但不会修改旧 WrapErrno
func WrapFormat(e Errno, a ...interface{}) WrapErrno {
	return Wrap(e).Format(a...)
}
