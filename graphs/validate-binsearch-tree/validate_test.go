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

func TestIsValidBSTExample4(t *testing.T) {
	assert := assert.New(t)
	result := isValidBST(&TreeNode{
		Val: 120,
		Left: &TreeNode{
			Val: 70,
			Left: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val: 20,
				},
				Right: &TreeNode{
					Val: 55,
				},
			},
			Right: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val: 75,
				},
				Right: &TreeNode{
					Val: 110,
				},
			},
		},
		Right: &TreeNode{
			Val: 140,
			Left: &TreeNode{
				Val: 130,
				Left: &TreeNode{
					Val: 119,
				},
				Right: &TreeNode{
					Val: 135,
				},
			},
			Right: &TreeNode{
				Val: 160,
				Left: &TreeNode{
					Val: 150,
				},
				Right: &TreeNode{
					Val: 200,
				},
			},
		},
	})
	assert.False(result)
}

func TestIsValidBSTExample5(t *testing.T) {
	assert := assert.New(t)
	result := isValidBST(&TreeNode{
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
	})
	assert.True(result)
}
