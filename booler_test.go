package c0nvert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/uccu/go-c0nvert"
)

type TestBooler struct{}

func (TestBooler) Bool() bool {
	return true
}

func TestInvalidToBool(t *testing.T) {
	assert.Equal(t, To[bool](nil), false)
}

func TestInterToBool(t *testing.T) {
	assert.Equal(t, To[bool](TestBooler{}), true)
}

func TestBoolToBool(t *testing.T) {
	assert.Equal(t, To[bool](true), true)
	assert.Equal(t, To[bool](false), false)
}

func TestIntToBool(t *testing.T) {
	assert.Equal(t, To[bool](1), true)
	assert.Equal(t, To[bool](0), false)
	assert.Equal(t, To[bool](-1), true)
}

func TestUintToBool(t *testing.T) {
	assert.Equal(t, To[bool](uint(1)), true)
	assert.Equal(t, To[bool](uint(0)), false)
}

func TestFloatToBool(t *testing.T) {
	assert.Equal(t, To[bool](float64(0)), false)
	assert.Equal(t, To[bool](float64(0.1)), true)
	assert.Equal(t, To[bool](float64(0.000000000000001)), true)
}

func TestStringToBool(t *testing.T) {
	assert.Equal(t, To[bool]("", true), false)
	assert.Equal(t, To[bool]("0", true), false)
	assert.Equal(t, To[bool]("false", true), false)
	assert.Equal(t, To[bool]("False", true), false)
	assert.Equal(t, To[bool]("1", true), true)
	assert.Equal(t, To[bool](""), false)
	assert.Equal(t, To[bool]("0"), true)
}

func TestSliceToBool(t *testing.T) {
	assert.Equal(t, To[bool]([]byte("x2")), true)
	assert.Equal(t, To[bool]([]int32("xx")), true)
	assert.Equal(t, To[bool]([]int32("false"), true), false)
	assert.Equal(t, To[bool]([]int32(nil), true), false)
	assert.Equal(t, To[bool]([]int32("false")), true)
	assert.Equal(t, To[bool]([]int32("")), true)
	assert.Equal(t, To[bool]([]int32(nil)), false)
}

func TestOtherToBool(t *testing.T) {
	assert.Equal(t, To[bool]([]int(nil)), false)
}
