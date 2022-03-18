package c0nvert

import (
	"reflect"
	"strings"
)

type Booler interface {
	Bool() bool
}

func toBool(src interface{}, fuzzy ...bool) interface{} {

	srcVal := reflect.ValueOf(src)
	srcType := srcVal.Type()

	if e, ok := srcVal.Interface().(Booler); ok {
		return e.Bool()
	}

	if srcType.Kind() != reflect.String {

		if srcType.Kind() != reflect.Slice || (srcType.Elem().Kind() != reflect.Uint8 && srcType.Elem().Kind() != reflect.Int32) {
			return !srcVal.IsZero()
		}

		strType := reflect.TypeOf(new(string)).Elem()
		srcVal = srcVal.Convert(strType)

	}

	str := srcVal.String()
	if len(fuzzy) == 0 || !fuzzy[0] {
		return str != ""
	}

	return str != "" && str != "0" && strings.ToLower(str) != "false"
}
