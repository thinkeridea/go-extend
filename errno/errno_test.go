package errno

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	t1 := New(0, "ok")
	t2 := New(0, "ok")
	t3 := New(-1, "message")

	if t1 == t2 {
		t.Errorf("multiple New returns the same instance %p == %p", t1, t2)
	}

	if e, ok := t1.(*errno); !ok {
		t.Error("t1  is not an instance of *errno")
	} else if e.code != 0 {
		t.Errorf("t1.code %d != 0", e.code)
	} else if e.message != "ok" {
		t.Errorf("t1.message %s != ok", e.message)
	}

	if e, ok := t2.(*errno); !ok {
		t.Error("t2  is not an instance of *errno")
	} else if e.code != 0 {
		t.Errorf("t2.code %d != 0", e.code)
	} else if e.message != "ok" {
		t.Errorf("t2.message %s != ok", e.message)
	}

	if e, ok := t3.(*errno); !ok {
		t.Error("t3  is not an instance of *errno")
	} else if e.code != -1 {
		t.Errorf("t3.code %d != -1", e.code)
	} else if e.message != "message" {
		t.Errorf("t3.message %s != message", e.message)
	}
}

func TestNewCode(t *testing.T) {
	t1 := NewCode(0)
	t2 := NewCode(0)
	t3 := NewCode(-1)

	if t1 == t2 {
		t.Errorf("multiple NewCode returns the same instance %p == %p", t1, t2)
	}

	if e, ok := t1.(*errno); !ok {
		t.Error("t1  is not an instance of *errno")
	} else if e.code != 0 {
		t.Errorf("t1.code %d != 0", e.code)
	} else if e.message != "" {
		t.Errorf("t1.message %s not empty", e.message)
	}

	if e, ok := t2.(*errno); !ok {
		t.Error("t2  is not an instance of *errno")
	} else if e.code != 0 {
		t.Errorf("t2.code %d != 0", e.code)
	} else if e.message != "" {
		t.Errorf("t2.message %s not empty ", e.message)
	}

	if e, ok := t3.(*errno); !ok {
		t.Error("t3  is not an instance of *errno")
	} else if e.code != -1 {
		t.Errorf("t3.code %d != -1", e.code)
	} else if e.message != "" {
		t.Errorf("t3.message %s not empty", e.message)
	}
}

func TestErrno_Code(t *testing.T) {
	t1 := NewCode(1)
	t2 := NewCode(2)
	t3 := NewCode(-1)

	if t1.Code() != 1 {
		t.Errorf("t1.Code() %d != 1", t1.Code())
	}

	if t2.Code() != 2 {
		t.Errorf("t2.Code() %d != 2", t2.Code())
	}

	if t3.Code() != -1 {
		t.Errorf("t3.Code() %d != -1", t3.Code())
	}
}

func TestErrno_Message(t *testing.T) {
	t1 := New(1, "Error Message 1")
	t2 := New(2, "Error Message 2")
	t3 := New(-1, "System Error")

	if t1.Message() != "Error Message 1" {
		t.Errorf("t1.Message() %s != 'Error Message 1'", t1.Message())
	}

	if t2.Message() != "Error Message 2" {
		t.Errorf("t2.Message() %s != 'Error Message 2'", t2.Message())
	}

	if t3.Message() != "System Error" {
		t.Errorf("t3.Message() %s != 'System Error'", t3.Message())
	}
}

func TestErrno_Error(t *testing.T) {
	err := New(-1, "System Error")
	if err.Error() != "Error - code: -1, message: System Error" {
		t.Errorf("err.Error():'%s' != 'Error - code: -1, message: System Error'", err.Error())
	}

	if fmt.Sprintf("%v", err) != "Error - code: -1, message: System Error" {
		t.Errorf("fmt.Sprintf(\"%%v\", err):'%s' != 'Error - code: -1, message: System Error'", fmt.Sprintf("%v", err))
	}
}
