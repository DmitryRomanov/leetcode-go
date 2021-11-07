package subsets_2

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{{}}
	left := 0
	right := len(nums)

	for right > left {
		subNums := nums[left:right]

		for i := range subNums {
			subSubNums := subNums[0 : i+1]
			if !hasDup(result, subSubNums) {
				result = append(result, subSubNums)
			}
		}
		left++
	}

	return result
}

func hasDup(result [][]int, nums []int) bool {
	for _, resultNum := range result {
		if isEqual(resultNum, nums) {
			return true
		}
	}
	return false
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
