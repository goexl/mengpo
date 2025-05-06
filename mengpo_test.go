package mengpo_test

import (
	"testing"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/mengpo"
)

type (
	user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	ptr struct {
		True    bool  `default:"true"`
		False   *bool `default:"false"`
		Nil     *bool
		Timeout time.Duration `default:"1h"`
		Bytes   gox.Bytes     `default:"1g"`
		User    user          `default:"{'username': 'storezhang', 'password': 'test'}"`
		Users   []*user       `default:"[{'username': 'storezhang', 'password': 'test'}]"`
	}

	envPtr struct {
		Env string `default:"${TEST_ENV}"`
	}
)

func TestSetByPtr(t *testing.T) {
	_ptr := new(ptr)
	if err := mengpo.New().Build().Set(_ptr); nil != err {
		t.Fatal(err)
	}

	if true != _ptr.True {
		t.Fatalf(`期望：true，实际：%v`, _ptr.True)
	}

	if false != *_ptr.False {
		t.Fatalf(`期望：false，实际：%v`, *_ptr.False)
	}

	if nil != _ptr.Nil {
		t.Fatalf(`期望：nil，实际：%v`, _ptr.Nil)
	}

	if time.Hour != _ptr.Timeout {
		t.Fatalf(`期望：%d，实际：%v`, time.Hour, _ptr.Timeout)
	}

	if gox.BytesGB != _ptr.Bytes {
		t.Fatalf(`期望：%d，实际：%v`, gox.BytesGB, _ptr.Bytes)
	}

	if `storezhang` != _ptr.User.Username || `test` != _ptr.User.Password {
		t.Fatalf(`期望：{"username": "storezhang", "password": "test"}，实际：%v`, _ptr.User)
	}
	if 1 != len(_ptr.Users) {
		t.Fatalf(`期望：[{"username": "storezhang", "password": "test"}]，实际：%v`, _ptr.Users)
	}
}

type getter struct{}

func (g *getter) Get(key string) string {
	return key
}

func TestEnvGetter(t *testing.T) {
	_ptr := new(envPtr)
	if err := mengpo.New().Getter(new(getter)).Build().Set(_ptr); nil != err {
		t.Fatal(err)
	}
	if "TEST_ENV" != _ptr.Env {
		t.Fatalf("期望：TEST_ENV，实际：%v", _ptr.Env)
	}
}

type custom uint8

func (c *custom) Unmarshal(_ []byte) (err error) {
	*c = custom(1)

	return
}

type customType struct {
	Normal  custom  `default:"5"`
	Pointer *custom `default:"6"`
}

func TestCustom(t *testing.T) {
	_custom := new(customType)
	if err := mengpo.New().Build().Set(_custom); nil != err {
		t.Fatal(err)
	}
	if custom(1) != _custom.Normal {
		t.Fatalf("期望：1，实际：%v", _custom.Normal)
	}
	if custom(1) != *_custom.Pointer {
		t.Fatalf("期望：1，实际：%v", _custom.Normal)
	}
}
