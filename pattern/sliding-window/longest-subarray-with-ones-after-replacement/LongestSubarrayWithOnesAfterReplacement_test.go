package longestsubarraywithonesafterreplacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarrayWithOnesAfterReplacement(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		K                int // number of replacement 0
		array            []int
		expectedResult   int
		errorDescription string
	}{
		{2, []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 6, "Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6."},
		{3, []int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 9, "Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9."},
	}

	for _, e := range theTests {
		actualResult := LongestSubarrayWithOnesAfterReplacement(e.K, e.array)
		assert.Equal(e.expectedResult, actualResult, e.errorDescription)
	}
}
