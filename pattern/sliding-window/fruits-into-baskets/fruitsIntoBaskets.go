package fruitsintobaskets

import "github.com/baconYao/WayToHell/utils"

// Time: O(n) | Space: O(1)
func FruitsIntoBaskets(fruits []string) int {
	windowStart := 0
	maxWindowSize := -1
	charFrequency := make(map[string]int)

	for windowEnd, v := range fruits {
		sc := string(v)
		charFrequency[sc] += 1
		for len(charFrequency) > 2 {
			charAtStart := string(fruits[windowStart])
			charFrequency[charAtStart] -= 1
			windowStart += 1
			if charFrequency[charAtStart] == 0 {
				delete(charFrequency, charAtStart)
			}
		}
		maxWindowSize = utils.Max(maxWindowSize, windowEnd-windowStart+1)
	}
	return maxWindowSize
}
