package param

import (
	"reflect"

	"github.com/goexl/mengpo/internal/internal/constant"
	"github.com/goexl/mengpo/internal/internal/getter"
	"github.com/goexl/mengpo/internal/internal/processor"
	"github.com/goexl/mengpo/internal/kernel"
)

type Mengpo struct {
	Tag        string
	Initialize bool
	Mode       constant.Mode
	Getters    []kernel.Getter
	Processors []kernel.Processor
}

func NewMengpo() (mengpo *Mengpo) {
	mengpo = new(Mengpo)
	mengpo.Tag = "default"
	mengpo.Initialize = true
	mengpo.Mode = constant.ModReturn
	mengpo.Getters = []kernel.Getter{
		getter.NewEnvironment(),
	}
	mengpo.Processors = []kernel.Processor{
		processor.NewDefault(&mengpo.Getters),
	}

	return
}

func (m *Mengpo) Process(from string, field reflect.StructField) (to string, err error) {
	for _, _processor := range m.Processors {
		if to, err = _processor.Process(from, field); nil != err {
			break
		} else {
			from = to
		}
	}

	return
}
