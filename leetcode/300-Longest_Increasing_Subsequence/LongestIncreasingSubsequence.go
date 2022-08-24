package longestincreasingsubsequence

// Method 1
// Time: O(N^2)
func LengthOfLIS(nums []int) int {
	// 宣告一長度為 nums 個數的 DP slice，並初始化值為 1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	// 根據先前已算出來的 dp 數值，推算當前的 index i 所在位置的 dp 值
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
	}

	// 找出最大的 DP 值就是答案
	maxValue := 0
	for _, v := range dp {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Method 2
// 二分法查詢 - 與紙牌遊戲 patience game 有關
// 只要分類出牌堆的數量，該數量就表示最長子序列
func LengthOfLIS2(nums []int) int {
	// 初始化牌堆數為 0
	piles := 0
	// top 是用來記錄牌堆最上面的數字大小
	// 由於可能會是一張牌就成一堆，因此最長就 len(nums)
	top := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		// 當前的撲克
		poker := nums[i]
		left, right := 0, piles
		// 搜尋左側邊界的二分搜尋法
		for left < right {
			mid := (left + right) / 2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 沒找到合適的牌堆，則新建一堆
		if left == piles {
			piles++
		}
		// 把此牌放到牌堆頂端
		top[left] = poker
	}
	return piles
}

// Method 3
// Leetcode 的解答
func LengthOfLIS3(nums []int) int {
	arr := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > arr[len(arr)-1] {
			arr = append(arr, nums[i])
		} else {
			arr[search(arr, nums[i])] = nums[i]
		}
	}
	return len(arr)
}

func search(arr []int, num int) int {
	left, right := 0, len(arr)-1
	for left < right {
		// 為了防止 overflow，因此用此種方式計算 mid。
		mid := left + (right-left)/2
		if num == arr[mid] {
			return mid
		}
		if num < arr[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
