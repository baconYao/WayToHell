package main

// Method 1, sliding window
func LengthOfLongestSubstring(s string) int {
	charFrequency := make(map[string]int)
	front := 0
	longestSize := 0

	for rear, char := range s {
		c := string(char)
		charFrequency[c] += 1
		for charFrequency[c] > 1 && front <= rear {
			frontChar := string(s[front])
			charFrequency[frontChar] -= 1
			front += 1
		}
		longestSize = Max(longestSize, rear-front+1)
	}
	return longestSize
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// ============================================================

// Method 2
func LengthOfLongestSubstring2(s string) int {
	// 256 是 letter, symbol, digits and space 的總數
	location := [256]int{}

	for i := range location {
		location[i] = -1
	}

	left, maxLen := 0, 0

	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			// 移動到該同樣字元的"時間點"之後
			left = location[s[i]] + 1
		} else if i-left+1 > maxLen {
			maxLen = i + 1 - left
		}
		// 紀錄字元出現的"時間點"
		location[s[i]] = i
	}
	return maxLen
}
