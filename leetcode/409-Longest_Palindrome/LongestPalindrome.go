package longestpalindrome

// Method 1
func LongestPalindrome(s string) int {
	hashTable := make(map[string]int)

	for _, char := range s {
		c := string(char)
		hashTable[c] += 1
	}
	result := 0
	// 加偶數的數量並且將該 char 移除
	// 加奇數的數量 - 1，並保留在 hash table
	for k, v := range hashTable {
		if v >= 2 && v%2 == 0 {
			result += v
			delete(hashTable, k)
		} else if v >= 2 && v%2 == 1 {
			result += v - 1
			hashTable[k] = 1
		}
	}
	// 目前 hashTable 的 char 的 frequency 都為1
	if len(hashTable) != 0 {
		result += 1
	}
	return result
}

// ==========================================================
// Method 2
func LongestPalindrome2(s string) int {
	smap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		smap[s[i]]++
	}
	sum := 0
	// 用來表示是否有奇數
	flag := false
	for _, v := range smap {
		if v&1 == 1 {
			flag = true
		}
		sum += v / 2
	}
	if flag {
		return sum*2 + 1
	}
	return sum * 2
}
