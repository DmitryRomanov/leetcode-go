package countnodes

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
