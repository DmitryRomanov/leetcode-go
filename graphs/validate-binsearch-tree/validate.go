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
	if root.Left != nil {
		if root.Val > root.Left.Val && root.Val < limits.Max && root.Val > limits.Min {
			leftIsValidBST := isValidSubBST(root.Left, MinMaxLimits{Min: Min(root.Val, limits.Min), Max: Max(limits.Max, root.Val)})
			if !leftIsValidBST {
				return false
			}
		} else {
			return false
		}
	}

	if root.Right != nil {
		if root.Val < root.Right.Val && root.Val > limits.Max && root.Val < limits.Min {
			rightIsValidBST := isValidSubBST(root.Right, MinMaxLimits{Min: Min(root.Val, limits.Min), Max: Max(limits.Max, root.Val)})
			if !rightIsValidBST {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
