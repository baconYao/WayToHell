package groupanagrams

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strconv"
)

type sortByte []byte

func (s sortByte) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortByte) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByte) Len() int {
	return len(s)
}

func SortString(str string) string {
	chars := []byte(str)
	sort.Sort(sortByte(chars))
	return string(chars)
}

// Method 1
// Hash table + sort string
func groupAnagrams(strs []string) [][]string {
	var result [][]string
	hashTable := make(map[string]int)
	for _, s := range strs {
		sorted := SortString(s)
		if idx, ok := hashTable[sorted]; ok {
			result[idx] = append(result[idx], s)
		} else {
			hashTable[sorted] = len(result)
			result = append(result, []string{s})
		}
	}
	return result
}

// Another way to sort
func sortString(str string) string {
	strBytes := []byte(str)
	sort.Slice(strBytes, func(i, j int) bool { return strBytes[i] < strBytes[j] })
	return string(strBytes)
}

//==================================================================
// Method 2
// Hash table + hash 過的 26 alphabet 字母
func groupAnagrams2(strs []string) [][]string {
	var result [][]string
	hashTable := make(map[string]int)
	for _, s := range strs {
		es := encode(s)
		if idx, ok := hashTable[es]; ok {
			result[idx] = append(result[idx], s)
		} else {
			hashTable[es] = len(result)
			result = append(result, []string{s})
		}
	}
	return result
}

func encode(strs string) string {
	alphabet := make([]int, 26)
	for idx := range strs {
		alphabet[strs[idx]-'a']++
	}
	return computeHashKeyForList(alphabet, "0")
}

func computeHashKeyForList(list []int, delim string) string {
	var buffer bytes.Buffer
	for i, _ := range list {
		buffer.WriteString(strconv.Itoa(list[i]))
		buffer.WriteString(delim)
	}
	return GetMD5Hash(buffer.String())
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
