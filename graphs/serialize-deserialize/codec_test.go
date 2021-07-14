package serialize_deserialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	ans := "2,1,3"

	ser := Constructor()
	assert.Equal(ans, ser.serialize(root))

	deser := Constructor()
	assert.Equal(root, deser.deserialize(ans))
}

func TestSerializeExample2(t *testing.T) {
	assert := assert.New(t)
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	ans := "3,1,5,0,2,4,6"

	ser := Constructor()
	assert.Equal(ans, ser.serialize(root))

	deser := Constructor()
	assert.Equal(root, deser.deserialize(ans))
}
