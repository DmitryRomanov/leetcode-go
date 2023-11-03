package word_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := exist(
		[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
		"ABCCED",
	)
	assert.Equal(true, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := exist(
		[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
		"SEE",
	)
	assert.Equal(true, result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := exist(
		[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
		"ABCB",
	)
	assert.Equal(false, result)
}
