package implementqueue

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.stackIn = append(q.stackIn, x)
}

func (q *MyQueue) Pop() int {
	q.Peek()
	var result int
	result, q.stackOut = q.stackOut[len(q.stackOut)-1], q.stackOut[:len(q.stackOut)-1]
	return result
}

func (q *MyQueue) Peek() int {
	if len(q.stackOut) == 0 {
		for len(q.stackIn) > 0 {
			var x int
			x, q.stackIn = q.stackIn[len(q.stackIn)-1], q.stackIn[:len(q.stackIn)-1]
			q.stackOut = append(q.stackOut, x)
		}
	}
	return q.stackOut[len(q.stackOut)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.stackIn) == 0 && len(q.stackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
