package smallestsubarraywithagreatersum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestSubarrayWithAGreaterSum(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		S                int
		input            []int
		expectedResult   int
		errorDescription string
	}{
		{7, []int{2, 1, 5, 2, 3, 2}, 2, "The length of the smallest subarray with a sum greater than or equal to '7' is '2'"},
		{7, []int{2, 1, 5, 2, 8}, 1, "The length of the smallest subarray with a sum greater than or equal to '7' is '1'"},
		{8, []int{3, 4, 1, 1, 6}, 3, "The length of the smallest subarray with a sum greater than or equal to '7' is '3'"},
		{8, []int{1, 1, 1, 1}, 0, "The length of the smallest subarray with a sum greater than or equal to '8' is '0'"},
	}

	for _, e := range theTests {
		actualResult := SmallestSubarrayWithAGreaterSum(e.S, e.input)
		assert.Equal(actualResult, e.expectedResult, e.errorDescription)
	}
}
