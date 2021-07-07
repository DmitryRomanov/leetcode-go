package pathsum

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	childrenValue := targetSum - root.Val

	if root.Left != nil || root.Right != nil {
		if root.Left != nil {
			leftSum := hasPathSum(root.Left, childrenValue)
			if leftSum {
				return true
			}
		}
		if root.Right != nil {
			rightSum := hasPathSum(root.Right, childrenValue)
			if rightSum {
				return true
			}
		}
	} else {
		if root.Val == targetSum {
			return true
		} else {
			return false
		}
	}

	return false
}
