package maximumsubarray

import "math"

// Method 1
// Sliding Window
// Time: O(N), Space: O(1)
func MaximumSubarray(nums []int) int {
	maxSub := nums[0]
	currentSum := 0
	for _, v := range nums {
		// 只要當前的總和為負值，就表示不值得被我們考慮，因此要歸零
		// 也就是捨棄當前 index 前面所有的 subarray
		if currentSum < 0 {
			currentSum = 0
		}
		currentSum += v
		maxSub = Max(maxSub, currentSum)
	}
	return maxSub
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// ================= Method 2,3 為 Kadane 演算法 (DP 的一種) =================
// wiki:
// 		https://zh.wikipedia.org/zh-tw/%E6%9C%80%E5%A4%A7%E5%AD%90%E6%95%B0%E5%88%97%E9%97%AE%E9%A2%98

// Method 2
// DP
// Time: O(N), Space: O(N)
func MaximumSubarray2(nums []int) int {
	// 宣告一長度為 nums 個數的 DP slice
	dp := make([]int, len(nums))
	// base case
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = Max(nums[i], nums[i]+dp[i-1])
	}

	maxSub := math.MinInt
	for _, v := range dp {
		maxSub = Max(maxSub, v)
	}
	return maxSub
}

// Method 3
// DP，不使用 array
// Time: O(N), Space: O(1)
func MaximumSubarray3(nums []int) int {
	// base case
	dp := nums[0]
	maxSub := nums[0]

	for i := 1; i < len(nums); i++ {
		dp = Max(nums[i], nums[i]+dp)
		maxSub = Max(maxSub, dp)
	}

	return maxSub
}

// Method 4
// Prefix sum 解法
// https://leetcode.com/problems/maximum-subarray/discuss/2072037/4-different-C%2B%2B-Solutions-or-Additional-Java-and-Python-solution

// Method 5
// Divide and Conquer 解法
// https://leetcode.com/problems/maximum-subarray/discuss/2072037/4-different-C%2B%2B-Solutions-or-Additional-Java-and-Python-solution
