package domino_and_tromino_tiling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := numTilings(3)
	assert.Equal(5, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := numTilings(1)
	assert.Equal(1, result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := numTilings(2)
	assert.Equal(2, result)
}

func TestExample4(t *testing.T) {
	assert := assert.New(t)
	result := numTilings(4)
	assert.Equal(11, result)
}
