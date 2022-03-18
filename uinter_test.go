package c0nvert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/uccu/go-c0nvert"
)

type TestUinter struct{}

func (TestUinter) Uint() uint64 {
	return 3
}

func TestInvalidToUint(t *testing.T) {
	assert.Equal(t, To[uint](nil), uint(0))
}

func TestInterToUint(t *testing.T) {
	assert.Equal(t, To[uint](TestUinter{}), uint(3))
}

func TestBoolToUint(t *testing.T) {
	assert.Equal(t, To[uint](true), uint(1))
	assert.Equal(t, To[uint](false), uint(0))
}

func TestIntToUint(t *testing.T) {
	assert.Equal(t, To[uint](2), uint(2))
}

func TestUintToUint(t *testing.T) {
	assert.Equal(t, To[uint](uint(3)), uint(3))
}

func TestFloatToUint(t *testing.T) {
	assert.Equal(t, To[uint](3.1), uint(3))
}

func TestStringToUint(t *testing.T) {
	assert.Equal(t, To[uint]("1"), uint(1))

	assert.Equal(t, To[uint]([]byte("2.1")), uint(2))
	assert.Equal(t, To[uint]([]rune("3.1")), uint(3))

	assert.Equal(t, To[uint]("1.012"), uint(1))
	assert.Equal(t, To[uint]("1.019"), uint(1))
}

func TestToUintOther(t *testing.T) {
	assert.Equal(t, To[uint32]("1"), uint32(1))
	assert.Equal(t, To[uint64]("1"), uint64(1))
	assert.Equal(t, To[uint8]("1"), uint8(1))
	assert.Equal(t, To[uint16]("1"), uint16(1))
	assert.Equal(t, To[uintptr]("1"), uintptr(1))
}

func TestOtherToUint(t *testing.T) {
	assert.Equal(t, To[uint]([]int{}), uint(0))
}
