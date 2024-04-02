package minimumwindowsubsequence

import (
	"github.com/baconYao/WayToHell/utils"
)

func CharacterReplacement(s string, k int) int {
	charFrequency := make(map[string]int)
	lengthOfMaxSubstring := -1
	mostFreqChar := 0
	left := 0
	for right, char := range s {
		c := string(char)
		charFrequency[c] += 1
		// Obtain the count of the most frequently occuring characters
		mostFreqChar = utils.Max(mostFreqChar, charFrequency[c])

		if right-left+1-mostFreqChar > k {
			charFrequency[string(s[left])] -= 1
			left += 1
		}
		lengthOfMaxSubstring = utils.Max(right-left+1, lengthOfMaxSubstring)
	}

	return lengthOfMaxSubstring
}

func characterReplacement2(s string, k int) int {
	cnts := make([]int, 26)
	maxCnts, start, ans := 0, 0, 0

	for end, v := range s {
		cnts[v-'A']++
		maxCnts = utils.Max(maxCnts, cnts[v-'A'])
		for end-start-maxCnts+1 > k {
			cnts[s[start]-'A']--
			//maxCnts = max(maxCnts, cnts[s[start] - 'A'])
			start++
		}
		ans = utils.Max(ans, end-start+1)
	}
	return ans
}
