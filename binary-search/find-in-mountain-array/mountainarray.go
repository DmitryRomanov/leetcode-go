// https://leetcode.com/problems/find-in-mountain-array/
package mountainarray

func findInMountainArray(target int, mountainArr *MountainArray) int {
	peak := getPeak(mountainArr)
	var left, right int
	if peak != -1 {
		val := mountainArr.get(peak)
		if val == target {
			return peak
		} else if val > target {
			right = peak
			left = 0
		} else if val < target {
			left = peak
			right = mountainArr.length()
		}

		for left < right {
			median := left + (right-left)>>1
			medianValue := mountainArr.get(median)
			if medianValue == target {
				return median
			} else if medianValue > target {
				right = median - 1
			} else if medianValue < target {
				left = median + 1
			}
		}
	}
	return -1
}

func getPeak(mountainArr *MountainArray) int {
	left := 0
	right := mountainArr.length() - 1
	for left < right {
		median := left + (right-left)>>1
		medianValue := mountainArr.get(median)
		beforeValue := mountainArr.get(median - 1)
		afterValue := mountainArr.get(median + 1)
		if beforeValue < medianValue && medianValue > afterValue {
			return median
		} else if beforeValue < medianValue && medianValue < afterValue {
			return median
		} else if beforeValue < medianValue && afterValue < medianValue {
			left = median + 1
		} else if beforeValue > medianValue && afterValue > medianValue {
			right = median - 1
		}
	}

	return -1
}

type MountainArray struct {
	arr []int
}

func NewMountainArray(arr []int) *MountainArray {
	return &MountainArray{arr}
}

func (a *MountainArray) get(index int) int {
	return a.arr[index]
}
func (a *MountainArray) length() int {
	return len(a.arr)
}
