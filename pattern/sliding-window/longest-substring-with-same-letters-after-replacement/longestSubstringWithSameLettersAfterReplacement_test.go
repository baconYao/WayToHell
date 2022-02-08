package longestsubstringwithsamelettersafterreplacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstringWithSameLettersAfterReplacement(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		K                int // number of replacement characters
		S                string
		expectedResult   int
		errorDescription string
	}{
		{2, "aabccbb", 5, "Replace the two 'c' with 'b' to have the longest repeating substring \"bbbbb\"."},
		{2, "adcbb", 4, "Replace the 'd' and 'c' with 'b' to have the longest repeating substring \"bbbb\"."},
		{1, "abbcb", 4, "Replace the 'c' with 'b' to have the longest repeating substring \"bbbb\"."},
		{1, "abccde", 3, "Replace the 'b' or 'd' with 'c' to have the longest repeating substring \"ccc\"."},
		{2, "abc", 3, "Replace the 'a' or 'd' or 'c' to have the longest repeating substring \"aaa\" or \"bbb\" or \"ccc\"."},
		{1, "abc", 2, "Replace the 'a' or 'b' or 'c' to have the longest repeating substring \"aa\", \"bb\", \"cc\"."},
	}

	for _, e := range theTests {
		actualResult := LongestSubstringWithSameLettersAfterReplacement(e.K, e.S)
		assert.Equal(e.expectedResult, actualResult, e.errorDescription)
	}
}
