package mengpo

type (
	option interface {
		apply(options *options)
	}

	options struct {
		tag    string
		before beforeFunc
	}
)

func defaultOptions() *options {
	return &options{
		tag:    `default`,
		before: beforeDefault,
	}
}
