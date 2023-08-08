// https://leetcode.com/problems/diameter-of-binary-tree/
package diameter_of_binary_tree

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		result = max(result, left+right)
		return 1 + max(left, right)
	}
	dfs(root)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
