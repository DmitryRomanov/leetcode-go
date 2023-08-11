package insert_interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	assert.Equal([][]int{{1, 5}, {6, 9}}, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	assert.Equal([][]int{{1, 2}, {3, 10}, {12, 16}}, result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := insert([][]int{}, []int{2, 5})
	assert.Equal([][]int{{2, 5}}, result)
}

func TestExample4(t *testing.T) {
	assert := assert.New(t)
	result := insert([][]int{{1, 7}}, []int{2, 5})
	assert.Equal([][]int{{1, 7}}, result)
}

func TestExample5(t *testing.T) {
	assert := assert.New(t)
	result := insert([][]int{{1, 5}}, []int{2, 7})
	assert.Equal([][]int{{1, 7}}, result)
}
