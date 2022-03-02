package removeduplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input            []int
		expectedResult   int
		errorDescription string
	}{
		{[]int{2, 3, 3, 3, 6, 9, 9}, 4, "The first four elements after removing the duplicates will be [2, 3, 6, 9]."},
		{[]int{2, 2, 2, 11}, 2, "The first two elements after removing the duplicates will be [2, 11]."},
	}

	for _, e := range theTests {
		actualResult := RemoveDuplicates(e.input)
		assert.Equal(e.expectedResult, actualResult, e.errorDescription)
	}
}
