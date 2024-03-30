package minimumwindowsubsequence

func MinWindow(str1 string, str2 string) string {
	minSubsequence := ""
	minLength := len(str1) + 1 // the maximum length won't exceed str1's length + 1
	indexS1, indexS2 := 0, 0
	for indexS1 < len(str1) {
		if str1[indexS1] == str2[indexS2] {
			if indexS2 == len(str2)-1 {
				start, end := indexS1, indexS1
				for {
					if str1[start] == str2[indexS2] {
						indexS2 -= 1
						if indexS2 < 0 {
							indexS2 = 0
							break
						}
					}
					start -= 1
				}
				if end-start+1 < minLength {
					minLength = end - start + 1
					minSubsequence = str1[start : end+1]
				}
				indexS1 = start + 1
				continue
			} else {
				indexS2 += 1
			}
		}
		indexS1 += 1
	}
	return minSubsequence
}
