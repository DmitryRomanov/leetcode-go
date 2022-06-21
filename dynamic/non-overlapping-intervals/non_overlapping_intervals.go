package non_overlapping_intervals

func eraseOverlapIntervals(intervals [][]int) int {
	k := 0
	for i := 0; i < len(intervals); i++ {
		start1 := intervals[i][0]
		end1 := intervals[i][1]
		for j := i + 1; j < len(intervals); j++ {
			start2 := intervals[j][0]
			end2 := intervals[j][1]

			if start1 > end2 || start2 < end1 {
				k++
			}
		}
	}
	return k
}
