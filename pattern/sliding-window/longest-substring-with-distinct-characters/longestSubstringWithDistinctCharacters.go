package longestsubstringwithdistinctcharacters

import "github.com/baconYao/WayToHell/utils"

func LongestSubstringWithDistinctCharacters(s string) int {
	charFrequency := make(map[string]int)
	windowStart := 0
	longestSize := 0

	for windowEnd, char := range s {
		c := string(char)
		charFrequency[c] += 1
		// 當發現有重複的字元時，需要從 start pointer 所在的位置開始往 end pointer 位置掃描，並逐一排除
		// start pointer 經過的所有字元，直到沒有重複的字元或是 start 和 end pointer 重疊時才停止。
		for charFrequency[c] > 1 && windowStart <= windowEnd {
			charAtStart := string(s[windowStart])
			charFrequency[charAtStart] -= 1
			windowStart += 1
		}
		longestSize = utils.Max(longestSize, windowEnd-windowStart+1)
	}

	return longestSize
}
