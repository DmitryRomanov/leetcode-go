package coin_change

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := coinChange([]int{1, 2, 5}, 11)
	assert.Equal(3, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := coinChange([]int{2}, 3)
	assert.Equal(-1, result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := coinChange([]int{1}, 0)
	assert.Equal(0, result)
}

// BenchmarkTest-8   	   40114	    118566 ns/op	  164239 B/op	       1 allocs/op
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coinChange([]int{1, 2, 5}, i)
	}
}
