package findtheduplicatenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          []int
		expectedResult int
	}{
		{[]int{3, 4, 4, 4, 2}, 4},
		{[]int{1, 1}, 1},
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{1, 3, 6, 2, 7, 3, 5, 4}, 3},
		{[]int{1, 2, 2}, 2},
	}

	for _, e := range theTests {
		actualResult := FindDuplicate(e.input)
		assert.Equal(e.expectedResult, actualResult)
	}
}
