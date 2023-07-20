package quick_sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayExample1(t *testing.T) {
	assert := assert.New(t)
	result := sortArray([]int{5, 2, 3, 1})
	assert.Equal([]int{1, 2, 3, 5}, result)
}

func TestSortArrayExample2(t *testing.T) {
	assert := assert.New(t)
	result := sortArray([]int{5, 1, 1, 2, 0, 0})
	assert.Equal([]int{0, 0, 1, 1, 2, 5}, result)
}
