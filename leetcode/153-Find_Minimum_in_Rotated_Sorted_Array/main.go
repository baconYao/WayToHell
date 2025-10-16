package findminimuminrotatedsortedarray

func findMinInRotatedArray(arr []int) int {
	left, right, minimum := 0, len(arr)-1, arr[0]

	// "equal" means there's only one element in this array
	// "less than" means this array is sorted ascending
	if arr[left] <= arr[right] {
		return arr[0]
	}

	for left <= right {
		mid := left + (right-left)/2
		// if mid > mid+1, return mid+1, since pivot is mid+1 which is also the minimum element
		if arr[mid] > arr[mid+1] {
			minimum = arr[mid+1]
			break
		}
		// if mid-1 > mid, return mid, since pivot is mid which is also the minimum element
		if arr[mid] < arr[mid-1] {
			minimum = arr[mid]
			break
		}

		// If the value of the middle element is greater than the left element,
		// then the minimum value must be to the right of the middle element.
		if arr[left] < arr[mid] {
			left = mid + 1
		} else {
			// Otherwise, set the right pointer to the element
			// before the middle element [mid - 1].
			right = mid - 1
		}
	}

	return minimum
}
