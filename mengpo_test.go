package mengpo_test

import (
	"testing"

	"github.com/goexl/mengpo"
)

type (
	user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	ptr struct {
		True  *bool `default:"true"`
		False *bool `default:"false"`
		Nil   *bool
		User  user    `default:"{'username': 'storezhang', 'password': 'test'}"`
		Users []*user `default:"[{'username': 'storezhang', 'password': 'test'}]"`
	}
)

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

	if `storezhang` != _ptr.User.Username || `test` != _ptr.User.Password {
		t.Fatalf(`期望：{"username": "storezhang", "password": "test"}，实际：%v`, _ptr.User)
	}
	if 1 != len(_ptr.Users) {
		t.Fatalf(`期望：[{"username": "storezhang", "password": "test"}]，实际：%v`, _ptr.Users)
	}
}
