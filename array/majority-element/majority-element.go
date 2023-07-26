// https://leetcode.com/problems/majority-element/description/
package majority_element

func majorityElement(nums []int) int {
	target := nums[0]
	count := 0
	for _, v := range nums {
		if v == target {
			count++
		} else {
			count--
			if count == 0 {
				target = v
				count = 1
			}
		}
	}
	return target
}
