package main

import "fmt"

// Method 1
// Runtime: 29 ms
// Memory Usage: 5.4 MB
func TwoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left += 1
		} else {
			right -= 1
		}
	}

	return []int{}
}

// Method 2
func TwoSumMethod2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		fmt.Println("---------")
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left = rightBinarySearch(numbers, left+2, right-1, target-numbers[right])
		} else {
			fmt.Println("~~")
			right = leftBinarySearch(numbers, left+1, right-2, target-numbers[left])
		}
	}

	return []int{}
}

func leftBinarySearch(numbers []int, l, r, target int) int {
	l--
	r++
	fmt.Println(l, r)
	for r > l+1 {
		m := l + (r-l)/2
		fmt.Println("m:", m)
		if numbers[m] > target {
			r = m
		} else {
			l = m
		}
	}
	return r
}

func rightBinarySearch(numbers []int, l, r, target int) int {
	l--
	r++

	for r > l+1 {
		m := l + (r-l)/2
		if numbers[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	return l
}

func main() {
	res := TwoSumMethod2([]int{2, 7, 11, 15}, 9)
	fmt.Print(res)
}
