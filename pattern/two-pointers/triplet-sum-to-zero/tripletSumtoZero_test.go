package tripletsumtozero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTripletSumToZero(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          []int
		expectedResult [][3]int
		description    string
	}{
		{[]int{-3, 0, 1, 2, -1, 1, -2}, [][3]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}}, "There are four unique triplets whose sum is equal to zero."},
		{[]int{-5, 2, -1, -2, 3}, [][3]int{{-5, 2, 3}, {-2, -1, 3}}, "There are two unique triplets whose sum is equal to zero."},
		{[]int{-5, 2, 2, -1, -2, 3, 3}, [][3]int{{-5, 2, 3}, {-2, -1, 3}}, "There are two unique triplets whose sum is equal to zero."},
		{[]int{-4, 2, 2}, [][3]int{{-4, 2, 2}}, ""},
	}

	for _, e := range theTests {
		actualResult := TripletSumToZero(e.input)
		assert.Equal(e.expectedResult, actualResult)
	}
}
