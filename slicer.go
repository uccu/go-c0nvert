package c0nvert

import (
	"strings"
)

func splitToSliceString(i string, subs ...string) []string {

	if i == "" {
		return []string{}
	}

	sli := []string{i}
	if len(subs) == 0 {
		subs = []string{","}
	}

	for _, v := range subs {
		s := []string{}
		for _, v2 := range sli {
			s = append(s, strings.Split(v2, v)...)
		}
		sli = s
	}

	return sli

}

func Split[List []T, T any](i string, subs ...string) List {
	stringSlice := splitToSliceString(i, subs...)
	slice := make([]T, len(stringSlice))

	for k := range stringSlice {
		slice[k] = To[T](stringSlice[k])
	}
	return slice
}
