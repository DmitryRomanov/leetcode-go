package two_best_non_overlapping_events

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxTwoEventsExample1(t *testing.T) {
	assert := assert.New(t)
	result := maxTwoEvents([][]int{
		{1, 3, 2}, {4, 5, 2}, {2, 4, 3},
	})
	assert.Equal(4, result)
}

func TestMaxTwoEventsExample2(t *testing.T) {
	assert := assert.New(t)
	result := maxTwoEvents([][]int{
		{1, 3, 2}, {4, 5, 2}, {1, 5, 5},
	})
	assert.Equal(5, result)
}

func TestMaxTwoEventsExample3(t *testing.T) {
	assert := assert.New(t)
	result := maxTwoEvents([][]int{
		{1, 5, 3}, {1, 5, 1}, {6, 6, 5},
	})
	assert.Equal(8, result)
}

func TestMaxTwoEventsExample4(t *testing.T) {
	assert := assert.New(t)
	result := maxTwoEvents([][]int{
		{1, 1, 1}, {1, 1, 1},
	})
	assert.Equal(1, result)
}

func TestMaxTwoEventsExample5(t *testing.T) {
	assert := assert.New(t)
	result := maxTwoEvents([][]int{
		{19, 36, 24}, {70, 90, 11}, {61, 78, 36}, {38, 38, 70}, {39, 83, 72}, {8, 46, 5}, {64, 69, 49}, {88, 89, 39}, {53, 77, 24}, {35, 76, 26},
	})
	assert.Equal(142, result)
}
