package god

type setter interface {
	Defaults()
}

func _setter(val interface{}) {
	if s, ok := val.(setter); ok {
		s.Defaults()
	}
}
