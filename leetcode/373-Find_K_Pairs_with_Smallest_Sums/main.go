package findkpairswithsmallestsums

import (
	"container/heap"
)

type Set struct {
	sum    int // pair 的加總 (nums1[index1] + nums2[index2])
	index1 int // index of nums1
	index2 int // index of nums2
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
	return h[i].sum < h[j].sum
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

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, 0)
	minHeap := newMinHeap()
	counter := 1

	for i, _ := range nums1 {
		heap.Push(minHeap, Set{sum: nums1[i] + nums2[0], index1: i, index2: 0})
	}

	for !minHeap.Empty() && counter <= k {
		popElement := heap.Pop(minHeap).(Set)
		idx1, idx2 := popElement.index1, popElement.index2
		result = append(result, []int{nums1[idx1], nums2[idx2]})

		nextElement := idx2 + 1
		if len(nums2) > nextElement {
			heap.Push(minHeap, Set{sum: nums1[idx1] + nums2[nextElement], index1: idx1, index2: nextElement})
		}
		counter += 1
	}
	return result
}

func kSmallestPairsBest(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, &Set{
			index1: i,
			index2: 0,
			sum:    nums1[i] + nums2[0],
		})
	}

	resultList := make([][]int, k)
	for i := 0; i < k; i++ {
		node := (heap.Pop(h)).(*Set)
		resultList[i] = []int{nums1[node.index1], nums2[node.index2]}
		if node.index2+1 < len(nums2) {
			heap.Push(h, &Set{
				index1: node.index1,
				index2: node.index2 + 1,
				sum:    nums1[node.index1] + nums2[node.index2+1],
			})
		}
	}

	return resultList
}
