package errno

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestWrapError_Wrap(t *testing.T) {
	err1 := errors.New("1")
	err2 := errors.New("2")
	err3 := errors.New("3")
	err4 := fmt.Errorf("4 %w", err3)

	var we *wrapError
	testCases := []struct {
		we  *wrapError
		err error
		msg string
	}{
		{nil, nil, ""},
		{nil, err1, "1"},
		{we.Wrap(err1), err2, "1, 2"},
		{we.Wrap(err1).Wrap(err2), err3, "1, 2, 3"},
		{we.Wrap(err1).Wrap(err2).Wrap(err3), err4, "1, 2, 3, 4 3"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			w := tc.we.Wrap(tc.err)
			if tc.err == nil {
				if we != w {
					t.Errorf("we.Wrap(nil) = %v, want %v", w, we)
				}
				return
			}

			if w.old != tc.we {
				t.Errorf("we.Wrap(tc.err).old = %v, want %v", w.old, tc.we)
			}

			if w.new != tc.err {
				t.Errorf("we.Wrap(tc.err).new = %v, want %v", w.new, tc.err)
			}

			if w.msg != tc.msg {
				t.Errorf("we.Wrap(tc.err).msg = %v, want %v", w.msg, tc.msg)
			}
		})
	}
}

func TestWrapError_Error(t *testing.T) {
	err1 := errors.New("1")
	err2 := errors.New("2")
	err3 := errors.New("3")
	err4 := fmt.Errorf("4 %w", err3)

	var we *wrapError
	testCases := []struct {
		we  *wrapError
		msg string
	}{
		{nil, ""},
		{we.Wrap(err1), "1"},
		{we.Wrap(err1).Wrap(err2), "1, 2"},
		{we.Wrap(err1).Wrap(err2).Wrap(err3), "1, 2, 3"},
		{we.Wrap(err1).Wrap(err2).Wrap(err3).Wrap(err4), "1, 2, 3, 4 3"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if tc.we.Error() != tc.msg {
				t.Errorf("tc.we.Error() = %v, want %v", tc.we.Error(), tc.msg)
			}
		})
	}
}

func TestWrapError_Is(t *testing.T) {
	err1 := errors.New("1")
	err2 := errors.New("2")
	err3 := errors.New("3")
	err4 := fmt.Errorf("4 %w", err3)
	var uncomparable errorUncomparable

	var we *wrapError
	testCases := []struct {
		err    *wrapError
		target error
		match  bool
	}{
		{nil, nil, false},
		{we.Wrap(nil), we, true},
		{we.Wrap(err1), err1, true},
		{we.Wrap(err2), err2, true},
		{we.Wrap(err2), err3, false},
		{we.Wrap(err4), err3, true},
		{we.Wrap(err1).Wrap(err2), err2, true},
		{we.Wrap(err1).Wrap(err2), err1, true},
		{we.Wrap(err1).Wrap(err2), err3, false},
		{we.Wrap(err1).Wrap(err4), err3, true},
		{we.Wrap(err1).Wrap(err2).Wrap(err3), err1, true},
		{we.Wrap(err1).Wrap(err2).Wrap(err3), err2, true},
		{we.Wrap(err1).Wrap(err2).Wrap(err3), err3, true},
		{we.Wrap(err1).Wrap(err2).Wrap(err4), err3, true},
		{we.Wrap(err1).Wrap(err2).Wrap(err3), nil, false},
		{we.Wrap(err1).Wrap(err2).Wrap(err3), errors.New("not found"), false},
		{we.Wrap(uncomparable), uncomparable, true},
		{we.Wrap(uncomparable), &uncomparable, false},
		{we.Wrap(&uncomparable), uncomparable, true},
		{we.Wrap(&uncomparable), &uncomparable, true},
		{we.Wrap(uncomparable), err1, false},
		{we.Wrap(&uncomparable), err1, false},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if got := tc.err.Is(tc.target); got != tc.match {
				t.Errorf("tc.err(%v:%p).Is(%v:%p) = %v, want %v", tc.err, tc.err, tc.target, tc.target, got, tc.match)
			}

			if got := errors.Is(tc.err, tc.target); got != tc.match {
				t.Errorf("errors.Is(%v, %v) = %v, want %v", tc.err, tc.target, got, tc.match)
			}
		})
	}
}

