package course_schedule

type OperationType string

const (
	VISITED OperationType = "Visited"
	STUDIED OperationType = "Studied"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	visited := make(map[int]OperationType)

	for i := range prerequisites {
		graph[prerequisites[i][0]] = append(graph[prerequisites[i][0]], prerequisites[i][1])
	}

	var hasCycle func(i int) bool

	hasCycle = func(i int) bool {
		_, exists := graph[i]
		if !exists {
			return false
		}

		if visited[i] == VISITED {
			return true
		}

		if visited[i] == STUDIED {
			return false
		}

		visited[i] = VISITED

		for j := range graph[i] {
			if hasCycle(graph[i][j]) {
				return true
			}
		}

		visited[i] = STUDIED

		return false
	}

	for i := 0; i < numCourses; i++ {
		if hasCycle(i) {
			return false
		}
	}

	return true
}
