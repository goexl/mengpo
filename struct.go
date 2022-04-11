package mengpo

import (
	"reflect"
	"strings"
)

func _struct(field reflect.Value, tag string) (err error) {
	if `` == strings.TrimSpace(tag) || jsonStruct == tag {
		return
	}
	err = convertJson(tag, field.Addr().Interface())

	return
}
