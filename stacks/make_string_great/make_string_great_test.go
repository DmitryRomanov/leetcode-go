package make_string_great

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeGoodExample1(t *testing.T) {
	assert := assert.New(t)
	result := makeGood("leEeetcode")
	assert.Equal("leetcode", result)
}

func TestMakeGoodExample2(t *testing.T) {
	assert := assert.New(t)
	result := makeGood("abBAcC")
	assert.Equal("", result)
}

func TestMakeGoodExample3(t *testing.T) {
	assert := assert.New(t)
	result := makeGood("s")
	assert.Equal("s", result)
}
