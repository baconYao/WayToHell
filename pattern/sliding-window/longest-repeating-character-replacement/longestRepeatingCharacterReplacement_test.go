package minimumwindowsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input1         string
		input2         int
		expectedResult int
	}{
		{"abab", 2, 4},
		{"aababba", 1, 4},
		{"aaacbbbaabab", 2, 6},
		{"aaacbbbaabab", 1, 4},
		{"dippitydip", 4, 6},
		{"coollooc", 2, 6},
		{"aaaaaaaaaa", 2, 10},
	}

	for _, e := range theTests {
		actualResult := CharacterReplacement(e.input1, e.input2)
		assert.Equal(e.expectedResult, actualResult)
	}
}
