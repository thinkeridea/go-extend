package errno

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func wrapEqual(t *testing.T, we1, we2 WrapErrno, err error) {
	w1, ok := we1.(*wrap)
	if !ok {
		t.Error("we1 is not an instance of *wrap")
	}

	w2, ok := we2.(*wrap)
	if !ok {
		t.Error("we2 is not an instance of *wrap")
	}

	if w1.errno != w2.errno {
		t.Errorf("w1.errno:%v != w2.errno:%v", w1.errno, w2.errno)
	}

	if w1.comment != w2.comment {
		t.Errorf("w1.comment:%v != w2.comment:%v", w1.comment, w2.comment)
	}

	if err != nil && !errors.Is(w1, err) {
		t.Errorf("errors.Is(err:%v, v.err:%v) != true", err, err)
	}

	if !reflect.DeepEqual(w1.format, w2.format) {
		t.Errorf("w1.format:%v != w2.format:%v", w1.format, w2.format)
	}
}

func TestWrap(t *testing.T) {
	err := &errno{code: -1, message: "System Error"}
	w := Wrap(err)
	wrapEqual(t, w, &wrap{errno: *err}, err)

	ow := &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w = Wrap(ow)
	if ow == w {
		t.Errorf("multiple Wrap returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, nil)

	ow = &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w = Wrap(ow)
	if ow == w {
		t.Errorf("multiple Wrap returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, ow)
}

func TestWrapErr(t *testing.T) {
	we := (*wrapError)(nil)
	wErr := errors.New("system error")
	err := &errno{code: -1, message: "System Error"}
	w := WrapErr(err, wErr)
	wrapEqual(t, w, &wrap{errno: *err}, err)
	wrapEqual(t, w, &wrap{errno: *err}, wErr)

	nErr := errors.New("wrap new err")
	ow := &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w = WrapErr(ow, nErr)
	if ow == w {
		t.Errorf("multiple WrapErr returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, ow)
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, nErr)

	ow = &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}, err: we.Wrap(wErr)}
	w = WrapErr(ow, nErr)
	if ow == w {
		t.Errorf("multiple WrapErr returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, wErr)
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, nErr)

	ow = &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}, err: we.Wrap(wErr)}
	w = WrapErr(ow, ow)
	if ow == w {
		t.Errorf("multiple WrapErr returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, wErr)
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, ow)
}

func TestWrapComment(t *testing.T) {
	err := &errno{code: -1, message: "System Error"}
	w := WrapComment(err, "wrap comment")
	wrapEqual(t, w, &wrap{errno: *err, comment: "wrap comment"}, err)

	ow := &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w = WrapComment(ow, "wrap comment")
	if ow == w {
		t.Errorf("multiple WrapComment returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "wrap comment", format: []interface{}{1, 10, "format"}}, nil)

	ow = &wrap{errno: *err, format: []interface{}{1, 10, "format"}}
	w = WrapComment(ow, "wrap comment")
	if ow == w {
		t.Errorf("multiple WrapComment returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "wrap comment", format: []interface{}{1, 10, "format"}}, nil)
}

func TestWrapFormat(t *testing.T) {
	err := &errno{code: -1, message: "System Error"}
	w := WrapFormat(err, 1, 10, "format")
	wrapEqual(t, w, &wrap{errno: *err, format: []interface{}{1, 10, "format"}}, err)

	ow := &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w = WrapFormat(ow, 5, 8, "test format")
	if ow == w {
		t.Errorf("multiple WrapFormat returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{5, 8, "test format"}}, nil)

	ow = &wrap{errno: *err, comment: "system error comment"}
	w = WrapFormat(ow, 5, 8, "test format")
	if ow == w {
		t.Errorf("multiple WrapFormat returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{5, 8, "test format"}}, nil)
}

func TestWrap_WrapErr(t *testing.T) {
	we := (*wrapError)(nil)
	wErr := errors.New("system error")
	err := &errno{code: -1, message: "System Error"}

	nErr := errors.New("wrap new err")
	ow := &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w := ow.WrapErr(nErr)
	if ow == w {
		t.Errorf("multiple WrapErr returns the same instance %p == %p", ow, w)
		wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, err)
	}

	ow = &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}, err: we.Wrap(wErr)}
	w = ow.WrapErr(nErr)
	if ow == w {
		t.Errorf("multiple WrapErr returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, nErr)

	ow = &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}, err: we.Wrap(wErr)}
	w = ow.WrapErr(ow)
	if ow == w {
		t.Errorf("multiple WrapErr returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: *err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, ow)
}

func TestWrap_WrapComment(t *testing.T) {
	err := errno{code: -1, message: "System Error"}
	ow := &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w := ow.WrapComment("wrap comment")
	if ow == w {
		t.Errorf("multiple wrap.WrapComment returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "wrap comment", format: []interface{}{1, 10, "format"}}, nil)

	ow = &wrap{errno: err, format: []interface{}{1, 10, "format"}}
	w = ow.WrapComment("wrap comment")
	if ow == w {
		t.Errorf("multiple wrap.WrapComment returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "wrap comment", format: []interface{}{1, 10, "format"}}, nil)
}

func TestWrap_WrapFormat(t *testing.T) {
	err := errno{code: -1, message: "System Error"}
	ow := &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w := ow.WrapFormat(5, 8, "test format")
	if ow == w {
		t.Errorf("multiple wrap.WrapFormat returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "system error comment", format: []interface{}{5, 8, "test format"}}, nil)

	ow = &wrap{errno: err, comment: "system error comment"}
	w = ow.WrapFormat(5, 8, "test format")
	if ow == w {
		t.Errorf("multiple wrap.WrapFormat returns the same instance %p == %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "system error comment", format: []interface{}{5, 8, "test format"}}, nil)
}

func TestWrap_Err(t *testing.T) {
	we := (*wrapError)(nil)
	wErr := errors.New("system error")
	err := errno{code: -1, message: "System Error"}
	nErr := errors.New("wrap new err")
	ow := &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w := ow.Err(nErr)
	if ow != w {
		t.Errorf("multiple wrap.Err returns the different instance %p != %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, ow)
	wrapEqual(t, w, &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, nErr)

	owe := we.Wrap(wErr)
	ow = &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}, err: owe}
	w = ow.Err(nErr)
	if ow != w {
		t.Errorf("multiple wrap.Err returns the different instance %p != %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, nErr)
	wrapEqual(t, w, &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, wErr)

	ow = &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}, err: owe}
	w = ow.Err(ow)
	if ow != w {
		t.Errorf("multiple wrap.Err returns the different instance %p != %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}}, ow)
}

func TestWrap_Comment(t *testing.T) {
	err := errno{code: -1, message: "System Error"}
	ow := &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w := ow.Comment("wrap comment")
	if ow != w {
		t.Errorf("multiple wrap.Comment returns the different instance %p != %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "wrap comment", format: []interface{}{1, 10, "format"}}, nil)

	ow = &wrap{errno: err, format: []interface{}{1, 10, "format"}}
	w = ow.Comment("wrap comment")
	if ow != w {
		t.Errorf("multiple wrap.Comment returns the different instance %p != %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "wrap comment", format: []interface{}{1, 10, "format"}}, nil)
}

func TestWrap_Format(t *testing.T) {
	err := errno{code: -1, message: "System Error"}
	ow := &wrap{errno: err, comment: "system error comment", format: []interface{}{1, 10, "format"}}
	w := ow.Format(5, 8, "test format")
	if ow != w {
		t.Errorf("multiple wrap.WrapFormat returns the different instance %p != %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "system error comment", format: []interface{}{5, 8, "test format"}}, nil)

	ow = &wrap{errno: err, comment: "system error comment"}
	w = ow.Format(5, 8, "test format")
	if ow != w {
		t.Errorf("multiple wrap.WrapFormat returns the different instance %p != %p", ow, w)
	}
	wrapEqual(t, w, &wrap{errno: err, comment: "system error comment", format: []interface{}{5, 8, "test format"}}, nil)
}

func TestWrap_Message(t *testing.T) {
	e1 := Wrap(New(100, "1"))
	e2 := Wrap(New(100, "[%d] %s"))
	testCases := []struct {
		err  WrapErrno
		want string
	}{
		{e1, "1"},
		{e2, "[%d] %s"},
		{WrapComment(e1, "c1"), "1: c1"},
		{WrapFormat(e2, 1, "e2"), "[1] e2"},
		{WrapFormat(e2, 1, "e2").Comment("c2"), "[1] e2: c2"},
		{Wrap(NewCode(100)), ""},
		{Wrap(NewCode(100)).Comment("c3"), "c3"},
	}
	for _, tc := range testCases {
		if got := tc.err.Message(); got != tc.want {
			t.Errorf("(%v).Message() = %v, want %v", tc.err, got, tc.want)
		}
	}
}

func TestWrap_Error(t *testing.T) {
	e1 := Wrap(New(100, "1"))
	e2 := Wrap(New(100, "[%d] %s"))
	testCases := []struct {
		err  WrapErrno
		want string
	}{
		{e1, fmt.Sprintf("Error - code: %d, message: %s", 100, "1")},
		{e2, fmt.Sprintf("Error - code: %d, message: %s", 100, "[%d] %s")},
		{WrapComment(e1, "c1"), fmt.Sprintf("Error - code: %d, message: %s, comment: %s", 100, "1", "c1")},
		{WrapFormat(e2, 1, "e2"), fmt.Sprintf("Error - code: %d, message: [%d] %s", 100, 1, "e2")},
		{WrapFormat(e2, 1, "e2").Comment("c2"), fmt.Sprintf("Error - code: %d, message: [%d] %s, comment: %s", 100, 1, "e2", "c2")},
		{Wrap(NewCode(100)), fmt.Sprintf("Error - code: %d, message: %s", 100, "")},
		{Wrap(NewCode(100)).Comment("c3"), fmt.Sprintf("Error - code: %d, message: %s, comment: %s", 100, "", "c3")},

		{WrapComment(e1, "c1").Err(errors.New("err1")), fmt.Sprintf("Error - code: %d, message: %s, comment: %s, error: %s", 100, "1", "c1", "err1")},
		{WrapFormat(e2, 1, "e2").Err(errors.New("err1")), fmt.Sprintf("Error - code: %d, message: [%d] %s, error: %s", 100, 1, "e2", "err1")},
		{WrapFormat(e2, 1, "e2").Comment("c2").Err(errors.New("err1")), fmt.Sprintf("Error - code: %d, message: [%d] %s, comment: %s, error: %s", 100, 1, "e2", "c2", "err1")},
		{Wrap(NewCode(100)).Err(errors.New("err1")), fmt.Sprintf("Error - code: %d, message: %s, error: %s", 100, "", "err1")},
		{Wrap(NewCode(100)).Comment("c3").Err(errors.New("err1")), fmt.Sprintf("Error - code: %d, message: %s, comment: %s, error: %s", 100, "", "c3", "err1")},

		{WrapComment(e1, "c1").Err(errors.New("err1")).Err(errors.New("err2")), fmt.Sprintf("Error - code: %d, message: %s, comment: %s, error: %s", 100, "1", "c1", "err1, err2")},
		{WrapFormat(e2, 1, "e2").Err(errors.New("err1")).Err(errors.New("err2")), fmt.Sprintf("Error - code: %d, message: [%d] %s, error: %s", 100, 1, "e2", "err1, err2")},
		{WrapFormat(e2, 1, "e2").Comment("c2").Err(errors.New("err1")).Err(errors.New("err2")), fmt.Sprintf("Error - code: %d, message: [%d] %s, comment: %s, error: %s", 100, 1, "e2", "c2", "err1, err2")},
		{Wrap(NewCode(100)).Err(errors.New("err1")).Err(errors.New("err2")), fmt.Sprintf("Error - code: %d, message: %s, error: %s", 100, "", "err1, err2")},
		{Wrap(NewCode(100)).Comment("c3").Err(errors.New("err1")).Err(errors.New("err2")), fmt.Sprintf("Error - code: %d, message: %s, comment: %s, error: %s", 100, "", "c3", "err1, err2")},
	}
	for _, tc := range testCases {
		if got := tc.err.Error(); got != tc.want {
			t.Errorf("(%v).Error() = %v, want %v", tc.err, got, tc.want)
		}
	}
}

func TestWrap_Is(t *testing.T) {
	var w *wrap
	err1 := errors.New("1")
	err2 := errors.New("2")
	err3 := errors.New("3")
	err4 := fmt.Errorf("4 %w", err3)
	var uncomparable errorUncomparable

	err := New(100, "errno")
	testCases := []struct {
		err    WrapErrno
		target error
		match  bool
	}{
		{nil, nil, true},
		{w, err, false},
		{Wrap(err), nil, false},
		{Wrap(err), err, true},
		{Wrap(err), err1, false},
		{Wrap(err), err2, false},
		{Wrap(err).WrapErr(err1), err1, true},
		{Wrap(err).Err(err1), err1, true},
		{Wrap(err).Err(err1), err, true},
		{Wrap(err).Err(err1).Err(err2), err1, true},
		{Wrap(err).Err(err1).Err(err2), err, true},
		{Wrap(err).Err(err1).Err(err2), err2, true},
		{Wrap(err).Err(err1).Err(err2), err3, false},
		{Wrap(err).Err(err2).WrapErr(err3), err3, true},
		{Wrap(err).Err(err2).WrapErr(err4), err3, true},
		{Wrap(err).Err(err2).WrapErr(err3), nil, false},
		{Wrap(err).Err(err2).WrapErr(err3), errors.New("not found"), false},
		{Wrap(err).Err(uncomparable), uncomparable, true},
		{Wrap(err).Err(uncomparable), &uncomparable, false},
		{Wrap(err).Err(&uncomparable), uncomparable, true},
		{Wrap(err).Err(&uncomparable), &uncomparable, true},
		{Wrap(err).Err(uncomparable), err1, false},
		{Wrap(err).Err(&uncomparable), err1, false},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if tc.err != nil {
				w := tc.err.(*wrap)
				if got := w.Is(tc.target); got != tc.match {
					t.Errorf("tc.err(%v:%p).Is(%v:%p) = %v, want %v", tc.err, tc.err, tc.target, tc.target, got, tc.match)
				}
			}

			if got := errors.Is(tc.err, tc.target); got != tc.match {
				t.Errorf("errors.Is(%v, %v) = %v, want %v", tc.err, tc.target, got, tc.match)
			}
		})
	}
}

func TestWrap_As(t *testing.T) {
	var w *wrap
	var errno Errno
	var errT errorT
	var errP *os.PathError
	var timeout interface{ Timeout() bool }
	var p *poser
	_, errF := os.Open("non-existing")
	poserErr := &poser{"oh no", nil}
	err := New(100, "errno")

	testCases := []struct {
		err    WrapErrno
		target interface{}
		match  bool
		want   interface{} // value of target on match
	}{{
		nil,
		&errno,
		false,
		nil,
	}, {
		w,
		&errT,
		false,
		nil,
	}, {
		Wrap(err).Err(errF),
		&errP,
		true,
		errF,
	}, {
		Wrap(err).Err(errorT{"T"}),
		&errP,
		false,
		nil,
	}, {
		Wrap(err).Err(nil),
		&errT,
		false,
		nil,
	}, {
		Wrap(err).Err(&poser{"error", nil}),
		&errT,
		true,
		errorT{"poser"},
	}, {
		Wrap(err).Err(&poser{"path", nil}),
		&errP,
		true,
		poserPathErr,
	}, {
		Wrap(err).Err(poserErr).Err(&errT),
		&p,
		true,
		poserErr,
	}, {
		Wrap(err).Err(errors.New("err")),
		&timeout,
		false,
		nil,
	}, {
		Wrap(err).Err(errF).Err(&errT),
		&timeout,
		true,
		errF,
	}, {
		Wrap(err).Err(errF).Err(&errT),
		&timeout,
		true,
		errF,
	}}
	for i, tc := range testCases {
		name := fmt.Sprintf("%d:As(Errorf(..., %v), %v)", i, tc.err, tc.target)
		// Clear the target pointer, in case it was set in a previous test.
		rtarget := reflect.ValueOf(tc.target)
		rtarget.Elem().Set(reflect.Zero(reflect.TypeOf(tc.target).Elem()))
		t.Run(name, func(t *testing.T) {
			if tc.err != nil {
				w := tc.err.(*wrap)
				match := w.As(tc.target)
				if match != tc.match {
					t.Fatalf("match: got %v; want %v", match, tc.match)
				}
				if !match {
					return
				}
				if got := rtarget.Elem().Interface(); got != tc.want {
					t.Fatalf("got %#v, want %#v", got, tc.want)
				}
			}

			{
				match := errors.As(tc.err, tc.target)
				if match != tc.match {
					t.Fatalf("match: got %v; want %v", match, tc.match)
				}
				if !match {
					return
				}
				if got := rtarget.Elem().Interface(); got != tc.want {
					t.Fatalf("got %#v, want %#v", got, tc.want)
				}
			}
		})
	}
}

func TestWrap_Unwrap(t *testing.T) {
	var w *wrap
	err := New(100, "errno")
	err1 := errors.New("1")
	err2 := errors.New("2")

	testCases := []struct {
		err  WrapErrno
		want error
	}{
		{nil, nil},
		{w, nil},
		{&wrap{}, nil},
		{Wrap(err), err},
		{Wrap(err).Err(err1), err1},
		{Wrap(err).Err(err2), err2},
		{Wrap(err).Err(err1).Err(err2), err2},
		{Wrap(err).Err(err1).Err(nil), err1},
		{Wrap(err).Err(err1).Err(err), err},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			err := errors.Unwrap(tc.err)
			if got := errors.Unwrap(err); got != tc.want {
				t.Errorf("Unwrap(%v) = %v, want %v", tc.err, got, tc.want)
			}
			if tc.err != nil {
				w := tc.err.(*wrap)
				if got := errors.Unwrap(w.Unwrap()); got != tc.want {
					t.Errorf("Unwrap(%v) = %v, want %v", tc.err, got, tc.want)
				}
			}
		})
	}
}
