package non_overlapping_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEraseOverlapIntervalsExample1(t *testing.T) {
	assert := assert.New(t)
	result := eraseOverlapIntervals([][]int{
		{1, 2}, {2, 3}, {3, 4}, {1, 3},
	})
	assert.Equal(1, result)
}

func TestEraseOverlapIntervalsExample2(t *testing.T) {
	assert := assert.New(t)
	result := eraseOverlapIntervals([][]int{
		{1, 2}, {1, 2}, {1, 2},
	})
	assert.Equal(2, result)
}

func TestEraseOverlapIntervalsExample3(t *testing.T) {
	assert := assert.New(t)
	result := eraseOverlapIntervals([][]int{
		{1, 2}, {2, 3},
	})
	assert.Equal(0, result)
}

func TestEraseOverlapIntervalsExample4(t *testing.T) {
	assert := assert.New(t)
	result := eraseOverlapIntervals([][]int{
		{1, 100}, {11, 22}, {1, 11}, {2, 12},
	})
	assert.Equal(2, result)
}
