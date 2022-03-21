package c0nvert

import (
	"reflect"
	"strconv"
)

type Floater interface {
	Float() float64
}

func toFloat[T any](src any, ns ...int) any {

	srcVal := reflect.ValueOf(src)

	if e, ok := src.(Floater); ok {
		srcVal = reflect.ValueOf(e.Float())
	}

	if isByteString(srcVal.Type()) {
		srcVal = reflect.ValueOf(toFloat[float64](ByteStringToString(src)))
	}

	if srcVal.Kind() == reflect.Bool {
		if srcVal.Bool() {
			srcVal = reflect.ValueOf(1)
		} else {
			srcVal = reflect.ValueOf(0)
		}
	}

	if srcVal.Kind() == reflect.String {
		v, _ := strconv.ParseFloat(srcVal.String(), 64)
		srcVal = reflect.ValueOf(v)
	}

	dstVal := reflect.ValueOf(new(T)).Elem()
	dstType := dstVal.Type()

	if srcVal.CanConvert(dstType) {
		srcVal = srcVal.Convert(dstType)
		val := RoundedFixed(srcVal.Float(), ns...)
		if dstType.Kind() == reflect.Float32 {
			return float32(val)
		}
		return val
	}

	return 0.0
}
