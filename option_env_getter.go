package mengpo

var (
	_        = EnvGetter
	_ option = (*optionEnvGetter)(nil)
)

type optionEnvGetter struct {
	getter envGetter
}

// EnvGetter 环境变量获取器
func EnvGetter(getter envGetter) *optionEnvGetter {
	return &optionEnvGetter{
		getter: getter,
	}
}

func (eg *optionEnvGetter) apply(options *options) {
	options.envGetter = eg.getter
}
