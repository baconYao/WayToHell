package squaringasortedarray

import "github.com/baconYao/WayToHell/utils"

// time: O(n), space: O(k) where k is the number of element in input slice
func SquaringASortedArray(input []int) []int {
	frontPointer := 0
	rearPointer := len(input) - 1
	result := []int{}

	for frontPointer <= rearPointer {
		if frontPointer == rearPointer {
			tmp := utils.SquareInt(input[frontPointer])
			// insert to head
			result = append([]int{tmp}, result...)
			frontPointer += 1
			continue
		}
		frontSquareValue := utils.SquareInt(input[frontPointer])
		rearSquareValue := utils.SquareInt(input[rearPointer])

		if frontSquareValue > rearSquareValue {
			result = append([]int{frontSquareValue}, result...)
			frontPointer += 1
		} else {
			result = append([]int{rearSquareValue}, result...)
			rearPointer -= 1
		}
	}
	return result
}
