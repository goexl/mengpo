package core

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/mengpo/internal/internal/constant"
	"github.com/goexl/mengpo/internal/internal/param"
	"github.com/goexl/mengpo/internal/internal/runtime"
	"github.com/goexl/mengpo/internal/kernel"
)

type Mengpo struct {
	params      *param.Mengpo
	unmarshaler reflect.Type
}

func MewMengpo(params *param.Mengpo) *Mengpo {
	return &Mengpo{
		params:      params,
		unmarshaler: reflect.TypeOf((*kernel.Unmarshaler)(nil)).Elem(),
	}
}

func (m *Mengpo) Set(target runtime.Pointer) (err error) {
	kind := reflect.TypeOf(target).Kind()
	value := reflect.ValueOf(target).Elem()
	typ := value.Type()
	if reflect.Ptr != kind {
		err = exception.New().Message(constant.ErrorInvalidType).Field(field.New("kind", kind.String())).Build()
	} else if reflect.Struct != typ.Kind() {
		err = exception.New().Message(constant.ErrorInvalidType).Field(field.New("type", typ.String())).Build()
	}
	if nil != err {
		return
	}

	// 截获错误并按用户配置处理
	defer func() {
		if nil == err {
			return
		}

		switch m.params.Mode {
		case constant.ModSilent:
			err = nil
		case constant.ModPanic:
			panic(err)
		default:
			panic(err)
		}
	}()

	for index := 0; index < typ.NumField(); index++ {
		setField := typ.Field(index)
		tag := setField.Tag.Get(m.params.Tag)
		if constant.TagIgnore == tag {
			continue
		}

		if tag, err = m.params.Process(tag, setField); nil != err {
			return
		}
		if err = m.setField(value.Field(index), tag); nil != err {
			return
		}
	}

	return
}

func (m *Mengpo) setField(field reflect.Value, tag string) (err error) {
	if !m.canSet(field, tag) {
		// 不做任何操作
	} else if reflect.PointerTo(field.Type()).Implements(m.unmarshaler) || field.Type().Implements(m.unmarshaler) {
		// 实现了反序列化接口
		m.setUnmarshaler(field, tag)
	} else if reflect.DeepEqual(reflect.Zero(field.Type()).Interface(), field.Interface()) { // 判断是否可以被设置值
		err = m.setSettable(field, tag)
	} else {
		err = m.setNotSettable(field, tag)
	}

	return
}

func (m *Mengpo) setUnmarshaler(field reflect.Value, tag string) {
	value := field
	kind := field.Kind()
	if reflect.Ptr == kind && field.IsNil() {
		field.Set(reflect.New(field.Type().Elem())) // 初始化指针字段
	} else {
		value = reflect.New(field.Type())
	}
	method := value.MethodByName("Unmarshal")
	if method.IsValid() { // 调用设置值
		method.Call([]reflect.Value{reflect.ValueOf([]byte(tag))})
	}
	if reflect.Ptr != kind { // 将指针实例的值赋回原字段
		field.Set(value.Elem())
	}

	return
}

func (m *Mengpo) setSettable(field reflect.Value, tag string) (err error) {
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
	default:
		err = nil
	}

	return
}

func (m *Mengpo) setNotSettable(field reflect.Value, tag string) (err error) {
	switch field.Kind() {
	case reflect.Ptr:
		if field.Elem().Kind() == reflect.Struct {
			// 不关注错误，后面的代码必须执行
			_ = m.setField(field.Elem(), tag)
			_ = m.setter(field.Interface())
		}
	case reflect.Struct:
		if err = m.Set(field.Addr().Interface()); nil != err {
			return
		}
	case reflect.Slice:
		for index := 0; index < field.Len(); index++ {
			if err = m.setField(field.Index(index), tag); nil != err {
				return
			}
		}
	default:
		err = nil
	}

	return
}

func (m *Mengpo) convertJson(from string, value any) (err error) {
	// 将JSON字符串转换成易写的形式
	data := strings.ReplaceAll(from, "'", `"`)
	err = json.Unmarshal([]byte(data), value)

	return
}

