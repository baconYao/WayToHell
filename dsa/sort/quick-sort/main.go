package main

import "fmt"

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	if left < right {
		partitionIndex := partition(arr, left, right)
		quickSort(arr, left, partitionIndex-1)
		quickSort(arr, partitionIndex+1, right)
	}
}

// https://www.youtube.com/watch?v=COk73cpQbFQ
func partition(arr []int, left, right int) int {
	pivot := arr[right]
	partitionIndex := left
	for i := left; i < right; i++ {
		if arr[i] <= pivot {
			arr[partitionIndex], arr[i] = arr[i], arr[partitionIndex]
			partitionIndex++
		}
	}

	arr[partitionIndex], arr[right] = arr[right], arr[partitionIndex]
	return partitionIndex
}

// func partition(arr []int, left, right int) int {
// 	if left < 0 || right >= len(arr) {
// 		return left
// 	}

// 	pivotIndex := left
// 	start := left
// 	end := right // end 具有重要意義，最終會成為 pivot 的正確位置

// 	for start < end {
// 		for arr[start] <= arr[pivotIndex] && start < end {
// 			start++
// 		}

// 		// end 的大於 pivot 時，end 往左移動。終止條件是 start <= end，因為 start 的值最終會大於 pivot。
// 		// 因此不能使用 start < end 作為終止條件。
// 		for arr[end] > arr[pivotIndex] && start <= end {
// 			end--
// 		}

// 		if start < end {
// 			arr[start], arr[end] = arr[end], arr[start]
// 		}

// 	}
// 	arr[pivotIndex], arr[end] = arr[end], arr[pivotIndex]
// 	return end
// }
