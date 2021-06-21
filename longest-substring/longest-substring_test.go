package longestsubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstringExample1(t *testing.T) {
	assert := assert.New(t)
	result := lengthOfLongestSubstring("abcabcbb")
	assert.Equal(3, result)
}

func TestLengthOfLongestSubstringExample2(t *testing.T) {
	assert := assert.New(t)
	result := lengthOfLongestSubstring("bbbbb")
	assert.Equal(1, result)
}

func TestLengthOfLongestSubstringExample3(t *testing.T) {
	assert := assert.New(t)
	result := lengthOfLongestSubstring("pwwkew")
	assert.Equal(3, result)
}

func TestLengthOfLongestSubstringExample4(t *testing.T) {
	assert := assert.New(t)
	result := lengthOfLongestSubstring("")
	assert.Equal(0, result)
}
