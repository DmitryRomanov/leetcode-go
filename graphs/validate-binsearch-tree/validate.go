package validate_binsearch_tree

import "math"

type MinMaxLimits struct {
	Min float64
	Max float64
}

func isValidBST(root *TreeNode) bool {
	limits := &MinMaxLimits{
		Min: math.Inf(-1),
		Max: math.Inf(1),
	}
	return isValidSubBST(root, limits)
}

func isValidSubBST(root *TreeNode, limits *MinMaxLimits) bool {
	if root == nil {
		return true
	}

	leftLimit := &MinMaxLimits{Min: limits.Min, Max: float64(root.Val)}
	leftIsValid := isValidSubBST(root.Left, leftLimit)

	rightLimit := &MinMaxLimits{Min: float64(root.Val), Max: limits.Max}
	rightIsValid := isValidSubBST(root.Right, rightLimit)

	return leftIsValid && rightIsValid && limits.Min < float64(root.Val) && float64(root.Val) < limits.Max
}
