package mengpo_test

import (
	"testing"

	"github.com/goexl/mengpo"
)

type ptr struct {
	True  *bool `default:"true"`
	False *bool `default:"false"`
	Nil   *bool
}

func TestSetByPtr(t *testing.T) {
	_ptr := new(ptr)
	if err := mengpo.Set(_ptr); nil != err {
		t.Fatal(err)
	}

	if true != *_ptr.True {
		t.Fatalf(`期望：true，实际：%v`, *_ptr.True)
	}

	if false != *_ptr.False {
		t.Fatalf(`期望：false，实际：%v`, *_ptr.False)
	}

	if nil != _ptr.Nil {
		t.Fatalf(`期望：nil，实际：%v`, _ptr.Nil)
	}
}