func TestWrapError_As(t *testing.T) {
	var we *wrapError
	var errT errorT
	var errP *os.PathError
	var timeout interface{ Timeout() bool }
	var p *poser
	_, errF := os.Open("non-existing")
	poserErr := &poser{"oh no", nil}

	testCases := []struct {
		err    *wrapError
		target interface{}
		match  bool
		want   interface{} // value of target on match
	}{{
		nil,
		&errP,
		false,
		nil,
	}, {
		we.Wrap(errorT{"T"}),
		&errT,
		true,
		errorT{"T"},
	}, {
		we.Wrap(errF),
		&errP,
		true,
		errF,
	}, {
		we.Wrap(errorT{"T"}),
		&errP,
		false,
		nil,
	}, {
		we.Wrap(nil),
		&errT,
		false,
		nil,
	}, {
		we.Wrap(&poser{"error", nil}),
		&errT,
		true,
		errorT{"poser"},
	}, {
		we.Wrap(&poser{"path", nil}),
		&errP,
		true,
		poserPathErr,
	}, {
		we.Wrap(poserErr).Wrap(&errT),
		&p,
		true,
		poserErr,
	}, {
		we.Wrap(errors.New("err")),
		&timeout,
		false,
		nil,
	}, {
		we.Wrap(errF).Wrap(&errT),
		&timeout,
		true,
		errF,
	}, {
		we.Wrap(errF).Wrap(&errT),
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
			{
				match := tc.err.As(tc.target)
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

type poser struct {
	msg string
	f   func(error) bool
}

var poserPathErr = &os.PathError{Op: "poser"}

func (p *poser) Error() string     { return p.msg }
func (p *poser) Is(err error) bool { return p.f(err) }
func (p *poser) As(err interface{}) bool {
	switch x := err.(type) {
	case **poser:
		*x = p
	case *errorT:
		*x = errorT{"poser"}
	case **os.PathError:
		*x = poserPathErr
	default:
		return false
	}
	return true
}

func TestAsValidation(t *testing.T) {
	var s string
	var we *wrapError
	testCases := []interface{}{
		nil,
		(*int)(nil),
		"error",
		&s,
	}
	err := we.Wrap(errors.New("error"))
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%T(%v)", tc, tc), func(t *testing.T) {
			defer func() {
				recover()
			}()
			{
				if err.As(tc) {
					t.Errorf("As(err, %T(%v)) = true, want false", tc, tc)
					return
				}
				t.Errorf("As(err, %T(%v)) did not panic", tc, tc)
			}

			{
				if errors.As(err, tc) {
					t.Errorf("As(err, %T(%v)) = true, want false", tc, tc)
					return
				}
				t.Errorf("As(err, %T(%v)) did not panic", tc, tc)
			}
		})
	}
}

type errorT struct{ s string }

func (e errorT) Error() string { return fmt.Sprintf("errorT(%s)", e.s) }

type errorUncomparable struct {
	f []string
}

func (errorUncomparable) Error() string {
	return "uncomparable error"
}

func (errorUncomparable) Is(target error) bool {
	_, ok := target.(errorUncomparable)
	return ok
}

func TestUnwrap(t *testing.T) {
	err1 := errors.New("1")
	err2 := errors.New("2")
	erra := &wrapError{new: err1, msg: "a1"}
	errb := &wrapError{new: err2, msg: "b2"}

	testCases := []struct {
		err  *wrapError
		want error
	}{
		{nil, nil},
		{erra, err1},
		{errb, err2},
		{&wrapError{old: errb, new: err1}, err1},
		{&wrapError{old: errb, new: nil}, nil},
		{&wrapError{old: &wrapError{old: erra, new: err2}, new: err1}, err1},
		{&wrapError{old: &wrapError{old: erra, new: err2}, new: erra}, erra},
	}
	for _, tc := range testCases {
		if got := errors.Unwrap(tc.err); got != tc.want {
			t.Errorf("Unwrap(%v) = %v, want %v", tc.err, got, tc.want)
		}

		if got := tc.err.Unwrap(); got != tc.want {
			t.Errorf("Unwrap(%v) = %v, want %v", tc.err, got, tc.want)
		}
	}
}
