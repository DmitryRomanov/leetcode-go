package guessnumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGuessNumberExample1(t *testing.T) {
	assert := assert.New(t)
	pick = 6
	result := guessNumber(10)
	assert.Equal(pick, result)
}

func TestGuessNumberExample2(t *testing.T) {
	assert := assert.New(t)
	pick = 1
	result := guessNumber(1)
	assert.Equal(pick, result)
}

func TestGuessNumberExample3(t *testing.T) {
	assert := assert.New(t)
	pick = 1
	result := guessNumber(2)
	assert.Equal(pick, result)
}
