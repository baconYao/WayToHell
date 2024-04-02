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
