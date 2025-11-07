package minimumcosttoconnectsticks

import (
	"container/heap"
)

type MinHeap []int

func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j] // Min Heap is based on integer value
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Len() int {
	return len(h)
}
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

func connectSticks(sticks []int) int {
	if len(sticks) == 1 {
		return 0
	}

	pq := newMinHeap()

	for _, v := range sticks {
		heap.Push(pq, v)
	}

	cost := 0

	for pq.Len() > 1 {
		top := heap.Pop(pq).(int)
		second := heap.Pop(pq).(int)
		cost = cost + top + second
		heap.Push(pq, top+second)
	}

	return cost
}
