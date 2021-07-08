package validate_binsearch_tree

func isValidBST(root *TreeNode) bool {
	if root.Left != nil {
		if root.Val > root.Left.Val {
			leftIsValidBST := isValidBST(root.Left)
			if !leftIsValidBST {
				return false
			}
		} else {
			return false
		}
	}

	if root.Right != nil {
		if root.Val < root.Right.Val {
			rightIsValidBST := isValidBST(root.Right)
			if !rightIsValidBST {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
