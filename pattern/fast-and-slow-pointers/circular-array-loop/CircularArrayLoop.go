package circulararrayloop

func CircularArrayLoop(nums []int) bool {
	size := len(nums)
	for i := 0; i < size; i++ {
		slow, fast := i, i
		// true if it's forward direction, else flase
		direction := nums[i] > 0
		for {
			slow = nextStep(slow, nums[slow], size)
			if isNotCycle(nums, direction, slow) {
				break
			}
			// Move fast pointer two times
			fast = nextStep(fast, nums[fast], size)
			if isNotCycle(nums, direction, fast) {
				break
			}
			fast = nextStep(fast, nums[fast], size)
			if isNotCycle(nums, direction, fast) {
				break
			}
			// If slow and fast pointers meet, a cycle is found
			if slow == fast {
				return true
			}
		}
	}
	return false
}

// nextStep function calculates the next position (index) of nums
func nextStep(index, value, size int) int {
	result := (index + value) % size
	if result < 0 {
		result += size
	}
	return result
}

// isNotCycle function detects a cycle doesn't exist
func isNotCycle(nums []int, preDirection bool, index int) bool {
	currDirection := nums[index] > 0
	// 1. If the direction changes or 2. Only one element in the loop
	if currDirection != preDirection || nums[index]%len(nums) == 0 {
		return true
	}
	return false
}
