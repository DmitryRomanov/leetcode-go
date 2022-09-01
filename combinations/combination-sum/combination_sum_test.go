package combination_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSumExample1(t *testing.T) {
	assert := assert.New(t)
	result := combinationSum([]int{2, 3, 6, 7}, 7)

	assert.Equal([][]int{{2, 2, 3}, {7}}, result)
}

func TestCombinationSumExample2(t *testing.T) {
	assert := assert.New(t)
	result := combinationSum([]int{2, 3, 5}, 8)

	assert.Equal([][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, result)
}

func TestCombinationSumExample3(t *testing.T) {
	assert := assert.New(t)
	result := combinationSum([]int{2}, 1)

	assert.Equal([][]int{}, result)
}
