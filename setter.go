package mengpo

type setter interface {
	// Default 设置默认值
	Default()
}

func (m *mengpo) setter(val any) {
	if s, ok := val.(setter); ok {
		s.Default()
	}
}
