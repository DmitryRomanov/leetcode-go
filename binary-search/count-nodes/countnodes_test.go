package countnodes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountNodesExample1(t *testing.T) {
	assert := assert.New(t)
	result := countNodes(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	})
	assert.Equal(6, result)
}

func TestCountNodesExample2(t *testing.T) {
	assert := assert.New(t)
	result := countNodes(&TreeNode{})
	assert.Equal(1, result)
}

func TestCountNodesExample3(t *testing.T) {
	assert := assert.New(t)
	result := countNodes(&TreeNode{
		Val: 1,
	})
	assert.Equal(1, result)
}

func TestCountNodesExample4(t *testing.T) {
	assert := assert.New(t)
	result := countNodes(nil)
	assert.Equal(0, result)
}

func TestCountNodesExample5(t *testing.T) {
	assert := assert.New(t)
	result := countNodes(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	})
	assert.Equal(2, result)
}

func TestCountNodesExample6(t *testing.T) {
	assert := assert.New(t)
	result := countNodes(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	})
	assert.Equal(3, result)
}
