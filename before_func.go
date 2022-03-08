package mengpo

import (
	`reflect`
)

type beforeFunc func(from string, field reflect.StructField) (to string, err error)
