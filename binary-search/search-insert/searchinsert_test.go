package searchinsert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsertExample1(t *testing.T) {
	assert := assert.New(t)
	result := searchInsert([]int{1, 3, 5, 6}, 5)
	assert.Equal(2, result)
}

func TestSearchInsertExample2(t *testing.T) {
	assert := assert.New(t)
	result := searchInsert([]int{1, 3, 5, 6}, 2)
	assert.Equal(1, result)
}

func TestSearchInsertExample3(t *testing.T) {
	assert := assert.New(t)
	result := searchInsert([]int{1, 3, 5, 6}, 7)
	assert.Equal(4, result)
}

func TestSearchInsertExample4(t *testing.T) {
	assert := assert.New(t)
	result := searchInsert([]int{2, 3, 5, 6}, 1)
	assert.Equal(0, result)
}

func TestSearchInsertExample5(t *testing.T) {
	assert := assert.New(t)
	result := searchInsert([]int{-1, 3, 5, 6}, 0)
	assert.Equal(1, result)
}
