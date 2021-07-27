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

func TestCanFinishExample3(t *testing.T) {
	assert := assert.New(t)
	result := canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}})
	assert.Equal(true, result)
}

func TestCanFinishExample4(t *testing.T) {
	assert := assert.New(t)
	result := canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}})
	assert.Equal(true, result)
}
