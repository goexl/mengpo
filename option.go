package mengpo

import (
	`os`
)

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
		tag: `default`,
		before: func(original string) string {
			return os.ExpandEnv(original)
		},
	}
}
