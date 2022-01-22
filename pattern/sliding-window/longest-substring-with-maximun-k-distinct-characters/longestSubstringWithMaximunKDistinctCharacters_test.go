package longestsubstringwithmaximunkdistinctcharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstringWithMaximunKDistinctCharacters(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		K                int // number of distinc characters
		S                string
		expectedResult   int
		errorDescription string
	}{
		{2, "araaci", 4, "The longest substring with no more than '2' distinct characters is \"araa\""},
		{1, "araaci", 2, "The longest substring with no more than '1' distinct characters is \"aa\""},
		{3, "cbbebi", 5, "The longest substrings with no more than '3' distinct characters are \"cbbeb\" & \"bbebi\""},
		{10, "cbbebi", 6, "The longest substring with no more than '10' distinct characters is \"cbbebi\""},
	}

	for _, e := range theTests {
		actualResult := longestSubstringWithMaximunKDistinctCharacters(e.K, e.S)
		assert.Equal(actualResult, e.expectedResult, e.errorDescription)
	}
}
