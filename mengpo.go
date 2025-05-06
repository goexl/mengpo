package mengpo

import (
	"github.com/goexl/mengpo/internal/builder"
	"github.com/goexl/mengpo/internal/core"
)

type Mengpo = core.Mengpo

func New() *builder.Mengpo {
	return builder.New()
}
