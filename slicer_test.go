package c0nvert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/uccu/go-c0nvert"
)

func TestEmptyStringToSliceString(t *testing.T) {
	assert.Equal(t, ToSliceString(""), []string{})
}

func TestToSliceStringWithNoSubs(t *testing.T) {
	assert.Equal(t, ToSliceString("1,2"), []string{"1", "2"})
}

func TestToSliceStringWithOneSubs(t *testing.T) {
	assert.Equal(t, ToSliceString("1.2", "."), []string{"1", "2"})
}

func TestToSliceStringWithMoreSubs(t *testing.T) {
	assert.Equal(t, ToSliceString("1.2", "."), []string{"1", "2"})
}

func TestToSliceInt64(t *testing.T) {
	assert.Equal(t, ToSliceInt64("1,2"), []int64{1, 2})
}
func TestToSliceInt(t *testing.T) {
	assert.Equal(t, ToSliceInt("1,2"), []int{1, 2})
}
