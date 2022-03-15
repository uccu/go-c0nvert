package c0nvert

import (
	"reflect"
	"strings"
)

type Booler interface {
	Bool() bool
}

func ToBool(i interface{}, subs ...interface{}) bool {

	valueElement := ToElemReflectValue(i)
	valueTypeKind := valueElement.Kind()

	if valueTypeKind == reflect.Invalid || !valueElement.CanInterface() {
		return false
	}

	if e, ok := valueElement.Interface().(Booler); ok {
		return e.Bool()
	}

	if valueTypeKind == reflect.Bool {
		return valueElement.Bool()
	}

	if valueTypeKind <= reflect.Int64 {
		return valueElement.Int() != 0
	}

	if valueTypeKind <= reflect.Uintptr {
		return valueElement.Uint() != 0
	}

	if valueTypeKind <= reflect.Float64 {
		return valueElement.Float() != 0
	}

	if valueTypeKind == reflect.String {
		str := valueElement.String()

		if len(subs) > 0 {
			if ToBool(subs[0]) {
				if str == "" || str == "0" || strings.ToLower(str) == "false" {
					return false
				}
				return true
			}
		}

		if str != "" {
			return true
		}
	}

	if valueTypeKind == reflect.Slice {
		if e, ok := valueElement.Interface().([]byte); ok {
			return ToBool(string(e), subs...)
		}

		if e, ok := valueElement.Interface().([]rune); ok {
			return ToBool(string(e), subs...)
		}
	}

	return false
}
