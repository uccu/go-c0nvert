package c0nvert

import (
	"reflect"
	"strconv"
)

type Inter interface {
	Int() int64
}

func ToInt(i interface{}) int {
	return int(ToInt64(i))
}

func ToInt8(i interface{}) int8 {
	return int8(ToInt64(i))
}

func ToInt16(i interface{}) int16 {
	return int16(ToInt64(i))
}

func ToInt32(i interface{}) int32 {
	return int32(ToInt64(i))
}

func ToInt64(i interface{}) int64 {

	valueElement := ToElemReflectValue(i)
	valueTypeKind := valueElement.Kind()

	if valueTypeKind == reflect.Invalid || !valueElement.CanInterface() {
		return 0
	}

	if e, ok := valueElement.Interface().(Inter); ok {
		return e.Int()
	}

	if valueTypeKind == reflect.Bool {
		if valueElement.Bool() {
			return 1
		}
		return 0
	}

	if valueTypeKind <= reflect.Int64 {
		return valueElement.Int()
	}

	if valueTypeKind <= reflect.Uintptr {
		return int64(valueElement.Uint())
	}

	if valueTypeKind <= reflect.Float64 {
		return int64(valueElement.Float())
	}

	if valueTypeKind == reflect.String {
		v, _ := strconv.ParseFloat(valueElement.String(), 64)
		return int64(v)
	}

	if valueTypeKind == reflect.Slice {

		if e, ok := valueElement.Interface().([]byte); ok {
			v, _ := strconv.ParseFloat(string(e), 64)
			return int64(v)
		}

		if e, ok := valueElement.Interface().([]rune); ok {
			v, _ := strconv.ParseFloat(string(e), 64)
			return int64(v)
		}

	}

	return 0
}
