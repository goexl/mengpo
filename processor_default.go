package mengpo

import (
	"reflect"
	"strings"

	"github.com/drone/envsubst"
)

type processorDefault struct {
	options *options
}

func newDefaultProcessor(_options *options) *processorDefault {
	return &processorDefault{
		options: _options,
	}
}

func (p *processorDefault) Process(tag string, _ reflect.StructField) (to string, err error) {
	hasEnvCount := 0
	to = tag
	envPrefix := `$`
	for {
		if to, err = envsubst.Eval(to, p.options.getter.Get); nil != err {
			break
		}

		if hasEnvCount >= 2 || !strings.Contains(to, envPrefix) {
			break
		}
		if strings.Contains(to, envPrefix) {
			hasEnvCount++
		}
	}

	return
}
