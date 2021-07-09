package validate_binsearch_tree

func isValidBST(root *TreeNode) bool {
	if root.Left != nil {
		if root.Left.Right != nil && root.Left.Right.Val > root.Val {
			return false
		}

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
			if root.Right.Left != nil && root.Right.Left.Val < root.Val {
				return false
			}

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
