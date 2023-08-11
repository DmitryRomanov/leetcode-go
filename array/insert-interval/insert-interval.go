// https://leetcode.com/problems/insert-interval/
package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	mergedStart, mergedFinish := newInterval[0], newInterval[1]

	left := [][]int{}
	right := [][]int{}
	for _, v := range intervals {
		if mergedStart > v[1] {
			left = append(left, v)
		} else if mergedFinish < v[0] {
			right = append(right, v)
		} else {
			mergedStart = min(v[0], mergedStart)
			mergedFinish = max(v[1], mergedFinish)
		}
	}
	return append(append(left, []int{mergedStart, mergedFinish}), right...)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
