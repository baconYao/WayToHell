package tripletsumclosetotarget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTripletSumCloseToTarget(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          []int
		target         int
		expectedResult int
		description    string
	}{
		{[]int{-2, 0, 1, 2}, 2, 1, "The triplet [-2, 1, 2] has the closest sum to the target."},
		{[]int{-3, -1, 1, 2}, 1, 0, "The triplet [-3, 1, 2] has the closest sum to the target."},
		{[]int{1, 0, 1, 1}, 100, 3, "The triplet [1, 1, 1] has the closest sum to the target."},
	}

	for _, e := range theTests {
		actualResult := TripletSumCloseToTarget(e.input, e.target)
		assert.Equal(e.expectedResult, actualResult)
	}
}
