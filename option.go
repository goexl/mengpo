package mengpo

type (
	option interface {
		apply(options *options)
	}

	options struct {
		tag    string
		before []beforeFunc
	}
)

func defaultOptions() *options {
	return &options{
		tag: `default`,
		before: []beforeFunc{
			beforeDefault,
		},
	}
}

func (o *options) doBefore(from string) (to string, err error) {
	for _, before := range o.before {
		if to, err = before(from); nil != err {
			break
		} else {
			from = to
		}
	}

	return
}
