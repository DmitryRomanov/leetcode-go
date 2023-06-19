package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchExample1(t *testing.T) {
	assert := assert.New(t)
	result := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	assert.Equal(4, result)
}

func TestSearchExample2(t *testing.T) {
	assert := assert.New(t)
	result := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	assert.Equal(-1, result)
}
