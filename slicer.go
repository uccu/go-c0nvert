package c0nvert

import (
	"strconv"
	"strings"
)

func ToSliceString(i string, subs ...string) []string {

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

func ToSliceInt64(i string, subs ...string) []int64 {

	stringSlice := ToSliceString(i, subs...)
	intSlice := make([]int64, len(stringSlice))

	for k, v := range stringSlice {
		intSlice[k], _ = strconv.ParseInt(v, 10, 64)
	}

	return intSlice
}

func ToSliceInt(i string, subs ...string) []int {

	stringSlice := ToSliceString(i, subs...)
	intSlice := make([]int, len(stringSlice))

	for k, v := range stringSlice {
		intSlice[k], _ = strconv.Atoi(v)
	}

	return intSlice
}
