package climbingstairs

// 這題是一個略微變形的費氏數列

// method 1
// Time: O(N), space: O(n)
func ClimbingStairs(n int) int {
	if n <= 2 {
		return n
	}
	// dp 用來記錄到達 n 層階梯時的所有走法總和
	dp := []int{}
	// base case
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// method 2
// Time: O(N), space: O(1)
func ClimbingStairs2(n int) int {
	if n <= 2 {
		return n
	}
	// dp 用來記錄到達 n 層階梯時的所有走法總和
	var dp_n int
	// base case
	dp_n_diff_2, dp_n_diff_1 := 1, 2
	for i := 3; i <= n; i++ {
		dp_n = dp_n_diff_1 + dp_n_diff_2
		dp_n_diff_2 = dp_n_diff_1
		dp_n_diff_1 = dp_n
	}
	return dp_n
}

// method 3, recursive
// 但在Leetcode 會超時
func ClimbingStairs3(n int) int {
	if n <= 2 {
		return n
	}
	return ClimbingStairs3(n-1) + ClimbingStairs3(n-2)
}
