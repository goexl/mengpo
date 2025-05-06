package processor

import (
	"reflect"
	"strings"

	"github.com/drone/envsubst"
	"github.com/goexl/mengpo/internal/kernel"
)

type Default struct {
	getters *[]kernel.Getter
}

func NewDefault(getters *[]kernel.Getter) *Default {
	return &Default{
		getters: getters,
	}
}

func (d *Default) Process(tag string, _ reflect.StructField) (to string, err error) {
	count := 0
	prefix := `$`

	to = tag
	for {
		if to, err = envsubst.Eval(to, d.get); nil != err {
			break
		}

		if count >= 2 || !strings.Contains(to, prefix) {
			break
		}
		if strings.Contains(to, prefix) {
			count++
		}
	}

	return
}

func (d *Default) get(from string) (to string) {
	for _, getter := range *d.getters {
		to = getter.Get(from)
		if "" != to {
			break
		}
	}

	return
}
