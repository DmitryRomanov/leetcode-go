package searchinsert

import "math"

func searchInsert(nums []int, target int) int {
	currentTarget := target
	for currentTarget > 0 {
		foundPosition := binarySearch(nums, currentTarget)
		if foundPosition != -1 {
			if currentTarget == target {
				return foundPosition
			} else {
				return foundPosition + 1
			}
		} else {
			currentTarget--
		}
	}

	return 0
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		median := int(math.Floor(float64((left + right) / 2)))
		if nums[median] < target {
			left = median + 1
		} else if nums[median] > target {
			right = median - 1
		} else {
			return median
		}
	}

	return -1
}
