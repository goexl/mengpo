package mengpo_test

import (
	`testing`

	`github.com/goexl/mengpo`
)

type boolPtr struct {
	True  *bool `default:"true"`
	False *bool `default:"false"`
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
