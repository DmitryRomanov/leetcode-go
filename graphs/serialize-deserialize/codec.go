package serialize_deserialize

import (
	"fmt"
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

	var result []string
	result = append(result, fmt.Sprint(root.Val))

	leftResult := codec.serialize(root.Left)
	if len(leftResult) > 0 {
		result = append(result, leftResult)
	}

	rightResult := codec.serialize(root.Right)
	if len(rightResult) > 0 {
		result = append(result, rightResult)
	}

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (codec *Codec) deserialize(data string) *TreeNode {
	//elements:=strings.Split(data,",")
	return &TreeNode{}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
