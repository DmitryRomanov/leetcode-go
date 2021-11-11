package subsets_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetsWithDupExample1(t *testing.T) {
	assert := assert.New(t)
	result := subsetsWithDup([]int{1, 2, 2})
	assert.Equal([][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}, []int{2, 2}, []int{1, 2, 2}}, result)
}

func TestSubsetsWithDupExample2(t *testing.T) {
	assert := assert.New(t)
	result := subsetsWithDup([]int{0})
	assert.Equal([][]int{{}, {0}}, result)
}

func TestSubsetsWithDupExample3(t *testing.T) {
	assert := assert.New(t)
	result := subsetsWithDup([]int{1, 2, 3})
	assert.Equal([][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}, []int{3}, []int{1, 3}, []int{2, 3}, []int{1, 2, 3}}, result)
}

func TestSubsetsWithDupExample4(t *testing.T) {
	assert := assert.New(t)
	result := subsetsWithDup([]int{1, 2, 2, 3})
	assert.Equal([][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}, []int{2, 2}, []int{1, 2, 2}, []int{3}, []int{1, 3}, []int{2, 3}, []int{1, 2, 3}, []int{2, 2, 3}, []int{1, 2, 2, 3}}, result)
}

func TestSubsetsWithDupExample5(t *testing.T) {
	assert := assert.New(t)
	result := subsetsWithDup([]int{4, 4, 4, 1, 4})
	assert.Equal([][]int{[]int{}, []int{4}, []int{4, 4}, []int{4, 4, 4}, []int{1}, []int{1, 4}, []int{1, 4, 4}, []int{1, 4, 4, 4}, []int{4, 4, 4, 4}, []int{1, 4, 4, 4, 4}}, result)
}

func TestSubsetsWithDupExample6(t *testing.T) {
	assert := assert.New(t)
	result := subsetsWithDup([]int{2, 1, 2, 1, 3})
	assert.Equal([][]int{[]int{}, []int{2}, []int{1}, []int{1, 2}, []int{2, 2}, []int{1, 2, 2}, []int{1, 1}, []int{1, 1, 2}, []int{1, 1, 2, 2}, []int{3}, []int{2, 3}, []int{1, 3}, []int{1, 2, 3}, []int{2, 2, 3}, []int{1, 2, 2, 3}, []int{1, 1, 3}, []int{1, 1, 2, 3}, []int{1, 1, 2, 2, 3}}, result)
}

func TestSubsetsWithDupExample7(t *testing.T) {
	assert := assert.New(t)
	result := subsetsWithDup([]int{3, 2, 4, 1})
	assert.Equal([][]int{[]int{}, []int{3}, []int{2}, []int{2, 3}, []int{4}, []int{3, 4}, []int{2, 4}, []int{2, 3, 4}, []int{1}, []int{1, 3}, []int{1, 2}, []int{1, 2, 3}, []int{1, 4}, []int{1, 3, 4}, []int{1, 2, 4}, []int{1, 2, 3, 4}}, result)
}
