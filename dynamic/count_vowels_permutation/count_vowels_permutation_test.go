package count_vowels_permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := countVowelPermutation(1)
	assert.Equal(5, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := countVowelPermutation(2)
	assert.Equal(10, result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := countVowelPermutation(5)
	assert.Equal(68, result)
}

func TestExample4(t *testing.T) {
	assert := assert.New(t)
	result := countVowelPermutation(144)
	assert.Equal(18208803, result)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countVowelPermutation(20_000)
	}
}
