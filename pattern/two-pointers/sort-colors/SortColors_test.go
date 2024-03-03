package sortcolors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortColors(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          []int
		expectedResult []int
	}{
		{[]int{0, 1, 0}, []int{0, 0, 1}},
		{[]int{1}, []int{1}},
		{[]int{2, 2}, []int{2, 2}},
		{[]int{1, 1, 0, 2}, []int{0, 1, 1, 2}},
		{[]int{2, 1, 1, 0, 0}, []int{0, 0, 1, 1, 2}},
	}

	for _, e := range theTests {
		actualResult := SortColors(e.input)
		assert.Equal(e.expectedResult, actualResult)
	}
}
