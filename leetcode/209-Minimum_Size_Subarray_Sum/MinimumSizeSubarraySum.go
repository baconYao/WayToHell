package minimumsizesubarraysum

import "math"

// Method 1
// Sliding Window
// Time: O(n), Spatial: O(1)
func MinimumSizeSubarraySum(target int, nums []int) int {
	window_left := 0
	accumulate := 0
	minLength := math.MaxInt
	for window_right := 0; window_right < len(nums); window_right++ {
		accumulate += nums[window_right]
		for accumulate >= target {
			minLength = Min(minLength, window_right-window_left+1)
			accumulate -= nums[window_left]
			window_left += 1
		}
	}

	if minLength != math.MaxInt {
		return minLength
	}
	return 0
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
