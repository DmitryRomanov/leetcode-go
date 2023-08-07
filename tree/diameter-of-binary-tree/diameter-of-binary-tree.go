// https://leetcode.com/problems/diameter-of-binary-tree/
package diameter_of_binary_tree

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(a [3]int) int {
		max := 0
		for _, v := range a {
			if v > max {
				max = v
			}
		}
		return max
	}

	return max([3]int{
		depth(root.Left) +
			depth(root.Right),
		diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right),
	})
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}
