package subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetsExample1(t *testing.T) {
	assert := assert.New(t)
	result := subsets([]int{0})
	assert.Equal([][]int{{}, {0}}, result)
}

func TestSubsetsExample2(t *testing.T) {
	assert := assert.New(t)
	result := subsets([]int{1, 2, 3})
	assert.Equal([][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}, []int{3}, []int{1, 3}, []int{2, 3}, []int{1, 2, 3}}, result)
}

func TestSubsetsExample3(t *testing.T) {
	assert := assert.New(t)
	result := subsets([]int{3, 2, 4, 1})
	assert.Equal([][]int{[]int{}, []int{3}, []int{2}, []int{2, 3}, []int{4}, []int{3, 4}, []int{2, 4}, []int{2, 3, 4}, []int{1}, []int{1, 3}, []int{1, 2}, []int{1, 2, 3}, []int{1, 4}, []int{1, 3, 4}, []int{1, 2, 4}, []int{1, 2, 3, 4}}, result)
}
