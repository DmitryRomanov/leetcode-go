// https://leetcode.com/problems/sort-an-array/description/
package quick_sort

func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	middle := len(nums) / 2
	return merge(sortArray(nums[:middle]), sortArray(nums[middle:]))
}

func merge(ldata, rdata []int) []int {
	result := make([]int, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx] < rdata[ridx]:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}

	return result
}
