package valid_parentless

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidExample1(t *testing.T) {
	assert := assert.New(t)
	result := isValid("()")
	assert.Equal(true, result)
}

func TestIsValidExample2(t *testing.T) {
	assert := assert.New(t)
	result := isValid("()[]{}")
	assert.Equal(true, result)
}

func TestIsValidExample3(t *testing.T) {
	assert := assert.New(t)
	result := isValid("(]")
	assert.Equal(false, result)
}

func TestIsValidExample4(t *testing.T) {
	assert := assert.New(t)
	result := isValid("([)]")
	assert.Equal(false, result)
}

func TestIsValidExample5(t *testing.T) {
	assert := assert.New(t)
	result := isValid("{[]}")
	assert.Equal(true, result)
}

func TestIsValidExample6(t *testing.T) {
	assert := assert.New(t)
	result := isValid("{[[]]}")
	assert.Equal(true, result)
}

func TestIsValidExample7(t *testing.T) {
	assert := assert.New(t)
	result := isValid("{[][]}")
	assert.Equal(true, result)
}
