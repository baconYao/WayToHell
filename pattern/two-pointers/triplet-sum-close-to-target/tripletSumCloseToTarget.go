package tripletsumclosetotarget

import (
	"math"
	"sort"

	"github.com/baconYao/WayToHell/utils"
)

// Time: O(nlogn + N2) = O(N2) | Space: O(n)
func TripletSumCloseToTarget(input []int, target int) int {
	sort.Ints(input)
	smallestDifference := math.MaxInt64
	for i := 0; i < len(input)-2; i++ {
		left := i + 1
		right := len(input) - 1
		for left < right {
			difference := target - input[i] - input[left] - input[right]
			// Find the exac sum
			if difference == 0 {
				return target
			}

			if utils.AbsInt(difference) < utils.AbsInt(smallestDifference) || (utils.AbsInt(difference) == utils.AbsInt(smallestDifference) && difference > smallestDifference) {
				smallestDifference = difference
			}
			if difference > 0 {
				left += 1
			} else {
				right -= 1
			}
		}

	}
	return target - smallestDifference
}
