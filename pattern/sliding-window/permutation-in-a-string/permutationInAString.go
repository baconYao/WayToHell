package permutationinastring

import "reflect"

func permutationInAString(input, pattern string) bool {
	windowStart := 0
	windowSize := len(pattern)
	inputFrequency := make(map[string]int)
	patternFrequency := make(map[string]int)
	for _, v := range pattern {
		patternFrequency[string(v)] += 1
	}

	for windowEnd, v := range input {
		inputFrequency[string(v)] += 1

		if windowEnd-windowStart+1 > windowSize {
			wsc := string(input[windowStart])
			inputFrequency[wsc] -= 1
			windowStart += 1
			if inputFrequency[wsc] == 0 {
				delete(inputFrequency, wsc)
			}
		}

		if windowEnd-windowStart+1 == windowSize {
			isEqual := reflect.DeepEqual(patternFrequency, inputFrequency)
			if isEqual {
				return true
			}
		}
	}

	return false
}
