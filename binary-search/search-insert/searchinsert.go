package searchinsert

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		median := left + (right - left)
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
