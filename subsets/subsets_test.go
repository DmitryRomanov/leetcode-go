package subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeakIndexInMountainArrayExample1(t *testing.T) {
	assert := assert.New(t)
	result := subsets([]int{0})
	assert.Equal([][]int{{}, {0}}, result)
}
