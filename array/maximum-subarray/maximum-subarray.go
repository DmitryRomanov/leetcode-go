// https://leetcode.com/problems/maximum-subarray/description/
package maximum_subarray

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentMax := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentMax < 0 {
			currentMax = 0
		}
		currentMax += nums[i]
		if currentMax > maxSum {
			maxSum = currentMax
		}
	}
	return maxSum
}
