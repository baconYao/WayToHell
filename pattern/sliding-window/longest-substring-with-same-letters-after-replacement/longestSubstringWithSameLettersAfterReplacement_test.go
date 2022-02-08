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
		{1, "abbcb", 4, "Replace the 'c' with 'b' to have the longest repeating substring \"bbbb\"."},
		{1, "abccde", 3, "Replace the 'b' or 'd' with 'c' to have the longest repeating substring \"ccc\"."},
	}

	for _, e := range theTests {
		actualResult := LongestSubstringWithSameLettersAfterReplacement(e.K, e.S)
		assert.Equal(e.expectedResult, actualResult, e.errorDescription)
	}
}
