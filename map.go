package mengpo

import (
	"reflect"
	"strings"
)

func (m *mengpo) mapping(field reflect.Value, tag string) (err error) {
	ref := reflect.New(field.Type())
	ref.Elem().Set(reflect.MakeMap(field.Type()))
	if "" == strings.TrimSpace(tag) || jsonMap == tag {
		return
	}

	if err = m.convertJson(tag, ref.Interface()); nil == err {
		field.Set(ref.Elem().Convert(field.Type()))
	}

	return
}
