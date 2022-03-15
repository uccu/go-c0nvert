package c0nvert

import (
	"reflect"
	"strconv"
)

type Uinter interface {
	Uint() uint64
}

func ToUint(i interface{}) uint {
	return uint(ToUint64(i))
}

func ToUint8(i interface{}) uint8 {
	return uint8(ToUint64(i))
}

func ToUint16(i interface{}) uint16 {
	return uint16(ToUint64(i))
}

func ToUint32(i interface{}) uint32 {
	return uint32(ToUint64(i))
}

func ToUintptr(i interface{}) uintptr {
	return uintptr(ToUint64(i))
}

func ToUint64(i interface{}) uint64 {

	valueElement := ToElemReflectValue(i)
	valueTypeKind := valueElement.Kind()

	if valueTypeKind == reflect.Invalid || !valueElement.CanInterface() {
		return 0
	}

	if e, ok := valueElement.Interface().(Uinter); ok {
		return e.Uint()
	}

	if valueTypeKind == reflect.Bool {
		if valueElement.Bool() {
			return 1
		}
		return 0
	}

	if valueTypeKind <= reflect.Int64 {
		return uint64(valueElement.Int())
	}

	if valueTypeKind <= reflect.Uintptr {
		return valueElement.Uint()
	}

	if valueTypeKind <= reflect.Float64 {
		return uint64(valueElement.Float())
	}

	if valueTypeKind == reflect.String {
		v, _ := strconv.ParseFloat(valueElement.String(), 64)
		return uint64(v)
	}

	if valueTypeKind == reflect.Slice {

		if e, ok := valueElement.Interface().([]byte); ok {
			v, _ := strconv.ParseFloat(string(e), 64)
			return uint64(v)
		}

		if e, ok := valueElement.Interface().([]rune); ok {
			v, _ := strconv.ParseFloat(string(e), 64)
			return uint64(v)
		}

	}

	return 0
}
