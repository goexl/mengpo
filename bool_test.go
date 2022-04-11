package mengpo_test

import (
	"testing"
	"time"

	"github.com/goexl/mengpo"
)

type boolPtr struct {
	Test  time.Duration `default:"1h"`
	True  *bool         `default:"true"`
	False *bool         `default:"false"`
	Int8  int32         `default:"11"`
	Nil   *bool
}

func TestBoolPtr(t *testing.T) {
	ptr := new(boolPtr)
	if err := mengpo.Set(ptr); nil != err {
		t.Fatal(err)
	}

	if true != *ptr.True {
		t.Fatalf(`期望：true，实际：%v`, *ptr.True)
	}

	if false != *ptr.False {
		t.Fatalf(`期望：false，实际：%v`, *ptr.False)
	}

	if nil != ptr.Nil {
		t.Fatalf(`期望：nil，实际：%v`, ptr.Nil)
	}
}
