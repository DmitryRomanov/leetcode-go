package remove_dublicates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesExample1(t *testing.T) {
	assert := assert.New(t)
	result := removeDuplicates("abbaca")
	assert.Equal("ca", result)
}

func TestRemoveDuplicatesExample2(t *testing.T) {
	assert := assert.New(t)
	result := removeDuplicates("azxxzy")
	assert.Equal("ay", result)
}
