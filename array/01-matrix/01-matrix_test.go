package matrix_01

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})
	assert.Equal([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	})
	assert.Equal([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 2, 1},
	}, result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := updateMatrix([][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 0, 0},
	})
	assert.Equal([][]int{
		{1, 2, 1},
		{0, 1, 0},
		{0, 0, 0},
	}, result)
}

func TestExample4(t *testing.T) {
	assert := assert.New(t)
	result := updateMatrix([][]int{
		{1, 1},
		{1, 0},
	})
	assert.Equal([][]int{
		{2, 1},
		{1, 0},
	}, result)
}

func TestExample5(t *testing.T) {
	assert := assert.New(t)
	result := updateMatrix([][]int{
		{1, 1, 0},
		{1, 0, 0},
		{0, 0, 0},
	})
	assert.Equal([][]int{
		{2, 1, 0},
		{1, 0, 0},
		{0, 0, 0},
	}, result)
}

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(1 * time.Second)
		updateMatrix([][]int{
			{1, 1},
			{1, 0},
		})
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(3 * time.Second)

		updateMatrix([][]int{
			{1, 1, 0},
			{1, 0, 0},
			{0, 0, 0},
		})
	}
}
