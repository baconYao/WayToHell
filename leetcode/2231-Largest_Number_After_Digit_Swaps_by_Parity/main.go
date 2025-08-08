package largestnumberafterdigitswapsbyparity

import (
	"container/heap"
	"slices" // Requires Go 1.21+
)

// Structure for MaxHeap
type MaxHeap []int

// newMaxHeap function intializes an instance of the MaxHeap
func newMaxHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

// Len function returns the length of the MaxHeap
func (h MaxHeap) Len() int {
	return len(h)
}

// Empty returns true if the MaxHeap is empty, false otherwise
func (h MaxHeap) Empty() bool {
	return len(h) == 0
}

// Greater returns true if the first of the given elements is greater than the second one
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap function swaps the values at the given indices
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MaxHeap
func (h MaxHeap) Top() int {
	return h[0]
}

// Push function inserts the element in the MaxHeap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MaxHeap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func largestInteger(num int) int {
	oddMaxHeap := newMaxHeap()
	evenMaxHeap := newMaxHeap()
	nums := make([]int, 0)
	ans := 0

	for num != 0 {
		remainder := num % 10
		num = num / 10
		if remainder%2 == 0 {
			heap.Push(evenMaxHeap, remainder)
		} else {
			heap.Push(oddMaxHeap, remainder)
		}
		nums = slices.Insert(nums, 0, remainder)
	}

	for _, v := range nums {
		ele := 0
		ans = ans * 10
		if v%2 == 0 {
			ele = heap.Pop(evenMaxHeap).(int)
		} else {
			ele = heap.Pop(oddMaxHeap).(int)
		}
		ans = ans + ele
	}

	return ans
}
