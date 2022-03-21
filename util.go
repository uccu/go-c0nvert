package c0nvert

import (
	"math"
	"reflect"
)

func RoundedFixed(val float64, ns ...int) float64 {
	if len(ns) == 0 {
		return val
	}

	shift := math.Pow(10, float64(ns[0]))
	fv := 0.0000000001 + val
	return math.Floor(fv*shift+.5) / shift
}

func isByteString(t reflect.Type) bool {
	if t.Kind() == reflect.Slice {
		return t.Elem().Kind() == reflect.Int32 || t.Elem().Kind() == reflect.Uint8
	}
	return false
}

func ByteStringToString(i any) string {

	if e, ok := i.([]byte); ok {
		return string(e)
	}

	if e, ok := i.([]rune); ok {
		return string(e)
	}

	return ""
}
