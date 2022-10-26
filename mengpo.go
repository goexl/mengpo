package mengpo

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

var (
	_ = Set
	_ = Settable
)

// Settable 是否可被设置默认值
func Settable(ptr interface{}) bool {
	return settable(reflect.ValueOf(ptr))
}

// Set 设置默认值
func Set(ptr interface{}, opts ...option) (err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	kind := reflect.TypeOf(ptr).Kind()
	value := reflect.ValueOf(ptr).Elem()
	_type := value.Type()
	if reflect.Ptr != kind {
		err = exc.NewField(errorInvalidType, field.String(`kind`, kind.String()))
	} else if reflect.Struct != _type.Kind() {
		err = exc.NewField(errorInvalidType, field.String(`type`, _type.String()))
	}
	if nil != err {
		return
	}

	// 截获错误并按用户配置处理
	defer func() {
		if nil == err {
			return
		}

		switch _options.errorMod {
		case ErrorModSilent:
			err = nil
		case ErrorModPanic:
			panic(err)
		}
	}()

	for index := 0; index < _type.NumField(); index++ {
		_field := _type.Field(index)
		tag := _field.Tag.Get(_options.tag)
		if tagIgnore == tag {
			continue
		}

		if tag, err = _options.doProcessors(tag, _field); nil != err {
			return
		}
		if err = setField(value.Field(index), tag, _options); nil != err {
			return
		}
	}
	_setter(ptr)

	return
}

func setField(field reflect.Value, tag string, options *options) (err error) {
	if !canSet(field, tag, options) {
		return
	}

	_settable := settable(field)
	if _settable {
		switch field.Kind() {
		case reflect.Bool:
			err = _bool(field, tag)
		case reflect.Int:
			err = _int(field, tag)
		case reflect.Int8:
			err = _int8(field, tag)
		case reflect.Int16:
			err = _int16(field, tag)
		case reflect.Int32:
			err = _int32(field, tag)
		case reflect.Int64:
			err = _int64(field, tag)
		case reflect.Uint:
			err = _uint(field, tag)
		case reflect.Uint8:
			err = _uint8(field, tag)
		case reflect.Uint16:
			err = _uint16(field, tag)
		case reflect.Uint32:
			err = _uint32(field, tag)
		case reflect.Uint64:
			err = _uint64(field, tag)
		case reflect.Uintptr:
			err = _uintPtr(field, tag)
		case reflect.Float32:
			err = _float32(field, tag)
		case reflect.Float64:
			err = _float64(field, tag)
		case reflect.String:
			field.Set(reflect.ValueOf(tag).Convert(field.Type()))
		case reflect.Slice:
			err = slice(field, tag)
		case reflect.Map:
			err = _map(field, tag)
		case reflect.Struct:
			err = _struct(field, tag)
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
			_ = setField(field.Elem(), tag, options)
			_setter(field.Interface())
		}
	case reflect.Struct:
		if err = Set(field.Addr().Interface()); nil != err {
			return
		}
	case reflect.Slice:
		for index := 0; index < field.Len(); index++ {
			if err = setField(field.Index(index), tag, options); nil != err {
				return
			}
		}
	}

	return
}

func convertJson(from string, value interface{}) (err error) {
	// 将JSON字符串转换成易写的形式
	data := strings.ReplaceAll(from, `'`, `"`)
	err = json.Unmarshal([]byte(data), value)

	return
}

func settable(field reflect.Value) bool {
	return reflect.DeepEqual(reflect.Zero(field.Type()).Interface(), field.Interface())
}

func canSet(field reflect.Value, tag string, options *options) (set bool) {
	if !field.CanSet() {
		return
	}

	switch field.Kind() {
	case reflect.Struct:
		set = true
	case reflect.Ptr:
		if `` != tag {
			set = true
		} else {
			ek := field.Elem().Kind()
			set = reflect.Struct == ek || reflect.Slice == ek || reflect.Map == ek || reflect.Ptr == ek
		}
	case reflect.Slice:
		set = `` != tag || field.Len() > 0
	case reflect.Map:
		set = true
	default:
		set = `` != tag
	}
	set = options.initialize && set

	return
}
