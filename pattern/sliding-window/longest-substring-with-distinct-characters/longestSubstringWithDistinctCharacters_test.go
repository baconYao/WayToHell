package longestsubstringwithdistinctcharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstringWithDistinctCharacters(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		S                string
		expectedResult   int
		errorDescription string
	}{
		{"aabccbb", 3, "The longest substring with distinct characters is \"abc\"."},
		{"abbbb", 2, "The longest substring with distinct characters is \"ab\"."},
		{"abccde", 3, "Longest substrings with distinct characters are \"abc\" & \"cde\"."},
	}

	for _, e := range theTests {
		actualResult := LongestSubstringWithDistinctCharacters(e.S)
		assert.Equal(e.expectedResult, actualResult, e.errorDescription)
	}
}
