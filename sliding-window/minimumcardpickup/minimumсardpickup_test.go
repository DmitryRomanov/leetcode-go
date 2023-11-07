package minimumcardpickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := minimumCardPickup([]int{3, 4, 2, 3, 4, 7})
	assert.Equal(4, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := minimumCardPickup([]int{1, 0, 5, 3})
	assert.Equal(-1, result)
}

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minimumCardPickup([]int{3, 4, 2, 3, 4, 7})
	}
}
