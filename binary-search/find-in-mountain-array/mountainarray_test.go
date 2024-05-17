package mountainarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	arr := NewMountainArray([]int{1, 2, 3, 4, 5, 3, 1})
	result := findInMountainArray(3, arr)
	assert.Equal(2, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	arr := NewMountainArray([]int{0, 1, 2, 4, 2, 1})
	result := findInMountainArray(3, arr)
	assert.Equal(-1, result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	arr := NewMountainArray([]int{1, 5, 2})
	result := findInMountainArray(0, arr)
	assert.Equal(-1, result)
}

func TestExample4(t *testing.T) {
	assert := assert.New(t)
	arr := NewMountainArray([]int{1, 5, 2})
	result := findInMountainArray(1, arr)
	assert.Equal(0, result)
}

func TestExample5(t *testing.T) {
	assert := assert.New(t)
	arr := NewMountainArray([]int{1, 5, 2})
	result := findInMountainArray(2, arr)
	assert.Equal(2, result)
}
