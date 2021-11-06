package subsets_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetsWithDupExample1(t *testing.T) {
	assert := assert.New(t)
	result := subsetsWithDup([]int{1, 2, 2})
	assert.Equal([][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}, result)
}

func TestSubsetsWithDupExample2(t *testing.T) {
	assert := assert.New(t)
	result := subsetsWithDup([]int{0})
	assert.Equal([][]int{{}, {0}}, result)
}
