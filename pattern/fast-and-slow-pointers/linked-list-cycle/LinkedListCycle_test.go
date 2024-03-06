package linkedlistcycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCycle(t *testing.T) {
	assert := assert.New(t)

	var theTests = []struct {
		input          []int // an array representing the contents of a linked list
		input2         int   // index of the node to which the tail pointer is pointing
		expectedResult bool
	}{
		{
			[]int{2, 4, 6, 8, 10},
			2,
			true,
		},
		{
			[]int{1, 3, 5, 7, 9},
			-1,
			false,
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			true,
		},
		{
			[]int{0, 2, 3, 5, 6},
			-1,
			false,
		},
		{
			[]int{4, 4, 4, 4, 4, 4},
			-1,
			false,
		},
		{
			[]int{3, 6, 8, 9, 10},
			0,
			true,
		},
	}

	for _, e := range theTests {
		inputLinkedList := &LinkedList{}
		inputLinkedList.CreateLinkedList(e.input)
		if e.input2 != -1 {
			length := inputLinkedList.GetLength(inputLinkedList.head)
			lastNode := inputLinkedList.GetNode(inputLinkedList.head, length-1)
			lastNode.next = inputLinkedList.GetNode(inputLinkedList.head, e.input2)
		}
		actualResult := DetectCycle(inputLinkedList.head)
		assert.Equal(e.expectedResult, actualResult)
	}
}
