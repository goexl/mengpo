package mengpo

var _ option = (*optionAfter)(nil)

type optionAfter struct {
	after afterFunc
}

// After 配置生命周期后的操作
func After(after afterFunc) *optionAfter {
	return &optionAfter{
		after: after,
	}
}

func (a *optionAfter) apply(options *options) {
	options.after = a.after
}
