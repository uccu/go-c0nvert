package c0nvert

import (
	"reflect"
	"strconv"
)

type Floater interface {
	Float() float64
}

func ToFloat32(i interface{}) float32 {
	return float32(ToFloat(i))
}

func ToFloat(i interface{}, subs ...int) float64 {

	valueElement := ToElemReflectValue(i)
	valueTypeKind := valueElement.Kind()

	if valueTypeKind == reflect.Invalid || !valueElement.CanInterface() {
		return 0
	}

	if e, ok := valueElement.Interface().(Floater); ok {
		return RoundedFixed(e.Float(), subs...)
	}

	if valueTypeKind == reflect.Bool {
		if valueElement.Bool() {
			return 1
		}
		return 0
	}

	if valueTypeKind <= reflect.Int64 {
		return float64(valueElement.Int())
	}

	if valueTypeKind <= reflect.Uintptr {
		return float64(valueElement.Uint())
	}

	if valueTypeKind <= reflect.Float64 {
		return RoundedFixed(valueElement.Float(), subs...)
	}

	if valueTypeKind == reflect.String {
		v, _ := strconv.ParseFloat(valueElement.String(), 64)
		return RoundedFixed(v, subs...)
	}

	if valueTypeKind == reflect.Slice {

		if e, ok := valueElement.Interface().([]byte); ok {
			v, _ := strconv.ParseFloat(string(e), 64)
			return RoundedFixed(v, subs...)
		}

		if e, ok := valueElement.Interface().([]rune); ok {
			v, _ := strconv.ParseFloat(string(e), 64)
			return RoundedFixed(v, subs...)
		}

	}

	return 0
}
