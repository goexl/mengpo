package mengpo

import (
	"reflect"
	"strings"
)

func (m *mengpo) slice(field reflect.Value, tag string) (err error) {
	ref := reflect.New(field.Type())
	ref.Elem().Set(reflect.MakeSlice(field.Type(), 0, 0))
	if `` == strings.TrimSpace(tag) || jsonSlice == tag {
		return
	}

	if jsonErr := m.convertJson(tag, ref.Interface()); nil == jsonErr {
		field.Set(ref.Elem().Convert(field.Type()))
	} else {
		err = jsonErr
	}

	return
}
