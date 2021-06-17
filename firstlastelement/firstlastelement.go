package firstlastelement

import "math"

func searchRange(nums []int, target int) []int {
	oneValue := findOne(nums, target)
	firstValue := oneValue
	lastValue := oneValue

	firstNotEqualValue := firstNotEqual(nums[0:oneValue], target)
	if firstNotEqualValue >= 0 {
		firstValue = firstNotEqualValue + 1
	}

	lastNotEqualValue := lastNotEqual(nums[oneValue+1:], target)
	if lastNotEqualValue >= 0 {
		lastValue = oneValue + lastNotEqualValue
	}

	return []int{
		firstValue,
		lastValue,
	}
}

func findOne(nums []int, target int) int {
	left := 0
	right := len(nums)

	for (right - left) > 1 {
		median := int(math.Floor(float64((right - left) / 2)))
		medianValue := nums[median]
		if medianValue == target {
			return median
		}
		if medianValue < target {
			left = median
		} else if medianValue > target {
			right = median
		}
	}

	return -1
}

func firstNotEqual(nums []int, target int) int {
	return -1
}

func lastNotEqual(nums []int, target int) int {
	return -1
}
