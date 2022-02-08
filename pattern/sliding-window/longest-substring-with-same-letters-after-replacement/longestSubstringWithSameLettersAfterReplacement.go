package longestsubstringwithsamelettersafterreplacement

import "github.com/baconYao/WayToHell/utils"

// time: O(n), space: O(1)
func LongestSubstringWithSameLettersAfterReplacement(k int, s string) int {
	// 想法: 以 end pointer 走訪過的 letter 為新的基準，比較當前 window 內剩餘的 letter 和此基準點的數量，
	// 			若剩餘 letter 的數量多於 k 時，表示無法完全替換 window 內的 letter，因此需要 shrink window
	// 			(移動 start pointer)
	windowStart := 0
	charFrequency := make(map[string]int)
	maxRepeatLetterCount := 0
	longestSize := 0

	for windowEnd, ce := range s {
		charFrequency[string(ce)] += 1
		maxRepeatLetterCount = utils.Max(maxRepeatLetterCount, charFrequency[string(ce)])
		// window 內，扣掉重複letter後，剩下的就是要被替換的字元
		// 若要被替換的字元多餘 k 時，就 shrink window
		if windowEnd-windowStart+1-maxRepeatLetterCount > k {
			// char at start pointer
			cs := string(s[windowStart])
			charFrequency[cs] -= 1
			if charFrequency[cs] == 0 {
				delete(charFrequency, cs)
			}
			windowStart += 1
		}
		longestSize = utils.Max(longestSize, windowEnd-windowStart+1)
	}

	return longestSize
}
