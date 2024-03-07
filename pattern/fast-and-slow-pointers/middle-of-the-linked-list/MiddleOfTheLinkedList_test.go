package middleofthelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMiddleNode(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          []int // an array representing the contents of a linked list
		expectedResult int
	}{
		{
			[]int{1, 2, 3, 4, 5}, 3,
		},
		{
			[]int{1, 2, 3, 4, 5, 6}, 4,
		},
		{
			[]int{3, 2, 1}, 2,
		},
		{
			[]int{10}, 10,
		},
		{
			[]int{98, 99}, 99,
		},
	}

	for _, e := range theTests {
		inputLinkedList := &LinkedList{}
		inputLinkedList.CreateLinkedList(e.input)
		actualResult := GetMiddleNode(inputLinkedList.head)
		assert.Equal(e.expectedResult, actualResult.data)
	}
}
