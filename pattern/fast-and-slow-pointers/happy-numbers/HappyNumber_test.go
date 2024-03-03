package happynumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappyNumber(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          int
		expectedResult bool
	}{
		{2147483646, false},
		{1, true},
		{19, true},
		{8, false},
		{7, true},
	}

	for _, e := range theTests {
		actualResult := IsHappyNumber(e.input)
		assert.Equal(e.expectedResult, actualResult)
	}
}
