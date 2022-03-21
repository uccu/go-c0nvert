package c0nvert

import (
	"reflect"
	"strings"
)

type Booler interface {
	Bool() bool
}

func toBool(src any, fuzzy ...bool) any {

	srcVal := reflect.ValueOf(src)
	isfuzzy := true
	if len(fuzzy) == 0 || !fuzzy[0] {
		isfuzzy = false
	}

	if e, ok := srcVal.Interface().(Booler); ok {
		return e.Bool()
	}

	if srcVal.Kind() != reflect.String {

		if !isfuzzy || !isByteString(srcVal.Type()) {
			return !srcVal.IsZero()
		}

		strType := reflect.TypeOf(new(string)).Elem()
		srcVal = srcVal.Convert(strType)
	}

	str := srcVal.String()
	if !isfuzzy {
		return str != ""
	}

	return str != "" && str != "0" && strings.ToLower(str) != "false"
}
