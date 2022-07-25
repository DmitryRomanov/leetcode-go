package problem_2300

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessfulPairsExample1(t *testing.T) {
	assert := assert.New(t)
	result := successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7)
	assert.Equal([]int{4, 0, 3}, result)
}

func TestSuccessfulPairsExample2(t *testing.T) {
	assert := assert.New(t)
	result := successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16)
	assert.Equal([]int{2, 0, 2}, result)
}
