// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package best_time_to_buy_and_sell_stock

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt
	profit := 0

	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}
	return profit
}
