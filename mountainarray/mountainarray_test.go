package mountainarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeakIndexInMountainArrayExample1(t *testing.T) {
	assert := assert.New(t)
	result := peakIndexInMountainArray([]int{0, 1, 0})
	assert.Equal(1, result)
}

func TestPeakIndexInMountainArrayExample2(t *testing.T) {
	assert := assert.New(t)
	result := peakIndexInMountainArray([]int{0, 2, 1, 0})
	assert.Equal(1, result)
}

func TestPeakIndexInMountainArrayExample3(t *testing.T) {
	assert := assert.New(t)
	result := peakIndexInMountainArray([]int{0, 10, 5, 2})
	assert.Equal(1, result)
}

func TestPeakIndexInMountainArrayExample4(t *testing.T) {
	assert := assert.New(t)
	result := peakIndexInMountainArray([]int{3, 4, 5, 1})
	assert.Equal(2, result)
}

func TestPeakIndexInMountainArrayExample5(t *testing.T) {
	assert := assert.New(t)
	result := peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19})
	assert.Equal(2, result)
}

func TestPeakIndexInMountainArrayExample6(t *testing.T) {
	assert := assert.New(t)
	result := peakIndexInMountainArray([]int{24, 69, 99, 100, 99, 79, 78, 67, 36, 26, 19})
	assert.Equal(3, result)
}

func TestPeakIndexInMountainArrayExample7(t *testing.T) {
	assert := assert.New(t)
	result := peakIndexInMountainArray([]int{3, 5, 3, 2, 0})
	assert.Equal(1, result)
}
