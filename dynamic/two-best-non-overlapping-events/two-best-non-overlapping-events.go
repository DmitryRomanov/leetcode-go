package two_best_non_overlapping_events

import (
	"sort"
)

type Event struct {
	startTime int
	endTime   int
	value     int
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

	for i, event := range events {
		idxOfClosestEvent := findClosestNonOverlapping(events, event.endTime, i, len(events)-1)
		if idxOfClosestEvent >= len(events) {
			tmpMax := event.value
			if tmpMax > max {
				max = tmpMax
			}
		} else {
			restEvents := make([]Event, len(events)-idxOfClosestEvent)
			copy(restEvents, events[idxOfClosestEvent:])
			sort.Slice(restEvents, func(i, j int) bool {
				return restEvents[i].value > restEvents[j].value
			})
			tmpMax := event.value + restEvents[0].value
			if tmpMax > max {
				max = tmpMax
			}
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
