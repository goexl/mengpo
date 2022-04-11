package mengpo

import (
	"reflect"
	"strings"
)

func _map(field reflect.Value, tag string) (err error) {
	ref := reflect.New(field.Type())
	ref.Elem().Set(reflect.MakeMap(field.Type()))
	if `` == strings.TrimSpace(tag) || jsonMap == tag {
		return
	}

	if err = convertJson(tag, ref.Interface()); nil == err {
		field.Set(ref.Elem().Convert(field.Type()))
	}

	return
}
