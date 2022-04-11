package mengpo

import (
	"reflect"
	"strconv"
	"time"
)

func _int(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseInt(tag, 0, strconv.IntSize); nil == pie {
		field.Set(reflect.ValueOf(int(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func _int8(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseInt(tag, 0, 8); nil == pie {
		field.Set(reflect.ValueOf(int8(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func _int16(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseInt(tag, 0, 16); nil == pie {
		field.Set(reflect.ValueOf(int16(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func _int32(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseInt(tag, 0, 32); nil == pie {
		field.Set(reflect.ValueOf(int32(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func _int64(field reflect.Value, tag string) (err error) {
	var value interface{}
	switch field.Interface().(type) {
	case time.Duration:
		value, err = time.ParseDuration(tag)
	default:
		value, err = strconv.ParseInt(tag, 0, 64)
	}
	if nil != err {
		return
	}

	field.Set(reflect.ValueOf(value).Convert(field.Type()))

	return
}
