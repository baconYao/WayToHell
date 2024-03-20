package findmaximuminslidingwindow

func FindMaxSlidingWindowV2(nums []int, w int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	if w > n {
		w = n
	}

	output := make([]int, 0)
	currentWindow := []int{} // currWindow stores the index of nums

	// Iterate over the first w elemnts
	for i := 0; i < w; i++ {
		currentWindow = cleanUp(i, currentWindow, nums)
		currentWindow = append(currentWindow, i)
	}

	// Appending the maximum element of the current window
	output = append(output, nums[currentWindow[0]])

	// Iterate over the remaining
	for i := w; i < n; i++ {
		currentWindow = cleanUp(i, currentWindow, nums)
		// Remove first index from the currentWindow if it has fallen out of the current window
		if len(currentWindow) > 0 && currentWindow[0] <= i-w {
			currentWindow = currentWindow[1:]
		}
		currentWindow = append(currentWindow, i)
		output = append(output, nums[currentWindow[0]])
	}

	return output
}

// clenaUp function cleans up the window
func cleanUp(i int, currentWindow []int, nums []int) []int {
	for len(currentWindow) > 0 && nums[i] >= nums[currentWindow[len(currentWindow)-1]] {
		// Remove all the indexes from currentWindow whose corresponding values are
		// smaller than or equal to the current element --> nums[i]
		currentWindow = currentWindow[:len(currentWindow)-1]
	}
	return currentWindow
}
