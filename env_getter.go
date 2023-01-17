package mengpo

import (
	"github.com/goexl/env"
)

type envGetter struct{}

func newEnvGetter() *envGetter {
	return new(envGetter)
}

func (eg *envGetter) Get(key string) string {
	return env.Get(key)
}
