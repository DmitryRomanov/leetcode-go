package two_best_non_overlapping_events

import "sort"

type event struct {
	startTime int
	endTime   int
	value     int
}

type EventsSortable []event

func (a EventsSortable) Len() int           { return len(a) }
func (a EventsSortable) Less(i, j int) bool { return a[i].startTime < a[j].startTime }
func (a EventsSortable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func maxTwoEvents(events_array [][]int) int {
	max := 0
	events := make([]event, len(events_array))
	for i := 0; i < len(events_array); i++ {
		events[i] = event{
			startTime: events_array[i][0],
			endTime:   events_array[i][1],
			value:     events_array[i][2],
		}
	}

	sort.Sort(EventsSortable(events))

	for i, event := range events {
		idxOfClosestEvent := findClosestNonOverlapping(events, event.endTime, i, len(events)-1)
		if idxOfClosestEvent >= len(events) {
			tmpMax := event.value
			if tmpMax > max {
				max = tmpMax
			}
		} else {
			for j := idxOfClosestEvent; j < len(events); j++ {
				tmpMax := event.value + events[j].value
				if tmpMax > max {
					max = tmpMax
				}
			}
		}
	}

	return max
}

func findClosestNonOverlapping(events []event, end, left, right int) int {
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
