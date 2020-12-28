package errno

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestTo(t *testing.T) {
	systemErr := errors.New("system error")
	wrapE := errors.New("wrap error")
	for _, v := range []struct {
		err     error
		code    int
		message string
		wrapErr error
	}{
		{systemErr, -1, "System Error", systemErr},
		{systemErr, -1, "System Error", InternalServerError},
		{os.ErrNotExist, -1, "System Error", os.ErrNotExist},
		{bytes.ErrTooLarge, -1, "System Error", bytes.ErrTooLarge},
		{nil, 0, "ok", OK},
		{New(10, "testing"), 10, "testing", nil},
		{NewCode(20), 20, "", nil},
		{Wrap(New(10, "testing")).Comment("error comment"), 10, "testing: error comment", nil},
		{Wrap(New(10, "testing")).Err(wrapE), 10, "testing", wrapE},
		{Wrap(New(10, "testing %s")).Err(wrapE).Format("format"), 10, "testing format", wrapE},
	} {
		err := To(v.err)
		if err.Code() != v.code {
			t.Errorf("To(v.err).Code():%d != v.code:%d", err.Code(), v.code)
		}

		if err.Message() != v.message {
			t.Errorf("To(v.err).Message():%s != v.message:%s", err.Message(), v.message)
		}

		if v.wrapErr != nil && !errors.Is(err, v.wrapErr) {
			t.Errorf("errors.Is(err, v.wrapErr:%v) != true", v.wrapErr)
		}
	}
}

func TestParse(t *testing.T) {
	systemErr := errors.New("system error")
	wrapE := errors.New("wrap error")
	for _, v := range []struct {
		err     error
		code    int
		message string
		wrapErr error
	}{
		{systemErr, -1, "System Error", systemErr},
		{systemErr, -1, "System Error", InternalServerError},
		{os.ErrNotExist, -1, "System Error", os.ErrNotExist},
		{bytes.ErrTooLarge, -1, "System Error", bytes.ErrTooLarge},
		{nil, 0, "ok", nil},
		{New(10, "testing"), 10, "testing", nil},
		{NewCode(20), 20, "", nil},
		{Wrap(New(10, "testing")).Comment("error comment"), 10, "testing: error comment", nil},
		{Wrap(New(10, "testing")).Err(wrapE), 10, "testing", wrapE},
		{Wrap(New(10, "testing %s")).Err(wrapE).Format("format"), 10, "testing format", wrapE},
	} {
		t.Run("", func(t *testing.T) {
			code, message, err := Parse(v.err)
			if code != v.code {
				t.Errorf("code:%d != v.code:%d", code, v.code)
			}

			if message != v.message {
				t.Errorf("message:%s != v.message:%s", message, v.message)
			}

			if v.wrapErr != nil && !errors.Is(err, v.wrapErr) {
				t.Errorf("errors.Is(err:%v, v.wrapErr:%v) != true", err, v.wrapErr)
			}

			if _, ok := v.wrapErr.(Errno); !ok && v.wrapErr != nil && errors.Unwrap(err) != v.wrapErr {
				t.Errorf("errors.Unwrap(err:%v)!=v.wrapErr:%v", err, v.wrapErr)
			}
		})
	}
}
