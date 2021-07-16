package serialize_deserialize

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
		return "null"
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
	elements := strings.Split(data, ",")
	root, _ := deserializeNode(elements)
	return root
}

func deserializeNode(elements []string) (*TreeNode, []string) {
	root := &TreeNode{}
	if len(elements) > 0 {
		value, err := strconv.Atoi(elements[0])
		if err != nil {
			return nil, elements[1:]
		}
		root.Val = value

		var rightElementsString, otherLeafString []string
		root.Left, rightElementsString = deserializeNode(elements[1:])
		root.Right, otherLeafString = deserializeNode(rightElementsString)
		return root, otherLeafString
	}
	return nil, []string{}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
