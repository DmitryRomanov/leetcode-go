package two_best_non_overlapping_events

import "sort"

type event struct {
	startTime int
	endTime   int
	value     int
}

func maxTwoEvents(events_array [][]int) int {
	sort.Slice(events_array, func(i, j int) bool {
		return events_array[i][1] < events_array[j][1]
	})

	max := 0
	events := make([]event, len(events_array))
	for i := 0; i < len(events_array); i++ {
		tmpEvent := event{
			startTime: events_array[i][0],
			endTime:   events_array[i][1],
			value:     events_array[i][2],
		}
		events = append(events, tmpEvent)
	}

	for left := 0; left < len(events); left++ {
		if events[left].value > max {
			max = events[left].value
		}
		for right := left + 1; right < len(events); right++ {
			leftEvent := events[left]
			rightEvent := events[right]
			if leftEvent.endTime < rightEvent.startTime {
				tmpMax := leftEvent.value + rightEvent.value
				if tmpMax > max {
					max = tmpMax
				}
			}
		}
	}

	return max
}
