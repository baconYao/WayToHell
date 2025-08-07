package sortcolors

func sortColors(nums []int) {
	// front pointer point to the last position of red (0)
	// rear pointer point to the first position of blue (2)
	front, rear := 0, len(nums)-1
	i := 0
	for i <= rear {
		switch nums[i] {
		case 0:
			nums[front], nums[i] = nums[i], nums[front]
			i++
			front++
		case 1:
			i++
		case 2:
			nums[rear], nums[i] = nums[i], nums[rear]
			rear--
		}

	}
}
