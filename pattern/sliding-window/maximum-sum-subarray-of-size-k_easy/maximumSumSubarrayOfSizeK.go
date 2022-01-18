package maximumsumsubarrayofsizekeasy

func MaximumSumSubarrayOfSizeK(k int, nums []int) int {
	if len(nums) <= k || len(nums) <= 1 {
		return Sum(nums)
	}

	windowStart := 0
	maxSum, windowSum := 0, 0

	for windowEnd := range nums {
		windowSum += nums[windowEnd]
		if windowEnd >= k-1 {
			maxSum = Max(maxSum, windowSum)
			windowSum -= nums[windowStart]
			windowStart += 1
		}
	}

	return maxSum
}

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
