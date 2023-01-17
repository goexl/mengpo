package mengpo

import (
	"reflect"
)

type options struct {
	tag        string
	initialize bool
	errorMod   errorMode
	getter     getter
	processors []processor
}

func newDefaultOptions() (_options *options) {
	_options = new(options)
	_options.tag = "default"
	_options.initialize = true
	_options.errorMod = errorModReturn
	_options.getter = newEnvGetter()
	_options.processors = []processor{
		newDefaultProcessor(_options),
	}

	return
}

func (o *options) process(from string, field reflect.StructField) (to string, err error) {
	for _, _processor := range o.processors {
		if to, err = _processor.Process(from, field); nil != err {
			break
		} else {
			from = to
		}
	}

	return
}
