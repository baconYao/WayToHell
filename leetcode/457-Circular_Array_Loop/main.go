package circulararrayloop

func circularArrayLoop(nums []int) bool {
	size := len(nums)
	// for idx, val := range nums {
	for i := 0; i < len(nums); i++ {
		slow := i
		fast := i

		direction := nums[i] >= 0

		for {
			slow = moveTo(nums[slow], slow, size)
			if isNotCycle(nums, direction, slow) {
				break
			}

			fast = moveTo(nums[fast], fast, size)
			if isNotCycle(nums, direction, fast) {
				break
			}
			fast = moveTo(nums[fast], fast, size)
			if isNotCycle(nums, direction, fast) {
				break
			}

			if slow == fast {
				return true
			}
		}
	}
	return false
}

// val is the value of nums[idx]
// idx is the index of nums
// size is the length of nums
func moveTo(val, idx, size int) int {
	result := (val + idx) % size
	if result < 0 {
		result += size
	}
	return result
}

func isNotCycle(nums []int, prevDirection bool, idx int) bool {
	currentDirection := nums[idx] >= 0
	if (currentDirection != prevDirection) || (nums[idx]%len(nums) == 0) {
		return true
	}
	return false
}
