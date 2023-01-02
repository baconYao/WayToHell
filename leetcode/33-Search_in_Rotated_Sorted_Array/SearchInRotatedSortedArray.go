package searchinrotatedsortedarray

func search(nums []int, target int) int {
	n := len(nums)

	// Find the pivot.
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// n = 7
	// [4,5,6,7,0,1,2]
	// [4,5,6,7,0,1,2,4,5,6,7]

	pivot := left

	// Regular binary search
	left, right = pivot, pivot-1+n // 4, 4+7-1=10
	for left <= right {
		mid := left + (right-left)/2
		midVal := nums[mid%n]

		if midVal > target {
			right = mid - 1
		} else if midVal < target {
			left = mid + 1
		} else {
			return mid % n
		}
	}

	return -1
}
