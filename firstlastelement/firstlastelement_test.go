package firstlastelement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchRangeExample1(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	assert.Equal([]int{3, 4}, result)
}

func TestSearchRangeExample2(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	assert.Equal([]int{-1, 1}, result)
}

func TestSearchRangeExample3(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{}, 0)
	assert.Equal([]int{-1, 1}, result)
}

func TestSearchRangeExample4(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{5, 7, 7, 7, 8, 8, 10}, 7)
	assert.Equal([]int{1, 3}, result)
}
