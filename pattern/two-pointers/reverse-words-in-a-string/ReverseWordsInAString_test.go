package reversewordsinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWordsInAString(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          string
		expectedResult string
	}{
		{"We love GO ", "GO love We"},
		{"To be or not to be", "be to not or be To"},
		{"You are amazing", "amazing are You"},
		{"Hello     World", "World Hello"},
		{" Hey", "Hey"},
	}

	for _, e := range theTests {
		actualResult := ReverseWordsInAString(e.input)
		assert.Equal(e.expectedResult, actualResult)
	}
}
