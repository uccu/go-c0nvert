package c0nvert

import (
	"math"
	"reflect"
	"strconv"
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

func sliceToFloat(valueElement reflect.Value) float64 {

	if valueElement.Kind() == reflect.String {
		v, _ := strconv.ParseFloat(valueElement.String(), 64)
		return v
	}

	if valueElement.Kind() == reflect.Slice {

		if e, ok := valueElement.Interface().([]byte); ok {
			v, _ := strconv.ParseFloat(string(e), 64)
			return v
		}

		if e, ok := valueElement.Interface().([]rune); ok {
			v, _ := strconv.ParseFloat(string(e), 64)
			return v
		}
	}

	return 0
}
