package c0nvert

import (
	"reflect"
)

func ToSlice[List []T, T any](src any, subs ...any) List {

	dst := make(List, 0)

	srcVal := reflect.ValueOf(src)
	if !srcVal.IsValid() || !srcVal.CanInterface() || srcVal.Kind() != reflect.Slice {
		return dst
	}

	for k := 0; k < srcVal.Len(); k++ {
		dst = append(dst, To[T](srcVal.Index(k).Interface()))
	}

	return dst
}

func To[T any](src any, subs ...any) T {

	dst := new(T)
	dstVal := reflect.ValueOf(dst).Elem()

	srcVal := reflect.ValueOf(src)
	if !srcVal.IsValid() || !srcVal.CanInterface() {
		return dstVal.Interface().(T)
	}

	switch any(*new(T)).(type) {
	case bool:
		*dst = toBool(src, ToSlice[[]bool](subs)...).(T)
	case int, int8, int16, int32, int64:
		*dst = toInt[T](src).(T)
	case uint, uint8, uint16, uint32, uint64, uintptr:
		*dst = toUint[T](src).(T)
	case float64, float32:
		*dst = toFloat[T](src, ToSlice[[]int](subs)...).(T)
	case string:
		*dst = toString(src, subs...).(T)
	default:
		dst = nil
	}

	if dst != nil {
		return *dst
	}

	dstType := dstVal.Type()
	if srcVal.CanConvert(dstType) {
		return srcVal.Convert(dstType).Interface().(T)
	}

	return dstVal.Interface().(T)
}
