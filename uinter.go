package c0nvert

import (
	"reflect"
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

	return uint64(sliceToFloat(valueElement))

}
