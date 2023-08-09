package middle

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func build(in []int) *ListNode {
	if len(in) == 0 {
		return nil
	}
	sort.Sort(sort.Reverse(sort.IntSlice(in)))

	var result *ListNode
	for _, v := range in {
		result = &ListNode{
			Val:  v,
			Next: result,
		}
	}
	return result
}
