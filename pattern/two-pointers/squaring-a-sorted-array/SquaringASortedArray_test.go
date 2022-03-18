package squaringasortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquaringASortedArray(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          []int
		expectedResult []int
	}{
		{[]int{-2, -1, 0, 2, 3}, []int{0, 1, 4, 4, 9}},
		{[]int{-3, -1, 0, 1, 2}, []int{0, 1, 1, 4, 9}},
	}

	for _, e := range theTests {
		actualResult := SquaringASortedArray(e.input)
		assert.Equal(e.expectedResult, actualResult)
	}
}
