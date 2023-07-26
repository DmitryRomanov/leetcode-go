package longest_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := longestPalindrome("abccccdd")
	assert.Equal(7, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := longestPalindrome("a")
	assert.Equal(1, result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := longestPalindrome("aab")
	assert.Equal(3, result)
}

func TestExample4(t *testing.T) {
	assert := assert.New(t)
	result := longestPalindrome("aaa")
	assert.Equal(3, result)
}

func TestExample5(t *testing.T) {
	assert := assert.New(t)
	result := longestPalindrome("aaaAaaaa")
	assert.Equal(7, result)
}
