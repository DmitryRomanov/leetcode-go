package majority_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := majorityElement([]int{3, 2, 3})
	assert.Equal(3, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	assert.Equal(2, result)
}
