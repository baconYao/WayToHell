package meetingroomii

import (
	"container/heap"
	"sort"
)

type MinHeap []int

func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findSets(meetings [][]int) int {
	if len(meetings) == 0 {
		return 0
	}

	// Meetings are sorted according to their start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	// Initialize a new heap and add the ending time of the first meeting to the heap
	minHeap := newMinHeap()
	minHeap.Push(meetings[0][1]) // insert the first meeting's endtime

	for i := 1; i < len(meetings); i++ {
		start, end := meetings[i][0], meetings[i][1]
		// Check if the minimum element of the heap (i.e., the earliest ending meeting) is free
		if minHeap.Len() > 0 && start >= (*minHeap)[0] {
			// If the room is free, extract the earliest ending meeting and add the ending time of the current meeting
			heap.Pop(minHeap)
		}
		// Add the ending time of the current meeting to the heap
		heap.Push(minHeap, end)
	}

	// The size of the heap tells us the number of rooms allocated
	return minHeap.Len()
}
