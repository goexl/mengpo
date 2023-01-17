package mengpo

import (
	"reflect"
	"strconv"
)

func (m *mengpo) float32(field reflect.Value, tag string) (err error) {
	if value, pfe := strconv.ParseFloat(tag, 32); nil == pfe {
		field.Set(reflect.ValueOf(float32(value)).Convert(field.Type()))
	} else {
		err = pfe
	}

	return
}

func (m *mengpo) float64(field reflect.Value, tag string) (err error) {
	if value, pfe := strconv.ParseFloat(tag, 64); nil == pfe {
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	} else {
		err = pfe
	}

	return
}
