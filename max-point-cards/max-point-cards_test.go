package maxpointcards

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxScoreExample1(t *testing.T) {
	assert := assert.New(t)
	result := maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3)
	assert.Equal(12, result)
}

func TestMaxScoreExample2(t *testing.T) {
	assert := assert.New(t)
	result := maxScore([]int{2, 2, 2}, 2)
	assert.Equal(4, result)
}

func TestMaxScoreExample3(t *testing.T) {
	assert := assert.New(t)
	result := maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7)
	assert.Equal(55, result)
}

func TestMaxScoreExample4(t *testing.T) {
	assert := assert.New(t)
	result := maxScore([]int{1, 1000, 1}, 1)
	assert.Equal(1, result)
}

func TestMaxScoreExample5(t *testing.T) {
	assert := assert.New(t)
	result := maxScore([]int{1, 79, 80, 1, 1, 1, 200, 1}, 3)
	assert.Equal(202, result)
}