func (m *Mengpo) canSet(field reflect.Value, tag string) (settable bool) {
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
	settable = m.params.Initialize && settable

	return
}

func (m *Mengpo) bool(field reflect.Value, tag string) (err error) {
	if value, pbe := strconv.ParseBool(tag); nil == pbe {
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	} else {
		err = pbe
	}

	return
}

func (m *Mengpo) float32(field reflect.Value, tag string) (err error) {
	if value, pfe := strconv.ParseFloat(tag, 32); nil == pfe {
		field.Set(reflect.ValueOf(float32(value)).Convert(field.Type()))
	} else {
		err = pfe
	}

	return
}

func (m *Mengpo) float64(field reflect.Value, tag string) (err error) {
	if value, pfe := strconv.ParseFloat(tag, 64); nil == pfe {
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	} else {
		err = pfe
	}

	return
}

func (m *Mengpo) int(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseInt(tag, 0, strconv.IntSize); nil == pie {
		field.Set(reflect.ValueOf(int(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *Mengpo) int8(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseInt(tag, 0, 8); nil == pie {
		field.Set(reflect.ValueOf(int8(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *Mengpo) int16(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseInt(tag, 0, 16); nil == pie {
		field.Set(reflect.ValueOf(int16(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *Mengpo) int32(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseInt(tag, 0, 32); nil == pie {
		field.Set(reflect.ValueOf(int32(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *Mengpo) int64(field reflect.Value, tag string) (err error) {
	var value any
	switch field.Interface().(type) {
	case time.Duration:
		value, err = time.ParseDuration(tag)
	case gox.Bytes:
		value, err = gox.ParseBytes(tag)
	default:
		value, err = strconv.ParseInt(tag, 0, 64)
	}
	if nil == err {
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	}

	return
}

func (m *Mengpo) slice(field reflect.Value, tag string) (err error) {
	ref := reflect.New(field.Type())
	ref.Elem().Set(reflect.MakeSlice(field.Type(), 0, 0))
	if `` == strings.TrimSpace(tag) || constant.JsonSlice == tag {
		return
	}

	if jsonErr := m.convertJson(tag, ref.Interface()); nil == jsonErr {
		field.Set(ref.Elem().Convert(field.Type()))
	} else {
		err = jsonErr
	}

	return
}

func (m *Mengpo) structure(field reflect.Value, tag string) (err error) {
	if "" == strings.TrimSpace(tag) || constant.JsonStruct == tag {
		return
	}
	err = m.convertJson(tag, field.Addr().Interface())

	return
}

func (m *Mengpo) uint(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, strconv.IntSize); nil == pie {
		field.Set(reflect.ValueOf(uint(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *Mengpo) uint8(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, 8); nil == pie {
		field.Set(reflect.ValueOf(uint8(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *Mengpo) uint16(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, 16); nil == pie {
		field.Set(reflect.ValueOf(uint16(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *Mengpo) uint32(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, 32); nil == pie {
		field.Set(reflect.ValueOf(uint32(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *Mengpo) uint64(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, 64); nil == pie {
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *Mengpo) uintPtr(field reflect.Value, tag string) (err error) {
	if value, pie := strconv.ParseUint(tag, 0, strconv.IntSize); nil == pie {
		field.Set(reflect.ValueOf(uintptr(value)).Convert(field.Type()))
	} else {
		err = pie
	}

	return
}

func (m *Mengpo) mapping(field reflect.Value, tag string) (err error) {
	ref := reflect.New(field.Type())
	ref.Elem().Set(reflect.MakeMap(field.Type()))
	if "" == strings.TrimSpace(tag) || constant.JsonMap == tag {
		return
	}

	if err = m.convertJson(tag, ref.Interface()); nil == err {
		field.Set(ref.Elem().Convert(field.Type()))
	}

	return
}

func (m *Mengpo) setter(val any) (err error) {
	if _, ok := val.(kernel.Unmarshaler); ok {
		// s.Default()
	}

	return
}
