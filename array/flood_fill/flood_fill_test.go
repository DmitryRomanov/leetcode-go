package flood_fill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloodFillExample1(t *testing.T) {
	assert := assert.New(t)
	result := floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2)
	assert.Equal([][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}, result)
}

func TestFloodFillExample2(t *testing.T) {
	assert := assert.New(t)
	result := floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0)
	assert.Equal([][]int{{0, 0, 0}, {0, 0, 0}}, result)
}
