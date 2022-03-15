package c0nvert

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Stringer interface {
	String() string
}

const (
	CASE_LOWER byte = iota
	CASE_UPPER
)

func ToString(i interface{}, subs ...interface{}) string {

	valueElement := ToElemReflectValue(i)
	valueTypeKind := valueElement.Kind()

	if valueTypeKind == reflect.Invalid || !valueElement.CanInterface() {
		return ""
	}

	if e, ok := valueElement.Interface().(Stringer); ok {
		return e.String()
	}

	if valueTypeKind == reflect.String {
		return valueElement.String()
	}

	if valueTypeKind == reflect.Bool {
		b := fmt.Sprintf("%v", valueElement.Bool())
		if len(subs) > 0 && ToBool(subs[0]) {
			return strings.ToUpper(b)
		}
		return b
	}

	if valueTypeKind <= reflect.Int64 {
		return strconv.FormatInt(valueElement.Int(), 10)
	}

	if valueTypeKind <= reflect.Uintptr {
		return strconv.FormatUint(valueElement.Uint(), 10)
	}

	if valueTypeKind <= reflect.Float64 {
		if len(subs) > 0 {
			if sub := ToInt(subs[0]); sub > -1 {
				return strconv.FormatFloat(valueElement.Float(), 'f', sub, 64)
			}
		}
		return fmt.Sprintf("%v", valueElement.Float())
	}

	if valueTypeKind == reflect.Slice {
		if len(subs) == 0 {
			if e, ok := valueElement.Interface().([]byte); ok { // []byte/[]uint8
				return string(e)
			}

			if e, ok := valueElement.Interface().([]rune); ok { // []rune/[]int32
				return string(e)
			}
		}
		return sliceToString(valueElement, subs...)
	}

	return ""
}

func sliceToString(valueElement reflect.Value, subs ...interface{}) string {

	sli := []string{}
	sub := ","
	if len(subs) > 0 {
		sub = ToString(subs[0])
		subs = subs[1:]
	}

	for i := 0; i < valueElement.Len(); i++ {
		value := ToString(valueElement.Index(i), subs...)
		sli = append(sli, value)
	}

	return strings.Join(sli, sub)
}
