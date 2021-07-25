package remove_k_digits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveKdigitsExample1(t *testing.T) {
	assert := assert.New(t)
	result := removeKdigits("1432219", 3)
	assert.Equal("1219", result)
}

func TestRemoveKdigitsExample2(t *testing.T) {
	assert := assert.New(t)
	result := removeKdigits("10200", 1)
	assert.Equal("200", result)
}

func TestRemoveKdigitsExample3(t *testing.T) {
	assert := assert.New(t)
	result := removeKdigits("10", 2)
	assert.Equal("0", result)
}

func TestRemoveKdigitsExample4(t *testing.T) {
	assert := assert.New(t)
	result := removeKdigits("9", 1)
	assert.Equal("0", result)
}

func TestRemoveKdigitsExample5(t *testing.T) {
	assert := assert.New(t)
	result := removeKdigits("112", 1)
	assert.Equal("11", result)
}

func TestRemoveKdigitsExample6(t *testing.T) {
	assert := assert.New(t)
	result := removeKdigits("112", 2)
	assert.Equal("1", result)
}

func TestRemoveKdigitsExample7(t *testing.T) {
	assert := assert.New(t)
	result := removeKdigits("10001", 4)
	assert.Equal("0", result)
}
