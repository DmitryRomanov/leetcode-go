package countnodes

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var lh, rh = 0, 0

	for current := root; current != nil; current = current.Left {
		lh += 1
	}

	for current := root; current != nil; current = current.Right {
		rh += 1
	}

	if lh == rh {
		return int(math.Pow(float64(2), float64(lh))) - 1
	}

	lc := countNodes(root.Left)
	rc := countNodes(root.Right)

	return lc + rc + 1
}
