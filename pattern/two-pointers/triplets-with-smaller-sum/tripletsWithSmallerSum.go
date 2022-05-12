package tripletswithsmallersum

import (
	"sort"
)

// Time: O(n*logn + n^2) = O(N^2) | Space: O(1)
func TripletsWithSmallerSum(input []int, target int) int {
	countOfSmallerSum := 0
	if len(input) < 3 {
		return countOfSmallerSum
	}

	sort.Sort(sort.Reverse(sort.IntSlice(input)))

	for i := 0; i < len(input)-2; i++ {
		left := i + 1
		right := len(input) - 1
		for right > left {
			if input[i]+input[left]+input[right] < target {
				countOfSmallerSum++
			}
			left++
		}
	}

	return countOfSmallerSum
}
