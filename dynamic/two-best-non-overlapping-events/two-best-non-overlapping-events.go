package two_best_non_overlapping_events

import "sort"

type event struct {
	startTime int
	endTime   int
	value     int
}

func maxTwoEvents(events_array [][]int) int {
	sort.Slice(events_array, func(i, j int) bool {
		return events_array[i][0] < events_array[j][0]
	})

	max := 0
	events := make([]event, len(events_array))
	for i := 0; i < len(events_array); i++ {
		events[i] = event{
			startTime: events_array[i][0],
			endTime:   events_array[i][1],
			value:     events_array[i][2],
		}
	}

	for i, event := range events {
		idxOfClosestEvent := findClosestNonOverlapping(events, event.endTime, event.endTime, len(events)-1)
		tmpMax := events[i].value + events[idxOfClosestEvent].value
		if tmpMax > max {
			max = tmpMax
		}
	}

	return max
}

func findClosestNonOverlapping(events []event, end, left, right int) int {
	if right > left {
		return left
	}
	median := (left + right) / 2

	if events[median].endTime > end {
		return findClosestNonOverlapping(events, end, median+1, right)
	} else if events[median].endTime < end {
		return findClosestNonOverlapping(events, end, left, median-1)
	} else {
		return median
	}
}
