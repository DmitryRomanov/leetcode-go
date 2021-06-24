package allanagrams

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAnagramsExample1(t *testing.T) {
	assert := assert.New(t)
	result := findAnagrams("cbaebabacd", "abc")
	assert.Equal([]int{0, 6}, result)
}

func TestFindAnagramsExample2(t *testing.T) {
	assert := assert.New(t)
	result := findAnagrams("abab", "ab")
	assert.Equal([]int{0, 1, 2}, result)
}
