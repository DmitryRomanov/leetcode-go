// https://leetcode.com/problems/course-schedule-ii/
package course_schedule_2

import (
	"errors"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	digraph := getDigraph(prerequisites)
	sorted, err := kahn(digraph)
	if err != nil {
		return []int{}
	}
	if len(sorted) < numCourses {
		for i := numCourses - 1; i >= len(sorted)-2; i-- {
			sorted = append(sorted, i)
		}
		return sorted[:numCourses]
	} else {
		return sorted[:numCourses]
	}
}

// [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
func getDigraph(prerequisites [][]int) map[int][]int {
	result := make(map[int][]int)
	for _, v := range prerequisites {
		vertex, root := v[0], v[1]
		result[root] = append(result[root], vertex)
	}
	return result
}

// 0: [1, 2]
// 1: [3]
// 2: [3]
// 3: []
func kahn(digraph map[int][]int) ([]int, error) {
	indegrees := make(map[int]int)
	for u := range digraph {
		if digraph[u] != nil {
			for _, v := range digraph[u] {
				indegrees[v]++
			}
		}
	}

	var queue []int
	for u := range digraph {
		if _, ok := indegrees[u]; !ok {
			queue = append(queue, u)
		}
	}

	var order []int
	for len(queue) > 0 {
		u := queue[len(queue)-1]
		queue = queue[:(len(queue) - 1)]
		order = append(order, u)
		for _, v := range digraph[u] {
			indegrees[v]--
			if indegrees[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	for _, indegree := range indegrees {
		if indegree > 0 {
			return order, errors.New("not a DAG")
		}
	}
	return order, nil
}
