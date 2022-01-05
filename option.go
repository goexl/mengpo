package mengpo

type (
	option interface {
		apply(options *options)
	}

	options struct {
		tag    string
		before beforeFunc
		after  afterFunc
	}
)

func defaultOptions() *options {
	return &options{
		tag:    `default`,
		before: beforeDefault,
		after:  afterDefault,
	}
}
