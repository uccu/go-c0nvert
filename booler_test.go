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
	assert.Equal(t, ToBool(nil), false)
}

func TestInterToBool(t *testing.T) {
	assert.Equal(t, ToBool(TestBooler{}), true)
}

func TestBoolToBool(t *testing.T) {
	assert.Equal(t, ToBool(true), true)
	assert.Equal(t, ToBool(false), false)
}

func TestIntToBool(t *testing.T) {
	assert.Equal(t, ToBool(1), true)
	assert.Equal(t, ToBool(0), false)
	assert.Equal(t, ToBool(-1), true)
}

func TestUintToBool(t *testing.T) {
	assert.Equal(t, ToBool(uint(1)), true)
	assert.Equal(t, ToBool(uint(0)), false)
}

func TestFloatToBool(t *testing.T) {
	assert.Equal(t, ToBool(float64(0)), false)
	assert.Equal(t, ToBool(float64(0.1)), true)
	assert.Equal(t, ToBool(float64(0.000000000000001)), true)
}

func TestStringToBool(t *testing.T) {
	assert.Equal(t, ToBool("", true), false)
	assert.Equal(t, ToBool("0", true), false)
	assert.Equal(t, ToBool("false", true), false)
	assert.Equal(t, ToBool("False", true), false)
	assert.Equal(t, ToBool("1", true), true)
	assert.Equal(t, ToBool(""), false)
	assert.Equal(t, ToBool("0"), true)
}

func TestSliceToBool(t *testing.T) {
	assert.Equal(t, ToBool([]byte("x2")), true)
	assert.Equal(t, ToBool([]int32("xx")), true)
	assert.Equal(t, ToBool([]int32("false"), true), false)
}

func TestOtherToBool(t *testing.T) {
	assert.Equal(t, ToBool([]int{}), false)
}
