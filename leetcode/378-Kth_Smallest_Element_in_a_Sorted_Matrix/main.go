package kthsmallestelementinasortedmatrix

import (
	"container/heap"
)

type Set struct {
	value        int // the current value
	idxOfLists   int // index of lists
	idxOfElement int // index of element in a list
}

type MinHeap []Set

// newMinHeap intializes an instance of MinHeap
func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

// Len function returns the length of MinHeap
func (h MinHeap) Len() int {
	return len(h)
}

// Empty function returns true of empty, false otherwise
func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares the two elements of MinHeap given their indices
func (h MinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

// Swap function swaps the values of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MinHeap
func (h MinHeap) Top() interface{} {
	return h[0]
}

// Push function pushes the given element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}

// Pop function pops the top element of MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	minHeap := newMinHeap()
	ascendingList := make([]int, 0)

	for i := 0; i < len(matrix); i++ {
		heap.Push(minHeap, Set{
			value:        matrix[i][0],
			idxOfLists:   i,
			idxOfElement: 0,
		})
	}

	for !minHeap.Empty() && len(ascendingList) != k {
		element := heap.Pop(minHeap).(Set)
		ascendingList = append(ascendingList, element.value)

		if len(matrix[element.idxOfLists])-1 > element.idxOfElement {
			heap.Push(minHeap, Set{
				value:        matrix[element.idxOfLists][element.idxOfElement+1],
				idxOfLists:   element.idxOfLists,
				idxOfElement: element.idxOfElement + 1,
			})
		}
	}

	return ascendingList[k-1]
}
