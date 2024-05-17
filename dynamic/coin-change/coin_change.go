package coin_change

// []int{1, 2, 5}, 11
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}

	for i := 1; i < amount+1; i++ {
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
