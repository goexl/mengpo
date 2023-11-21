package mengpo

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
)

func (m *mengpo) Settable(ptr any) bool {
	return m.settable(reflect.ValueOf(ptr))
}

func (m *mengpo) Set(ptr any) (err error) {
	kind := reflect.TypeOf(ptr).Kind()
	value := reflect.ValueOf(ptr).Elem()
	_type := value.Type()
	if reflect.Ptr != kind {
		err = exception.New().Message(errorInvalidType).Field(field.New("kind", kind.String())).Build()
	} else if reflect.Struct != _type.Kind() {
		err = exception.New().Message(errorInvalidType).Field(field.New("type", _type.String())).Build()
	}
	if nil != err {
		return
	}

	// 截获错误并按用户配置处理
	defer func() {
		if nil == err {
			return
		}

		switch m.options.errorMod {
		case errorModSilent:
			err = nil
		case errorModPanic:
			panic(err)
		}
	}()

	for index := 0; index < _type.NumField(); index++ {
		_field := _type.Field(index)
		tag := _field.Tag.Get(m.options.tag)
		if tagIgnore == tag {
			continue
		}

		if tag, err = m.options.process(tag, _field); nil != err {
			return
		}
		if err = m.setField(value.Field(index), tag, m.options); nil != err {
			return
		}
	}
	m.setter(ptr)

	return
}

func (m *mengpo) setField(field reflect.Value, tag string, options *options) (err error) {
	if !m.canSet(field, tag, options) {
		return
	}

	_settable := m.settable(field)
	if _settable {
		switch field.Kind() {
		case reflect.Bool:
			err = m.bool(field, tag)
		case reflect.Int:
			err = m.int(field, tag)
		case reflect.Int8:
			err = m.int8(field, tag)
		case reflect.Int16:
			err = m.int16(field, tag)
		case reflect.Int32:
			err = m.int32(field, tag)
		case reflect.Int64:
			err = m.int64(field, tag)
		case reflect.Uint:
			err = m.uint(field, tag)
		case reflect.Uint8:
			err = m.uint8(field, tag)
		case reflect.Uint16:
			err = m.uint16(field, tag)
		case reflect.Uint32:
			err = m.uint32(field, tag)
		case reflect.Uint64:
			err = m.uint64(field, tag)
		case reflect.Uintptr:
			err = m.uintPtr(field, tag)
		case reflect.Float32:
			err = m.float32(field, tag)
		case reflect.Float64:
			err = m.float64(field, tag)
		case reflect.String:
			field.Set(reflect.ValueOf(tag).Convert(field.Type()))
		case reflect.Slice:
			err = m.slice(field, tag)
		case reflect.Map:
			err = m.mapping(field, tag)
		case reflect.Struct:
			err = m.structure(field, tag)
		case reflect.Ptr:
			field.Set(reflect.New(field.Type().Elem()))
		}
	}
	if nil != err {
		return
	}

	switch field.Kind() {
	case reflect.Ptr:
		if _settable || field.Elem().Kind() == reflect.Struct {
			// 不关注错误，后面的代码必须执行
			_ = m.setField(field.Elem(), tag, options)
			m.setter(field.Interface())
		}
	case reflect.Struct:
		if err = m.Set(field.Addr().Interface()); nil != err {
			return
		}
	case reflect.Slice:
		for index := 0; index < field.Len(); index++ {
			if err = m.setField(field.Index(index), tag, options); nil != err {
				return
			}
		}
	}

	return
}

func (m *mengpo) convertJson(from string, value any) (err error) {
	// 将JSON字符串转换成易写的形式
	data := strings.ReplaceAll(from, "'", `"`)
	err = json.Unmarshal([]byte(data), value)

	return
}

func (m *mengpo) settable(field reflect.Value) bool {
	return reflect.DeepEqual(reflect.Zero(field.Type()).Interface(), field.Interface())
}

func (m *mengpo) canSet(field reflect.Value, tag string, options *options) (settable bool) {
	if !field.CanSet() {
		return
	}

	switch field.Kind() {
	case reflect.Struct:
		settable = true
	case reflect.Ptr:
		if "" != tag {
			settable = true
		} else {
			ek := field.Elem().Kind()
			settable = reflect.Struct == ek || reflect.Slice == ek || reflect.Map == ek || reflect.Ptr == ek
		}
	case reflect.Slice:
		settable = "" != tag || field.Len() > 0
	case reflect.Map:
		settable = true
	default:
		settable = "" != tag
	}
	settable = options.initialize && settable

	return
}
