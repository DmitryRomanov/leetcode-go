package diameter_of_binary_tree

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Deserializes your encoded data to tree.
func (codec *Codec) deserialize(data string) *TreeNode {
	elements := strings.Split(data, ",")
	if elements[0] == "" {
		return nil
	}

	value, _ := strconv.Atoi(elements[0])
	root := &TreeNode{
		Val: value,
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	i := 1
	for {
		if i == len(elements) || len(queue) == 0 {
			break
		}
		current := queue[0]
		queue = queue[1:]
		value, _ := strconv.Atoi(elements[i])
		current.Left = &TreeNode{
			Val: value,
		}
		queue = append(queue, current.Left)
		i++

		if i == len(elements) || len(queue) == 0 {
			break
		}

		value, _ = strconv.Atoi(elements[i])
		current.Right = &TreeNode{
			Val: value,
		}
		queue = append(queue, current.Right)
		i++
	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
