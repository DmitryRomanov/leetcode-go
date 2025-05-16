package implementqueue

import (
	"fmt"
	"testing"
)

func TestExample0(t *testing.T) {
	myQueue := MyQueue{}
	myQueue.Push(1) // queue is: [1]
	myQueue.Push(2) // queue is: [1, 2] (leftmost is front of the queue)
	myQueue.Peek()  // return 1
	myQueue.Pop()   // return 1, queue is [2]
	myQueue.Empty() // return false

	fmt.Println(myQueue)

}
