package course_schedule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Explanation: There are a total of 2 courses to take.
//To take course 1 you should have finished course 0. So it is possible.
func TestCanFinishExample1(t *testing.T) {
	assert := assert.New(t)
	result := canFinish(2, [][]int{{1, 0}})
	assert.Equal(true, result)
}

//Explanation: There are a total of 2 courses to take.
//To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
func TestCanFinishExample2(t *testing.T) {
	assert := assert.New(t)
	result := canFinish(2, [][]int{{1, 0}, {0, 1}})
	assert.Equal(false, result)
}
