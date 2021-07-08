package validate_binsearch_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidBSTExample1(t *testing.T) {
	assert := assert.New(t)
	result := isValidBST(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	})
	assert.True(result)
}

func TestIsValidBSTExample2(t *testing.T) {
	assert := assert.New(t)
	result := isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	})
	assert.False(result)
}

func TestIsValidBSTExample3(t *testing.T) {
	assert := assert.New(t)
	result := isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	})
	assert.False(result)
}
