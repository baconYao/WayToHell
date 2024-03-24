package minimumwindowsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input1         string
		input2         string
		expectedResult string
	}{
		{"abcdebdde", "bde", "bcde"},
		{"fgrqsqsnodwmxzkzxwqegkndaa", "kzed", "kzxwqegknd"},
		{"michmznaitnjdnjkdsnmichmznait", "michmznait", "michmznait"},
		{"afgegrwgwga", "aa", "afgegrwgwga"},
		{"abcdbebe", "bbe", "bebe"},
	}

	for _, e := range theTests {
		actualResult := MinWindow(e.input1, e.input2)
		assert.Equal(e.expectedResult, actualResult)
	}
}
