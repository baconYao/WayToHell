package longestsubarraywithonesafterreplacement

import (
	"github.com/baconYao/WayToHell/utils"
)

func LongestSubarrayWithOnesAfterReplacement(k int, arr []int) int {
	windowStart := 0
	oneCount := 0
	longestSubarray := 0

	for windowEnd, number := range arr {
		if number == 1 {
			oneCount += 1
		}
		// 確認 window 內的 0 的數量是否多於 k
		for windowEnd-windowStart+1-oneCount > k {
			// number at start pointer
			ns := arr[windowStart]
			if ns == 1 {
				oneCount -= 1
			}
			windowStart += 1
		}
		longestSubarray = utils.Max(longestSubarray, windowEnd-windowStart+1)
	}
	return longestSubarray
}
