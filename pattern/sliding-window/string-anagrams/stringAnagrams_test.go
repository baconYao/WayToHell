package stringanagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutationInAString(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input            string
		pattern          string
		expectedResult   []int
		errorDescription string
	}{
		{"ppqp", "pq", []int{1, 2}, "The two anagrams of the pattern in the given string are \"pq\" and \"qp\"."},
		{"abbcabc", "abc", []int{2, 3, 4}, "The three anagrams of the pattern in the given string are \"bca\", \"cab\", and \"abc\"."},
		{"aaaa", "a", []int{0, 1, 2, 3}, "All the characters are anagram."},
		{"zxcv", "a", []int{}, "No string anagram."},
	}

	for _, e := range theTests {
		actualResult := stringAnagrams(e.input, e.pattern)
		assert.Equal(e.expectedResult, actualResult, e.errorDescription)
	}
}
