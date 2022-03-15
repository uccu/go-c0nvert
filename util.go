package c0nvert

import (
	"math"
	"reflect"
)

func ToElemReflectValue(v interface{}) reflect.Value {

	e, isReflectValue := v.(reflect.Value)

	if !isReflectValue {
		return ToElemReflectValue(reflect.ValueOf(v))
	}

	if e.Kind() == reflect.Ptr || e.Kind() == reflect.Interface {
		return ToElemReflectValue(e.Elem())
	}

	return e
}

func RoundedFixed(val float64, ns ...int) float64 {
	if len(ns) == 0 {
		return val
	}

	shift := math.Pow(10, float64(ns[0]))
	fv := 0.0000000001 + val
	return math.Floor(fv*shift+.5) / shift
}
