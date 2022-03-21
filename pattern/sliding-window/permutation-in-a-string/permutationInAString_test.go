package permutationinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutationInAString(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input            string
		pattern          string
		expectedResult   bool
		errorDescription string
	}{
		{"oidbcaf", "abc", true, "The string contains \"bca\" which is a permutation of the given pattern."},
		{"odicf", "dc", false, "No permutation of the pattern is present in the given string as a substring."},
		{"bcdxabcdy", "bcdyabcdx", true, "Both the string and the pattern are a permutation of each other."},
		{"aaacb", "abc", true, "The string contains \"acb\" which is a permutation of the given pattern."},
	}

	for _, e := range theTests {
		actualResult := permutationInAString(e.input, e.pattern)
		assert.Equal(e.expectedResult, actualResult, e.errorDescription)
	}
}
