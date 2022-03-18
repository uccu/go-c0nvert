package c0nvert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/uccu/go-c0nvert"
)

type TestInter struct{}

func (TestInter) Int() int64 {
	return 3
}

func TestInvalidToInt(t *testing.T) {
	assert.Equal(t, To[int](nil), 0)
}

func TestInterToInt(t *testing.T) {
	assert.Equal(t, To[int](TestInter{}), 3)
}

func TestBoolToInt(t *testing.T) {
	assert.Equal(t, To[int](true), 1)
	assert.Equal(t, To[int](false), 0)
}

func TestIntToInt(t *testing.T) {
	assert.Equal(t, To[int](2), 2)
}

func TestUintToInt(t *testing.T) {
	assert.Equal(t, To[int](uint(3)), 3)
}

func TestFloatToInt(t *testing.T) {
	assert.Equal(t, To[int](3.1), 3)
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, To[int]("1"), 1)

	assert.Equal(t, To[int]([]byte("2.1")), 2)
	assert.Equal(t, To[int]([]rune("3.1")), 3)

	assert.Equal(t, To[int]("1.012"), 1)
	assert.Equal(t, To[int]("1.019"), 1)
}

func TestToIntOther(t *testing.T) {
	assert.Equal(t, To[int32]("1"), int32(1))
	assert.Equal(t, To[int64]("1"), int64(1))
	assert.Equal(t, To[int8]("1"), int8(1))
	assert.Equal(t, To[int16]("1"), int16(1))
}

func TestOtherToInt(t *testing.T) {
	assert.Equal(t, To[int]([]int{}), 0)
}
