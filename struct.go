package mengpo

import (
	"reflect"
	"strings"
)

func (m *mengpo) structure(field reflect.Value, tag string) (err error) {
	if "" == strings.TrimSpace(tag) || jsonStruct == tag {
		return
	}
	err = m.convertJson(tag, field.Addr().Interface())

	return
}
