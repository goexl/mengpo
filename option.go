package mengpo

import (
	"reflect"

	"github.com/goexl/env"
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		tag        string
		initialize bool
		errorMod   errorMod
		envGetter  envGetter
		processors []processor
	}
)

func defaultOptions() (_options *options) {
	_options = new(options)
	_options.tag = `default`
	_options.initialize = true
	_options.errorMod = ErrorModReturn
	_options.envGetter = env.Get

	pd := new(processorDefault)
	pd.options = _options
	_options.processors = []processor{
		pd,
	}

	return
}

func (o *options) doProcessors(from string, field reflect.StructField) (to string, err error) {
	for _, _processor := range o.processors {
		if to, err = _processor.Process(from, field); nil != err {
			break
		} else {
			from = to
		}
	}

	return
}
