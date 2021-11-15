package mengpo

type (
	option interface {
		apply(options *options)
	}

	options struct {
		tag string
	}
)

func defaultOptions() *options {
	return &options{
		tag: `default`,
	}
}
