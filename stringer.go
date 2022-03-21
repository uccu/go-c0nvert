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

func toString(src any, subs ...any) any {

	srcVal := reflect.ValueOf(src)

	if e, ok := src.(Stringer); ok {
		srcVal = reflect.ValueOf(e.String())
	}

	if isByteString(srcVal.Type()) && len(subs) == 0 {
		return ByteStringToString(src)
	}

	if srcVal.Kind() == reflect.Bool {
		b := fmt.Sprintf("%v", srcVal.Bool())
		if len(subs) > 0 && To[bool](subs[0]) {
			return strings.ToUpper(b)
		}
		return b
	}

	if srcVal.CanInt() {
		return strconv.FormatInt(srcVal.Int(), 10)
	}

	if srcVal.CanUint() {
		return strconv.FormatUint(srcVal.Uint(), 10)
	}

	if srcVal.CanFloat() {
		if len(subs) > 0 {
			if sub := To[int](subs[0]); sub > -1 {
				return strconv.FormatFloat(srcVal.Float(), 'f', sub, 64)
			}
		}
		return fmt.Sprintf("%v", srcVal.Float())
	}

	if srcVal.Kind() == reflect.Slice {
		return sliceToString(srcVal, subs...)
	}

	dstVal := reflect.ValueOf(new(string)).Elem()
	dstType := dstVal.Type()

	if srcVal.CanConvert(dstType) {
		return srcVal.Convert(dstType).String()
	}

	return ""

}

func sliceToString(valueElement reflect.Value, subs ...any) string {

	sli := []string{}
	sub := ","
	if len(subs) > 0 {
		sub = To[string](subs[0])
		subs = subs[1:]
	}

	for i := 0; i < valueElement.Len(); i++ {
		value := To[string](valueElement.Index(i).Interface(), subs...)
		sli = append(sli, value)
	}

	return strings.Join(sli, sub)
}
