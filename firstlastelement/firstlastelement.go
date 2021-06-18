package firstlastelement

import "math"

func searchRange(nums []int, target int) []int {
	firstResult := -1
	lastResult := -1

	firstEqual := firstBorderEqual(nums, target, LEFT)
	if firstEqual >= 0 {
		firstResult = firstEqual
	}

	lastEqual := firstBorderEqual(nums, target, RIGHT)
	if lastEqual >= 0 {
		lastResult = lastEqual
	}

	return []int{
		firstResult,
		lastResult,
	}
}

type Direction string

const (
	RIGHT Direction = "RIGHT"
	LEFT  Direction = "LEFT"
)

func firstBorderEqual(nums []int, target int, direction Direction) int {
	left := 0
	right := len(nums) - 1
	result := -1

	for left <= right {
		median := int(math.Floor(float64((right + left) / 2)))
		medianValue := nums[median]

		if medianValue == target {
			result = median

			if direction == LEFT {
				right = median - 1
			} else if direction == RIGHT {
				left = median + 1
			}
		}
		if medianValue < target {
			left = median + 1
		} else if medianValue > target {
			right = median - 1
		}
	}

	return result
}
