package mountainarray

import "math"

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		median := int(math.Floor(float64((right + left) / 2)))

		if arr[median-1] < arr[median] && arr[median] > arr[median+1] {
			return median
		} else if arr[median-1] < arr[median] && arr[median] < arr[median+1] {
			left = median
		} else if arr[median-1] > arr[median] && arr[median] > arr[median+1] {
			right = median
		}
	}

	return -1
}
