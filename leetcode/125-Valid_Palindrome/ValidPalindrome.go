package validpalindrome

import (
	"regexp"
	"strings"
)

func ValidPalindrome(s string) bool {
	cs := sanitizeString(s)
	left, right := 0, len(cs)-1
	for left < right {
		if cs[left] != cs[right] {
			return false
		}
		left++
		right--
	}
	return true
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func sanitizeString(s string) string {
	// 如何濾掉 non alphabet 的字元
	// https://gosamples.dev/remove-non-alphanumeric/
	as := nonAlphanumericRegex.ReplaceAllString(s, "")
	return strings.ToUpper(as)
}
