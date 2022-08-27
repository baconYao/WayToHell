package ransomnote

// Method 1
// Hash table
func RansomNote(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	char_frequency := make(map[string]int)
	// 計算 ransomNote 和 magazine 出現的頻率
	for _, v := range ransomNote {
		c := string(v)
		char_frequency[c] -= 1
	}
	for _, v := range magazine {
		c := string(v)
		char_frequency[c] += 1
	}
	// 最後檢查是否有字母出現的頻率
	for _, v := range char_frequency {
		if v < 0 {
			return false
		}
	}
	return true
}

// Method 2
// Array
func RansomNote2(ransomNote string, magazine string) bool {
	arrMagazine := make([]int, 26)
	for idx := range magazine {
		arrMagazine[int(magazine[idx])-int('a')]++
	}

	for idx := range ransomNote {
		arrMagazine[int(ransomNote[idx])-int('a')]--
		if arrMagazine[int(ransomNote[idx])-int('a')] < 0 {
			return false
		}
	}
	return true
}
