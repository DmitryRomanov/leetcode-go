package course_schedule_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Explanation: There are a total of 2 courses to take.
// To take course 1 you should have finished course 0. So the correct course order is [0,1].
func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := findOrder(2, [][]int{{1, 0}})
	assert.Equal([]int{0, 1}, result)
}

// Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2.
// Both courses 1 and 2 should be taken after you finished course 0.
// So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	assert.Equal([]int{0, 2, 1, 3}, result)
}

// func TestExample4(t *testing.T) {
// 	assert := assert.New(t)
// 	result := findOrder(3, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
// 	assert.Equal([]int{0, 2, 1}, result)
// }

func TestExample7(t *testing.T) {
	assert := assert.New(t)
	result := findOrder(2, [][]int{{0, 1}, {1, 0}})
	assert.Equal([]int{}, result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := findOrder(1, [][]int{})
	assert.Equal([]int{0}, result)
}

func TestExample6(t *testing.T) {
	assert := assert.New(t)
	result := findOrder(2, [][]int{})
	assert.Equal([]int{1, 0}, result)
}

func TestExample5(t *testing.T) {
	assert := assert.New(t)
	result := findOrder(4, [][]int{{3, 0}, {0, 1}})
	assert.Equal([]int{2, 1, 0, 3}, result)
}
