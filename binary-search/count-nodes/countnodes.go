package countnodes

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Val == 0 && root.Right == nil && root.Left == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
