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
	assert.Equal([]int{-1, -1}, result)
}

func TestSearchRangeExample3(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{}, 0)
	assert.Equal([]int{-1, -1}, result)
}

func TestSearchRangeExample4(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{5, 7, 7, 7, 8, 8, 10}, 7)
	assert.Equal([]int{1, 3}, result)
}

func TestSearchRangeExample5(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{5, 7, 8, 8, 10}, 7)
	assert.Equal([]int{1, 1}, result)
}

func TestSearchRangeExample6(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	assert.Equal([]int{-1, -1}, result)
}

func TestSearchRangeExample7(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{1}, 1)
	assert.Equal([]int{0, 0}, result)
}

func TestSearchRangeExample8(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{2, 2}, 2)
	assert.Equal([]int{0, 1}, result)
}

func TestSearchRangeExample9(t *testing.T) {
	assert := assert.New(t)
	result := searchRange([]int{1, 3}, 1)
	assert.Equal([]int{0, 0}, result)
}

func TestFirstEqualExample1(t *testing.T) {
	assert := assert.New(t)
	result := firstEqual([]int{5, 7, 7}, 7)
	assert.Equal(1, result)
}

func TestFirstEqualExample2(t *testing.T) {
	assert := assert.New(t)
	result := firstEqual([]int{5, 5, 5}, 7)
	assert.Equal(-1, result)
}

func TestFirstEqualExample3(t *testing.T) {
	assert := assert.New(t)
	result := firstEqual([]int{1}, 1)
	assert.Equal(0, result)
}

func TestLastEqualExample1(t *testing.T) {
	assert := assert.New(t)
	result := lastEqual([]int{7, 7, 9, 10}, 7)
	assert.Equal(1, result)
}

func TestLastEqualExample2(t *testing.T) {
	assert := assert.New(t)
	result := lastEqual([]int{5, 5, 5}, 7)
	assert.Equal(-1, result)
}

func TestFindOneExample1(t *testing.T) {
	assert := assert.New(t)
	result := findOne([]int{1}, 1)
	assert.Equal(0, result)
}

func TestFindOneExample2(t *testing.T) {
	assert := assert.New(t)
	result := findOne([]int{1, 3}, 1)
	assert.Equal(0, result)
}
