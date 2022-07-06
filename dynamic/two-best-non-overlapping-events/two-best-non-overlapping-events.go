package two_best_non_overlapping_events

import (
	"sort"
)

type Event struct {
	startTime int
	endTime   int
	value     int
}

func buildMax(events []Event) []int {
	maxes := make([]int, len(events))
	for i := 0; i < len(events); i++ {
		maxes[i] = events[i].value
	}

	max := maxes[len(events)-1]

	for i := len(events) - 2; i >= 0; i-- {
		if max > maxes[i] {
			maxes[i] = max
		} else {
			max = maxes[i]
		}
	}
	return maxes
}

func maxTwoEvents(events_array [][]int) int {
	max := 0
	events := make([]Event, len(events_array))
	for i := 0; i < len(events_array); i++ {
		events[i] = Event{
			startTime: events_array[i][0],
			endTime:   events_array[i][1],
			value:     events_array[i][2],
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].startTime < events[j].startTime
	})

	maxes := buildMax(events)

	for i, event := range events {
		idxOfClosestEvent := findClosestNonOverlapping(events, event.endTime, i, len(events)-1)
		tmpMax := 0
		if idxOfClosestEvent >= len(events) {
			tmpMax = event.value
		} else {
			tmpMax = event.value + maxes[idxOfClosestEvent]
		}

		if tmpMax > max {
			max = tmpMax
		}
	}

	return max
}

func findClosestNonOverlapping(events []Event, end, left, right int) int {
	if right < left {
		return left
	}
	median := (left + right) / 2

	if end < events[median].startTime {
		return findClosestNonOverlapping(events, end, left, median-1)
	} else {
		return findClosestNonOverlapping(events, end, median+1, right)
	}
}
