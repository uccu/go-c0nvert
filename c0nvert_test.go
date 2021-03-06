package c0nvert_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/uccu/go-c0nvert"
)

func TestInvalidToSlice(t *testing.T) {
	assert.Equal(t, ToSlice[[]int](nil), []int{})
}

type TestStruct struct {
	A int
}

func TestToStruct(t *testing.T) {
	assert.Equal(t, To[TestStruct](TestStruct{1}), TestStruct{1})
}

func TestOtherToStruct(t *testing.T) {
	assert.Equal(t, To[TestStruct](1), TestStruct{0})
}
func TestByteStringToString(t *testing.T) {
	assert.Equal(t, ByteStringToString([]int{1}), "")

}
func Test33(t *testing.T) {

	a := new([]byte)

	v := reflect.ValueOf(a).Elem()

	v = reflect.Append(v, reflect.ValueOf(byte(1)))

	fmt.Println(*a, v.Interface())

}
