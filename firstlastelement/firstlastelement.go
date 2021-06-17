package firstlastelement

import "math"

func searchRange(nums []int, target int) []int {
	oneValue := findOne(nums, target)
	if oneValue == -1 {
		return []int{-1, -1}
	}

	firstResult := oneValue
	lastResult := oneValue

	firstEqual := firstEqual(nums[0:oneValue], target)
	if firstEqual >= 0 {
		firstResult = firstEqual
	}

	lastEqual := lastEqual(nums[oneValue:], target)
	if lastEqual >= 0 {
		lastResult = oneValue + lastEqual
	}

	return []int{
		firstResult,
		lastResult,
	}
}

func findOne(nums []int, target int) int {
	left := 0
	right := len(nums)

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

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

func firstEqual(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	result := -1

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	for (right - left) > 1 {
		median := int(math.Floor(float64((right - left) / 2)))
		medianValue := nums[median]

		if medianValue == target {
			result = median
			right = median
		}
		if medianValue < target {
			right = median
		} else if medianValue > target {
			left = median
		}
	}

	return result
}

func lastEqual(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	result := -1

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	for (right - left) > 1 {
		median := int(math.Floor(float64((right - left) / 2)))
		medianValue := nums[median]

		if medianValue == target {
			result = median
			left = median + 1
		}
		if medianValue < target {
			left = median
		} else if medianValue > target {
			right = median
		}
	}

	return result
}
