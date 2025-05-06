package getter

import (
	"github.com/goexl/env"
)

type Environment struct{}

func NewEnvironment() *Environment {
	return new(Environment)
}

func (e *Environment) Get(key string) string {
	return env.Get(key)
}
