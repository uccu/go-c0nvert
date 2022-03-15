package c0nvert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/uccu/go-c0nvert"
)

type TestFloater struct{}

func (TestFloater) Float() float64 {
	return 1.2
}

func TestInvalidToFloat(t *testing.T) {
	assert.Equal(t, ToFloat(nil), 0.0)
}

func TestFloaterToFloat(t *testing.T) {
	assert.Equal(t, ToFloat(TestFloater{}), 1.2)
}

func TestBoolToFloat(t *testing.T) {
	assert.Equal(t, ToFloat(true), 1.0)
	assert.Equal(t, ToFloat(false), 0.0)
}

func TestIntToFloat(t *testing.T) {
	assert.Equal(t, ToFloat(2), 2.0)
}

func TestUintToFloat(t *testing.T) {
	assert.Equal(t, ToFloat(uint(3)), 3.0)
}

func TestFloatToFloat(t *testing.T) {
	assert.Equal(t, ToFloat(1.2), 1.2)
}

func TestStringToFloat(t *testing.T) {
	assert.Equal(t, ToFloat("1"), 1.0)

	assert.Equal(t, ToFloat([]byte("2.1")), 2.1)
	assert.Equal(t, ToFloat([]rune("3.1")), 3.1)

	assert.Equal(t, ToFloat("1.012", 2), 1.01)
	assert.Equal(t, ToFloat("1.019", 2), 1.02)
}

func TestToFloat32(t *testing.T) {
	assert.Equal(t, ToFloat32("1"), float32(1.0))
}

func TestOtherToFloat(t *testing.T) {
	assert.Equal(t, ToFloat([]int{}), 0.0)
}
