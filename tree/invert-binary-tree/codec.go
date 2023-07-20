package invert_binary_tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (codec *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	queue := []*TreeNode{}
	enqueue := func(node *TreeNode, queue *[]*TreeNode) {
		*queue = append(*queue, node)
	}
	dequeue := func(queue *[]*TreeNode) (node *TreeNode) {
		node, *queue = (*queue)[0], (*queue)[1:len(*queue)]
		return
	}

	enqueue(root, &queue)
	var result []string
	for len(queue) > 0 {
		node := dequeue(&queue)
		if node != nil {
			result = append(result, fmt.Sprint(node.Val))
			enqueue(node.Left, &queue)
			enqueue(node.Right, &queue)
		}
	}

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (codec *Codec) deserialize(data string) *TreeNode {
	elements := strings.Split(data, ",")
	if elements[0] == "" {
		return nil
	}

	value, _ := strconv.Atoi(elements[0])
	root := addVertex(nil, value)
	for _, v := range elements[1:] {
		value, _ := strconv.Atoi(v)
		root = addVertex(root, value)
	}

	return root
}

func addVertex(tree *TreeNode, val int) *TreeNode {
	if tree == nil {
		tree = &TreeNode{Val: val}
	}
	if tree.Val > val {
		tree.Left = addVertex(tree.Left, val)
		return tree
	} else if tree.Val < val {
		tree.Right = addVertex(tree.Right, val)
		return tree
	}

	return &TreeNode{Val: val}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
