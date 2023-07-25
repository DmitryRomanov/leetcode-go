package ransom_note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := canConstruct("a", "b")
	assert.False(result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := canConstruct("aa", "ab")
	assert.False(result)
}

func TestExample3(t *testing.T) {
	assert := assert.New(t)
	result := canConstruct("aa", "aab")
	assert.True(result)
}
