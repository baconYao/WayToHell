package findmaximuminslidingwindow

// Time Limit Exceeded

func FindMaxSlidingWindow(nums []int, w int) []int {
	result := make([]int, 0)

	if w > len(nums) {
		w = len(nums)
	}

	front := 0
	for rear := w - 1; rear < len(nums); rear++ {
		currMaxInWindow := findMaximunNum(nums, front, rear)
		result = append(result, currMaxInWindow)
		front += 1
	}

	return result
}

func findMaximunNum(nums []int, front, rear int) int {
	max := -99999
	for i := front; i <= rear; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
