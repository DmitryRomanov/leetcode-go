package coin_change

func coinChange(coins []int, amount int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if i-c >= 0 {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
