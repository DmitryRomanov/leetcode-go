package minimumcardpickup

import (
	"math/rand"
	"testing"
	"time"

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

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := minimumCardPickup([]int{0, 0})
	assert.Equal(2, result)
}

func BenchmarkBruteforce(b *testing.B) {
	rand.Seed(time.Now().Unix())
	nums := 1_000
	input := make([]int, 0, nums)
	for i := 0; i < nums; i++ {
		input = append(input, randInt(1, 1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minimumCardPickupBruteforce(input)
	}
}

func BenchmarkHashmap(b *testing.B) {
	rand.Seed(time.Now().Unix())
	nums := 1_000
	input := make([]int, 0, nums)
	for i := 0; i < nums; i++ {
		input = append(input, randInt(1, 1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minimumCardPickup(input)
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
