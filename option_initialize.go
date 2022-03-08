package mengpo

var (
	_        = Initialize
	_        = DisableInitialize
	_ option = (*optionInitialize)(nil)
)

type optionInitialize struct {
	initialize bool
}

// Initialize 需要初始化
func Initialize() *optionInitialize {
	return &optionInitialize{
		initialize: true,
	}
}

// DisableInitialize 不初始化
func DisableInitialize() *optionInitialize {
	return &optionInitialize{
		initialize: false,
	}
}

func (i *optionInitialize) apply(options *options) {
	options.initialize = i.initialize
}
