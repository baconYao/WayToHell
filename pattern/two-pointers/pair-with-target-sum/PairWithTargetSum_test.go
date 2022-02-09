package pairwithtargetsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairWithTargetSum(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		target           int // number of replacement 0
		array            []int
		expectedResult   []int
		errorDescription string
	}{
		{6, []int{1, 2, 3, 4, 6}, []int{1, 3}, "The numbers at index 1 and 3 add up to 6: 2+4=6"},
		{11, []int{2, 5, 9, 11}, []int{0, 2}, "The numbers at index 0 and 2 add up to 11: 2+9=11"},
		{18, []int{2, 5, 9, 11}, []int{-1, -1}, "There's no set to match 18"},
	}

	for _, e := range theTests {
		actualResult := PairWithTargetSum(e.target, e.array)
		assert.Equal(e.expectedResult, actualResult, e.errorDescription)
	}
}
