package mengpo

import (
	"reflect"
	"strconv"
)

func _bool(field reflect.Value, tag string) (err error) {
	if value, pbe := strconv.ParseBool(tag); nil == pbe {
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	} else {
		err = pbe
	}

	return
}
