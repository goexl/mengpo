package mengpo

var (
	_        = ClearBefore
	_ option = (*optionClearBefore)(nil)
)

type optionClearBefore struct{}

// ClearBefore 清除
func ClearBefore() *optionClearBefore {
	return new(optionClearBefore)
}

func (c *optionClearBefore) apply(options *options) {
	options.before = make([]beforeFunc, 0)
}
