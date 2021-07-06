package open_the_lock

import (
	"fmt"
	"strconv"
)

type Node struct {
	value     string
	timestamp int
}

var (
	set map[string]struct{}
	all map[string]struct{}
)

func openLock(deadends []string, target string) int {
	var queue []Node
	queue = append(queue, Node{"0000", 0})
	all = make(map[string]struct{})
	all["0000"] = struct{}{}
	makeDeadEnds(deadends)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.value == target {
			return node.timestamp
		} else if isDeadEnd(node.value) {
			continue
		} else {
			queue = append(queue, getNodes(node)...)
		}
	}
	return -1
}

func getNodes(node Node) []Node {
	var nodes []Node
	timestamp := node.timestamp + 1

	for i, value := range []rune(node.value) {
		nextNodeValue := []rune(node.value)
		nextNodeValue[i] = turnDigit(value, true)
		if !isExists(string(nextNodeValue)) {
			nodes = append(nodes, Node{string(nextNodeValue), timestamp})
			all[string(nextNodeValue)] = struct{}{}
		}

		prevNodeValue := []rune(node.value)
		prevNodeValue[i] = turnDigit(value, false)
		if !isExists(string(prevNodeValue)) {
			nodes = append(nodes, Node{string(prevNodeValue), timestamp})
			all[string(prevNodeValue)] = struct{}{}
		}
	}
	return nodes
}

func turnDigit(digit rune, forward bool) rune {
	value, _ := strconv.Atoi(string(digit))

	if value == 9 && forward {
		value = 0
	} else if value == 0 && !forward {
		value = 9
	} else if forward {
		value++
	} else if !forward {
		value--
	}

	return []rune(fmt.Sprint(value))[0]
}

func makeDeadEnds(slice []string) {
	set = make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
}

func isDeadEnd(item string) bool {
	_, ok := set[item]
	return ok
}

func isExists(item string) bool {
	_, ok := all[item]
	return ok
}
