package mengpo

type mengpo struct {
	options *options
}

func newMengpo(options *options) *mengpo {
	return &mengpo{
		options: options,
	}
}
