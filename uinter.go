package c0nvert

import (
	"reflect"
)

type Uinter interface {
	Uint() uint64
}

func toUint[T any](src interface{}) interface{} {

	srcVal := reflect.ValueOf(src)

	if e, ok := src.(Uinter); ok {
		srcVal = reflect.ValueOf(e.Uint())
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
		srcVal = reflect.ValueOf(toFloat[float64](src))
	}

	dstVal := reflect.ValueOf(new(T)).Elem()
	dstType := dstVal.Type()

	if srcVal.CanConvert(dstType) {
		return srcVal.Convert(dstType).Interface().(T)
	}

	return dstVal.Interface().(T)
}
