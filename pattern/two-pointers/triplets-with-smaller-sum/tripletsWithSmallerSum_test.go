package tripletswithsmallersum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTripletsWithSmallerSum(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          []int
		target         int
		expectedResult int
		description    string
	}{
		{[]int{-1, 0, 2, 3}, 3, 2, "There are two triplets whose sum is less than the target: [-1, 0, 3], [-1, 0, 2]"},
		{[]int{-1, 4, 2, 1, 3}, 5, 4, "There are four triplets whose sum is less than the target: [-1, 1, 4], [-1, 1, 3], [-1, 1, 2], [-1, 2, 3]"},
		{[]int{39, 28, 4, -40}, 5, 2, "There are two triplets whose sum is less than the target: [39, 4, -40], [28, 4, -40]"},
	}

	for _, e := range theTests {
		actualResult := TripletsWithSmallerSum(e.input, e.target)
		assert.Equal(e.expectedResult, actualResult)
	}
}
