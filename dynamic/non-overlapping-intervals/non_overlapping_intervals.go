package non_overlapping_intervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	k := 0

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	var ends []int
	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		add := true
		for j := range ends {
			if start < ends[j] {
				add = false
				k++
				break
			}
		}
		if add {
			ends = append(ends, intervals[i][1])
		}

	}
	return k
}
