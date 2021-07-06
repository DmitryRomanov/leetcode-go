package open_the_lock

import (
	"fmt"
	"strconv"
)

type Node struct {
	value     string
	timestamp int
}

func openLock(deadends []string, target string) int {
	var queue []Node
	queue = append(queue, Node{"0000", 0})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.value == target {
			return node.timestamp
		} else if contains(deadends, node.value) {
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
	digits := []rune(node.value)
	for i := range digits {
		nextNodeValue := digits
		nextNodeValue[i] = turnDigit(nextNodeValue[i], true)
		nodes = append(nodes, Node{string(nextNodeValue), timestamp})

		prevNodeValue := digits
		prevNodeValue[i] = turnDigit(prevNodeValue[i], false)
		nodes = append(nodes, Node{string(prevNodeValue), timestamp})
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

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
