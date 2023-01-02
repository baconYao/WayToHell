package coinchange

import "math"

// method 1, DP DFS, Top-Down Memo
func coinChange(coins []int, amount int) int {
	// 這裡的 amount 做 + 1 是為了後續的 index 時使用。ex: 22, 34, 36 行。
	memo := make([]int, amount+1)

	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	res := getMinCoins(coins, memo, amount)

	if res >= math.MaxInt32 {
		return -1
	}

	return res
}

func getMinCoins(coins []int, memo []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return math.MaxInt32
	}
	if memo[amount] != -1 {
		return memo[amount]
	}

	minCoins := math.MaxInt32

	for _, c := range coins {
		minCoins = min(minCoins, 1+getMinCoins(coins, memo, amount-c))
	}

	memo[amount] = minCoins

	return minCoins
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// method 2, DP, Bottom Up
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c >= 0 {
				dp = min(dp[i], 1+dp[i-c])
			}
		}
	}

	if dp[amount] >= math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
