package binary_search

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		median := int(float64((left + right) / 2))
		median_value := nums[median]
		if median_value == target {
			return median
		} else if median_value < target {
			left = median + 1
		} else if median_value > target {
			right = median - 1
		}
	}

	return -1
}
