package builder

import (
	"github.com/goexl/mengpo/internal/core"
	"github.com/goexl/mengpo/internal/internal/constant"
	"github.com/goexl/mengpo/internal/internal/param"
	"github.com/goexl/mengpo/internal/kernel"
)

type Mengpo struct {
	params *param.Mengpo
}

func New() *Mengpo {
	return &Mengpo{
		params: param.NewMengpo(),
	}
}

func (m *Mengpo) Tag(tag string) *Mengpo {
	return m.set(func() {
		m.params.Tag = tag
	})
}

func (m *Mengpo) Initialize(initialize bool) *Mengpo {
	return m.set(func() {
		m.params.Initialize = initialize
	})
}

func (m *Mengpo) PanicOnError() *Mengpo {
	return m.set(func() {
		m.params.Mode = constant.ModPanic
	})
}

func (m *Mengpo) ReturnOnError() *Mengpo {
	return m.set(func() {
		m.params.Mode = constant.ModReturn
	})
}

func (m *Mengpo) SilentOnError() *Mengpo {
	return m.set(func() {
		m.params.Mode = constant.ModSilent
	})
}

func (m *Mengpo) Getter(getter kernel.Getter) *Mengpo {
	return m.set(func() {
		m.params.Getters = append(m.params.Getters, getter)
	})
}

func (m *Mengpo) Processor(processor kernel.Processor) *Mengpo {
	return m.set(func() {
		m.params.Processors = append(m.params.Processors, processor)
	})
}

func (m *Mengpo) Build() *core.Mengpo {
	return core.MewMengpo(m.params)
}

func (m *Mengpo) set(callback func()) (mengpo *Mengpo) {
	callback()
	mengpo = m

	return
}
