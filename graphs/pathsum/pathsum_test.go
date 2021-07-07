package pathsum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasPathSumExample1(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}, 22)
	assert.True(result)
}

func TestHasPathSumExample2(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}, 5)
	assert.False(result)
}

func TestHasPathSumExample3(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}, 0)
	assert.False(result)
}

func TestHasPathSumExample4(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(&TreeNode{}, 1)
	assert.False(result)
}

func TestHasPathSumExample5(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(nil, 1)
	assert.False(result)
}

func TestHasPathSumExample6(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}, 1)
	assert.False(result)
}

func TestHasPathSumExample7(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}, 1)
	assert.False(result)
}

func TestHasPathSumExample8(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(&TreeNode{
		Val: -2,
		Right: &TreeNode{
			Val: -3,
		},
	}, -5)
	assert.True(result)
}

func TestHasPathSumExample9(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(&TreeNode{
		Val: -2,
		Right: &TreeNode{
			Val: -3,
		},
	}, -2)
	assert.False(result)
}

func TestHasPathSumExample10(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -1,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -2,
			},
		},
	}, -1)
	assert.True(result)
}
