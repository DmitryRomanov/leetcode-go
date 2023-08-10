package maximum_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	assert.Equal(6, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := maxSubArray([]int{1})
	assert.Equal(1, result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := maxSubArray([]int{5, 4, -1, 7, 8})
	assert.Equal(23, result)
}

func TestExample4(t *testing.T) {
	assert := assert.New(t)
	result := maxSubArray([]int{-1})
	assert.Equal(-1, result)
}

func TestExample5(t *testing.T) {
	assert := assert.New(t)
	result := maxSubArray([]int{1})
	assert.Equal(1, result)
}
