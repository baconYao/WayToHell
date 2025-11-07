package longesthappystring

import (
	"container/heap"
)

type Pair struct {
	count int
	ch    rune
}

type MaxHeap []Pair

func newMaxHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count // Max Heap based on the count
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Len() int      { return len(h) }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func longestDiverseString(a int, b int, c int) string {
	pq := newMaxHeap()
	if a > 0 {
		heap.Push(pq, Pair{count: a, ch: 'a'})
	}
	if b > 0 {
		heap.Push(pq, Pair{count: b, ch: 'b'})
	}
	if c > 0 {
		heap.Push(pq, Pair{count: c, ch: 'c'})
	}

	output := ""

	for pq.Len() > 0 {
		top := heap.Pop(pq).(Pair)
		// The latest two chars are same as the top
		if len(output) >= 2 && output[len(output)-1] == byte(top.ch) && output[len(output)-2] == byte(top.ch) {
			if pq.Len() == 0 {
				break
			}
			second := heap.Pop(pq).(Pair)
			output = output + string(second.ch)
			// Push second back to heap
			if second.count-1 > 0 {
				heap.Push(pq, Pair{count: second.count - 1, ch: second.ch})
			}
			// Push top back to heap
			heap.Push(pq, top)
		} else {
			// Concatenate the value of top to output string
			output = output + string(top.ch)
			// Push top back to heap
			if top.count-1 > 0 {
				heap.Push(pq, Pair{count: top.count - 1, ch: top.ch})
			}
		}
	}

	return output
}
