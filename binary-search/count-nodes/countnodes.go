package countnodes

func countNodes(root *TreeNode) int {
	if root.Right != nil {
		return 2 * (1 + calculateByLeftNodes(root.Right))
	}

	if root.Val != 0 {
		return 1
	}

	return 0
}

func calculateByLeftNodes(root *TreeNode) int {
	if root.Left != nil {
		return 1 + calculateByLeftNodes(root.Left)
	}
	return 1
}
