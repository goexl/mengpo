package kernel

import (
	"reflect"
)

type Processor interface {
	Process(tag string, field reflect.StructField) (to string, err error)
}
