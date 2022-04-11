package mengpo

import (
	"reflect"
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		tag        string
		initialize bool
		silence    bool
		before     []beforeFunc
	}
)

func defaultOptions() *options {
	return &options{
		tag:        `default`,
		initialize: true,
		silence:    false,
		before: []beforeFunc{
			beforeDefault,
		},
	}
}

func (o *options) doBefore(from string, field reflect.StructField) (to string, err error) {
	for _, before := range o.before {
		if to, err = before(from, field); nil != err {
			break
		} else {
			from = to
		}
	}

	return
}
