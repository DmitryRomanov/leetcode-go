package searchinsert

import "math"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		median := int(math.Round(float64((left + right) / 2)))
		if nums[median] == target {
			return median
		} else if nums[median] < target {
			left = median + 1
		} else if nums[median] > target {
			right = median - 1
		}
	}

	return left
}
