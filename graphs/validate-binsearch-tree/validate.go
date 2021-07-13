package validate_binsearch_tree

import "math"

type MinMaxLimits struct {
	Min int
	Max int
}

func isValidBST(root *TreeNode) bool {
	return isValidSubBST(
		root,
		MinMaxLimits{
			Min: math.MinInt32,
			Max: math.MaxInt32,
		})
}

func isValidSubBST(root *TreeNode, limits MinMaxLimits) bool {
	if root == nil {
		return true
	}

	leftIsValid := isValidSubBST(root.Left, MinMaxLimits{Min: limits.Min, Max: root.Val})
	rightIsValid := isValidSubBST(root.Right, MinMaxLimits{Min: root.Val, Max: limits.Min})

	return leftIsValid && rightIsValid && limits.Min > root.Val && root.Val < limits.Max
}
