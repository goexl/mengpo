package mengpo

type setter interface {
	// Default 设置默认值
	Default()
}

func _setter(val interface{}) {
	if s, ok := val.(setter); ok {
		s.Default()
	}
}
