package circulararrayloop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularArrayLoop(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          []int
		expectedResult bool
	}{
		{[]int{1, 3, -2, -4, 1}, true},
		{[]int{2, 1, -1, -2}, false},
		{[]int{5, 4, -2, -1, 3}, false},
		{[]int{1, 2, -3, 3, 4, 7, 1}, true},
		{[]int{3, 3, 1, -1, 2}, true},
	}

	for _, e := range theTests {
		actualResult := CircularArrayLoop(e.input)
		assert.Equal(e.expectedResult, actualResult)
	}
}
