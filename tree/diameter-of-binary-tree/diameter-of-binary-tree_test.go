package diameter_of_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	in := "1,2,3,4,5"

	codec := Constructor()
	rootDeserialized := codec.deserialize(in)
	diameter := diameterOfBinaryTree(rootDeserialized)
	assert.Equal(3, diameter)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	in := "1,2"

	codec := Constructor()
	rootDeserialized := codec.deserialize(in)
	diameter := diameterOfBinaryTree(rootDeserialized)
	assert.Equal(1, diameter)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	in := "1"

	codec := Constructor()
	rootDeserialized := codec.deserialize(in)
	diameter := diameterOfBinaryTree(rootDeserialized)
	assert.Equal(0, diameter)
}
