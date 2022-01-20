package smallestsubarraywithagreatersum

import (
	"math"

	"github.com/baconYao/WayToHell/utils"
)

// Time: O(n) | Space: O(1)
func SmallestSubarrayWithAGreaterSum(S int, nums []int) int {
	smallestArrayLength := math.MaxInt64
	sumOfWindow := 0
	windowStart := 0

	for windowEnd, v := range nums {
		sumOfWindow += v
		for sumOfWindow >= S {
			smallestArrayLength = utils.Min(smallestArrayLength, windowEnd-windowStart+1)
			sumOfWindow -= nums[windowStart]
			windowStart += 1
		}
	}

	if smallestArrayLength == math.MaxInt64 {
		return 0
	}

	return smallestArrayLength
}
