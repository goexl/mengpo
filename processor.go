package mengpo

import (
	"reflect"
)

type processor interface {
	// Process 默认值处理
	Process(tag string, field reflect.StructField) (to string, err error)
}
