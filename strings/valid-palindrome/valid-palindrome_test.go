package valid_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := isPalindrome("A man, a plan, a canal: Panama")
	assert.True(result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := isPalindrome("race a car")
	assert.False(result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := isPalindrome(" ")
	assert.True(result)
}

func TestExample4(t *testing.T) {
	assert := assert.New(t)
	result := isPalindrome("racacar")
	assert.True(result)
}

func TestExample5(t *testing.T) {
	assert := assert.New(t)
	result := isPalindrome("0P")
	assert.False(result)
}
