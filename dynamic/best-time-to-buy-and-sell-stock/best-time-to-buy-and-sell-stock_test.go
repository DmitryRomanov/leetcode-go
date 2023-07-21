package best_time_to_buy_and_sell_stock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert := assert.New(t)
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	assert.Equal(5, result)
}

func TestExample2(t *testing.T) {
	assert := assert.New(t)
	result := maxProfit([]int{7, 6, 4, 3, 1})
	assert.Equal(0, result)
}
