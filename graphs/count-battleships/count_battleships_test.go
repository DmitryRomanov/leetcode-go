package count_battleships

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBattleshipsExample1(t *testing.T) {
	assert := assert.New(t)
	result := countBattleships([][]byte{
		{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'},
	})
	assert.Equal(2, result)
}

func TestCountBattleshipsExample2(t *testing.T) {
	assert := assert.New(t)
	result := countBattleships([][]byte{
		{'.'},
	})
	assert.Equal(0, result)
}
