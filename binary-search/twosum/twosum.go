package twosum

import "math"

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		secondIndex := binarySearch(numbers, target-numbers[i], i)
		if secondIndex != -1 {
			return []int{i + 1, secondIndex + 1}
		}
	}
	return nil
}

func binarySearch(numbers []int, target, from int) int {
	left := from
	right := len(numbers)

	for (right - left) > 1 {
		medianIndex := int(math.Floor(float64((left + right) / 2)))
		medianValue := numbers[medianIndex]

		if medianValue == target {
			return medianIndex
		}

		if medianValue < target {
			left = medianIndex
		} else if medianValue > target {
			right = medianIndex
		}
	}

	return -1
}
