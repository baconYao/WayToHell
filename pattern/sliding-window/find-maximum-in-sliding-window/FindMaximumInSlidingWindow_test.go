package findmaximuminslidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxSlidingWindow(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          []int
		input2         int
		expectedResult []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 6, []int{6}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, []int{4, 5, 6, 7, 8, 9, 10}},
		{[]int{-4, 2, -5, 3, 6}, 3, []int{2, 3, 6}},
		{[]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}, 4, []int{3, 3, 3, 3, 3, 3}},
		{[]int{10, 6, 9, -3, 23, -1, 34, 56, 67, -1, -4, -8, -2, 9, 10, 34, 67}, 3, []int{10, 9, 23, 23, 34, 56, 67, 67, 67, -1, -2, 9, 10, 34, 67}},
		{[]int{4, 5, 6, 1, 2, 3}, 1, []int{4, 5, 6, 1, 2, 3}},
		{[]int{9, 5, 3, 1, 6, 3}, 2, []int{9, 5, 3, 6, 6}},
	}

	for _, e := range theTests {
		actualResult := FindMaxSlidingWindow(e.input, e.input2)
		assert.Equal(e.expectedResult, actualResult)
	}
}
