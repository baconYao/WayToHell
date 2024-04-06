package minimunwindowsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		s              string
		t              string
		expectedResult string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"ABCD", "ABC", "ABC"},
		{"XYZYX", "XYZ", "XYZ"},
		{"ABXYZJKLSNFC", "ABC", "ABXYZJKLSNFC"},
		{"AAAAAAAAAAA", "A", "A"},
		{"ABDFGDCKAB", "ABCD", "DCKAB"},
	}

	for _, e := range theTests {
		actualResult := minWindow(e.s, e.t)
		assert.Equal(e.expectedResult, actualResult)
	}
}
