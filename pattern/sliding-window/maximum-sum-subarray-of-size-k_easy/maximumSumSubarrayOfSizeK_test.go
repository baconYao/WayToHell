package maximumsumsubarrayofsizekeasy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSumSubarrayOfSizeK(t *testing.T) {
	assert := assert.New(t)

	K := 1
	input := []int{2}
	actualResult := MaximumSumSubarrayOfSizeK(K, input)
	expectedResult := 2
	assert.Equal(actualResult, expectedResult, "The numbers of input is equal to the number of K")

	K = 3
	input = []int{2, 1, 5, 1, 3, 2}
	actualResult = MaximumSumSubarrayOfSizeK(K, input)
	expectedResult = 9
	assert.Equal(actualResult, expectedResult, "Result should be 9")

	K = 2
	input = []int{2, 3, 4, 1, 5}
	actualResult = MaximumSumSubarrayOfSizeK(K, input)
	expectedResult = 7
	assert.Equal(actualResult, expectedResult, "Result should be 7")
}
