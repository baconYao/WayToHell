package longestsubstringwithmaximunkdistinctcharacters

import (
	"github.com/baconYao/WayToHell/utils"
)

// Time: O(n) | Space: O(K) where K is the number of distinct characters
func longestSubstringWithMaximunKDistinctCharacters(K int, S string) int {
	windowStart := 0
	maxLength := 0
	charFrequency := make(map[string]int)

	for windowEnd, char := range S {
		// rune to string type
		sc := string(char)
		// Golang 對於不存在於 map 的 key，會給予其一個初始化得值
		// 在此就是 0 (因為 value 是 int)，所以我們可以忽略下面這段判斷式，直接 +1 直接即可
		// if _, exist := charFrequency[sc]; !exist {
		// 	charFrequency[sc] = 0
		// }
		charFrequency[sc] += 1
		// shrink the window if the number of distinct characters are more than K
		for len(charFrequency) > K {
			// wsc: window start character
			wsc := string(S[windowStart])
			charFrequency[wsc] -= 1
			windowStart += 1
			if charFrequency[wsc] == 0 {
				delete(charFrequency, wsc)
			}
		}
		maxLength = utils.Max(maxLength, windowEnd-windowStart+1)
	}
	return maxLength
}
