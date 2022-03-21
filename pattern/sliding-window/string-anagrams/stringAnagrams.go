package stringanagrams

import "reflect"

// Time: O(n) | Space: O(K) where K is the character number of pattern
func stringAnagrams(input, pattern string) []int {
	patternFrequency := make(map[string]int)
	charFrequency := make(map[string]int)
	windowStart := 0
	windowSize := len(pattern)
	returnIndices := []int{}

	for _, c := range pattern {
		patternFrequency[string(c)] += 1
	}

	for windowEnd, v := range input {
		charFrequency[string(v)] += 1
		if windowEnd-windowStart+1 == windowSize {
			if reflect.DeepEqual(patternFrequency, charFrequency) {
				returnIndices = append(returnIndices, windowStart)
			}
			wsc := string(input[windowStart])
			windowStart += 1
			charFrequency[wsc] -= 1
			if charFrequency[wsc] == 0 {
				delete(charFrequency, wsc)
			}
		}
	}

	return returnIndices
}
