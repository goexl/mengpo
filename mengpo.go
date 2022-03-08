package mengpo

import (
	`encoding/json`
	`reflect`
	`strconv`
	`strings`
	`time`
)

var (
	_ = Set
	_ = Settable
)

// Settable 是否可被设置默认值
func Settable(ptr interface{}) bool {
	return isInitialValue(reflect.ValueOf(ptr))
}

// Set 设置默认值
func Set(ptr interface{}, opts ...option) (err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		err = errInvalidType
	}
	if nil != err {
		return
	}

	value := reflect.ValueOf(ptr).Elem()
	_type := value.Type()
	if reflect.Struct != _type.Kind() {
		err = errInvalidType
	}
	if nil != err {
		return
	}

	for index := 0; index < _type.NumField(); index++ {
		sf := _type.Field(index)
		if tag := sf.Tag.Get(_options.tag); tagIgnore != tag {
			if tag, err = _options.doBefore(tag, sf); nil != err {
				return
			}
			if err = setField(value.Field(index), tag); nil != err {
				return
			}
		}
	}
	_setter(ptr)

	return
}

func setField(field reflect.Value, tag string) (err error) {
	if !field.CanSet() {
		return
	}

	if !isInitialField(field, tag) {
		return
	}

	isInitial := isInitialValue(field)
	if isInitial {
		switch field.Kind() {
		case reflect.Bool:
			if value, pbe := strconv.ParseBool(tag); nil == pbe {
				field.Set(reflect.ValueOf(value).Convert(field.Type()))
			}
		case reflect.Int:
			if value, pie := strconv.ParseInt(tag, 0, strconv.IntSize); nil == pie {
				field.Set(reflect.ValueOf(int(value)).Convert(field.Type()))
			}
		case reflect.Int8:
			if value, pie := strconv.ParseInt(tag, 0, 8); nil == pie {
				field.Set(reflect.ValueOf(int8(value)).Convert(field.Type()))
			}
		case reflect.Int16:
			if value, pie := strconv.ParseInt(tag, 0, 16); nil == pie {
				field.Set(reflect.ValueOf(int16(value)).Convert(field.Type()))
			}
		case reflect.Int32:
			if value, pie := strconv.ParseInt(tag, 0, 32); nil == pie {
				field.Set(reflect.ValueOf(int32(value)).Convert(field.Type()))
			}
		case reflect.Int64:
			if value, pde := time.ParseDuration(tag); nil == pde {
				field.Set(reflect.ValueOf(value).Convert(field.Type()))
			} else if intValue, pie := strconv.ParseInt(tag, 0, 64); nil == pie {
				field.Set(reflect.ValueOf(intValue).Convert(field.Type()))
			}
		case reflect.Uint:
			if value, pie := strconv.ParseUint(tag, 0, strconv.IntSize); nil == pie {
				field.Set(reflect.ValueOf(uint(value)).Convert(field.Type()))
			}
		case reflect.Uint8:
			if value, pie := strconv.ParseUint(tag, 0, 8); nil == pie {
				field.Set(reflect.ValueOf(uint8(value)).Convert(field.Type()))
			}
		case reflect.Uint16:
			if value, pie := strconv.ParseUint(tag, 0, 16); nil == pie {
				field.Set(reflect.ValueOf(uint16(value)).Convert(field.Type()))
			}
		case reflect.Uint32:
			if value, pie := strconv.ParseUint(tag, 0, 32); nil == pie {
				field.Set(reflect.ValueOf(uint32(value)).Convert(field.Type()))
			}
		case reflect.Uint64:
			if value, pie := strconv.ParseUint(tag, 0, 64); nil == pie {
				field.Set(reflect.ValueOf(value).Convert(field.Type()))
			}
		case reflect.Uintptr:
			if value, pie := strconv.ParseUint(tag, 0, strconv.IntSize); nil == pie {
				field.Set(reflect.ValueOf(uintptr(value)).Convert(field.Type()))
			}
		case reflect.Float32:
			if value, pfe := strconv.ParseFloat(tag, 32); nil == pfe {
				field.Set(reflect.ValueOf(float32(value)).Convert(field.Type()))
			}
		case reflect.Float64:
			if value, pfe := strconv.ParseFloat(tag, 64); nil == pfe {
				field.Set(reflect.ValueOf(value).Convert(field.Type()))
			}
		case reflect.String:
			field.Set(reflect.ValueOf(tag).Convert(field.Type()))
		case reflect.Slice:
			ref := reflect.New(field.Type())
			ref.Elem().Set(reflect.MakeSlice(field.Type(), 0, 0))
			if `` != tag && jsonSlice != tag {
				if err = json.Unmarshal(parseJson(tag), ref.Interface()); nil != err {
					return
				}
			}
			field.Set(ref.Elem().Convert(field.Type()))
		case reflect.Map:
			ref := reflect.New(field.Type())
			ref.Elem().Set(reflect.MakeMap(field.Type()))
			if `` != tag && jsonMap != tag {
				if err = json.Unmarshal(parseJson(tag), ref.Interface()); nil != err {
					return
				}
			}
			field.Set(ref.Elem().Convert(field.Type()))
		case reflect.Struct:
			if `` != tag && jsonStruct != tag {
				if err = json.Unmarshal(parseJson(tag), field.Addr().Interface()); nil != err {
					return
				}
			}
		case reflect.Ptr:
			field.Set(reflect.New(field.Type().Elem()))
		}
	}

	switch field.Kind() {
	case reflect.Ptr:
		if isInitial || field.Elem().Kind() == reflect.Struct {
			// 不关注错误，后面的代码必须执行
			_ = setField(field.Elem(), tag)
			_setter(field.Interface())
		}
	case reflect.Struct:
		if err = Set(field.Addr().Interface()); nil != err {
			return
		}
	case reflect.Slice:
		for index := 0; index < field.Len(); index++ {
			if err = setField(field.Index(index), tag); nil != err {
				return
			}
		}
	}

	return
}

func parseJson(from string) []byte {
	return []byte(strings.ReplaceAll(from, `'`, `"`))
}

func isInitialValue(field reflect.Value) bool {
	return reflect.DeepEqual(reflect.Zero(field.Type()).Interface(), field.Interface())
}

func isInitialField(field reflect.Value, tag string) (initial bool) {
	switch field.Kind() {
	case reflect.Struct:
		initial = true
	case reflect.Ptr:
		initial = !field.IsNil() && reflect.Struct == field.Elem().Kind()
	case reflect.Slice:
		initial = field.Len() > 0 || `` != tag
	default:
		initial = `` != tag
	}

	return
}
