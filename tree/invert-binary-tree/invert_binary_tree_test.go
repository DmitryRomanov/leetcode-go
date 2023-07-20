// https://leetcode.com/problems/invert-binary-tree/description/
package invert_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	in := "4,2,7,1,3,6,9"

	codec := Constructor()
	rootDeserialized := codec.deserialize(in)
	rootInverted := invertTree(rootDeserialized)
	assert.Equal("4,7,2,9,6,3,1", codec.serialize(rootInverted))
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	in := "2,1,3"

	codec := Constructor()
	rootDeserialized := codec.deserialize(in)
	rootInverted := invertTree(rootDeserialized)
	assert.Equal("2,3,1", codec.serialize(rootInverted))
}
