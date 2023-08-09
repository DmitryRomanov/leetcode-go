package middle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := middleNode(build([]int{1, 2, 3, 4, 5}))
	assert.Equal(build([]int{3, 4, 5}), result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := middleNode(build([]int{1, 2, 3, 4, 5, 6}))
	assert.Equal(build([]int{4, 5, 6}), result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := middleNode(build([]int{1}))
	assert.Equal(build([]int{1}), result)
}

func TestExample4(t *testing.T) {
	assert := assert.New(t)
	result := middleNode(build([]int{}))
	assert.Equal(build([]int{}), result)
}

func TestBuild1(t *testing.T) {
	assert := assert.New(t)
	result := build([]int{1, 2, 3})
	assert.Equal(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}, result)
}

func TestBuild2(t *testing.T) {
	assert := assert.New(t)
	result := build([]int{})
	assert.Equal((*ListNode)(nil), result)
}
