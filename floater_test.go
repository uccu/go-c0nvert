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
	assert.Equal(t, To[float64](nil), 0.0)
}

func TestFloaterToFloat(t *testing.T) {
	assert.Equal(t, To[float64](TestFloater{}), 1.2)
}

func TestBoolToFloat(t *testing.T) {
	assert.Equal(t, To[float64](true), 1.0)
	assert.Equal(t, To[float64](false), 0.0)
}

func TestIntToFloat(t *testing.T) {
	assert.Equal(t, To[float64](2), 2.0)
}

func TestUintToFloat(t *testing.T) {
	assert.Equal(t, To[float64](uint(3)), 3.0)
}

func TestFloatToFloat(t *testing.T) {
	assert.Equal(t, To[float64](1.2), 1.2)
}

func TestStringToFloat(t *testing.T) {
	assert.Equal(t, To[float64]("1"), 1.0)

	assert.Equal(t, To[float64]([]byte("2.1")), 2.1)
	assert.Equal(t, To[float64]([]rune("3.1")), 3.1)

	assert.Equal(t, To[float64]("1.012", 2), 1.01)
	assert.Equal(t, To[float64]("1.019", 2), 1.02)
}

func TestToFloat32(t *testing.T) {
	assert.Equal(t, To[float32]("1"), float32(1.0))
}

func TestOtherToFloat(t *testing.T) {
	assert.Equal(t, To[float64]([]int(nil)), 0.0)
}
