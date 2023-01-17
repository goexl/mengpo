package mengpo

import (
	"reflect"
	"strconv"
)

func (m *mengpo) uint(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, strconv.IntSize); nil == pie {
		field.Set(reflect.ValueOf(uint(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *mengpo) uint8(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, 8); nil == pie {
		field.Set(reflect.ValueOf(uint8(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *mengpo) uint16(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, 16); nil == pie {
		field.Set(reflect.ValueOf(uint16(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *mengpo) uint32(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, 32); nil == pie {
		field.Set(reflect.ValueOf(uint32(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *mengpo) uint64(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, 64); nil == pie {
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *mengpo) uintPtr(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, strconv.IntSize); nil == pie {
		field.Set(reflect.ValueOf(uintptr(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}
