package fruitsintobaskets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrutsIntoBaskets(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input            []string
		expectedResult   int
		errorDescription string
	}{
		{
			[]string{"A", "B", "C", "A", "C"},
			3,
			"We can put 2 'C' in one basket and one 'A' in the other from the subarray ['C', 'A', 'C']",
		},
		{
			[]string{"A", "B", "C", "B", "B", "C"},
			5,
			"We can put 3 'B' in one basket and two 'C' in the other basket. This can be done if we start with the second letter: ['B', 'C', 'B', 'B', 'C']",
		},
	}

	for _, e := range theTests {
		actualResult := FruitsIntoBaskets(e.input)
		assert.Equal(e.expectedResult, actualResult, e.errorDescription)
	}
}
