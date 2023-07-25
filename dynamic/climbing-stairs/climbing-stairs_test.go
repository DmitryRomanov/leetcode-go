package climbing_stairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := climbStairs(2)
	assert.Equal(2, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := climbStairs(3)
	assert.Equal(3, result)
}
