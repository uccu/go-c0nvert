package c0nvert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStringer struct{}

func (TestStringer) String() string {
	return "ab"
}

func TestInvalidToString(t *testing.T) {
	assert.Equal(t, ToString(nil), "")
}

func TestInterToString(t *testing.T) {
	assert.Equal(t, ToString(TestStringer{}), "ab")
}

func TestBoolToString(t *testing.T) {
	assert.Equal(t, ToString(true), "true")
	assert.Equal(t, ToString(false), "false")
	assert.Equal(t, ToString(false, 1), "FALSE")
	assert.Equal(t, ToString(true, true), "TRUE")
	assert.Equal(t, ToString(false, false), "false")
	assert.Equal(t, ToString(false, 0), "false")
}

func TestIntToString(t *testing.T) {
	assert.Equal(t, ToString(-2), "-2")
}

func TestUintToString(t *testing.T) {
	assert.Equal(t, ToString(uint(3)), "3")
}

func TestFloatToString(t *testing.T) {
	assert.Equal(t, ToString(3.1), "3.1")
	assert.Equal(t, ToString(3.122, 2), "3.12")
	assert.Equal(t, ToString(3.129, 2), "3.13")
}

func TestStringToString(t *testing.T) {
	assert.Equal(t, ToString("1"), "1")
}

func TestSliceToString(t *testing.T) {
	assert.Equal(t, ToString([]int{1, 2}, "."), "1.2")
	assert.Equal(t, ToString([]int32{1, 2}, "."), "1.2")
	assert.Equal(t, ToString([]byte("x2")), "x2")
	assert.Equal(t, ToString([]int32("xx")), "xx")
}

func TestOtherToString(t *testing.T) {
	assert.Equal(t, ToString(map[string]string{}), "")
}

func TestMoreElemToString(t *testing.T) {
	var a interface{}
	var b int = 1
	a = &b
	assert.Equal(t, ToString(&a), "1")
}
