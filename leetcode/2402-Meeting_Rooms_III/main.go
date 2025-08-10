package meetingroomsiii

import (
	"container/heap"
	"sort"
)

// MinHeap structure initialization
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of MinHeap given their indices
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap function swaps the value of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push function pushes an element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pair struct to hold the end time and room number
type Pair struct {
	endTime int64
	room    int
}

// PairHeap is a MinHeap of Pairs
type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}
func (h PairHeap) Empty() bool {
	return len(h) == 0
}

func (h PairHeap) Less(i, j int) bool {
	return h[i].endTime < h[j].endTime || (h[i].endTime == h[j].endTime && h[i].room < h[j].room)
}
func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mostBooked(meetings [][]int, rooms int) int {
	counter := make([]int, rooms) // The counter of each room be booked

	// Sort the meetings based on the first element (start) because each start is unique
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	// A min heap tracks the free room
	availableRooms := &MinHeap{}
	heap.Init(availableRooms)

	for r := 0; r < rooms; r++ {
		heap.Push(availableRooms, r)
	}

	//  A min heap tracks the endtime of meetings and its room number
	usedRooms := &PairHeap{}
	heap.Init(usedRooms)

	for i := 0; i < len(meetings); i++ {
		startTime := int64(meetings[i][0])
		endTime := int64(meetings[i][1])

		// Free the usedRooms if time's up
		for !usedRooms.Empty() && (*usedRooms)[0].endTime <= startTime {
			room := heap.Pop(usedRooms).(Pair).room
			heap.Push(availableRooms, room)
		}

		// No available room, need to pop the used room which ends the earliest
		if availableRooms.Len() == 0 {
			pair := heap.Pop(usedRooms).(Pair)
			// delay current meeting
			endTime = endTime + (pair.endTime - startTime)
			heap.Push(availableRooms, pair.room)
		}

		// Assign current meeting to the available room
		room := heap.Pop(availableRooms).(int)
		heap.Push(usedRooms, Pair{
			endTime: endTime,
			room:    room,
		})
		counter[room]++
	}

	maxUsedRoomIdx := 0
	for k, v := range counter {
		if v > counter[maxUsedRoomIdx] {
			maxUsedRoomIdx = k
		}
	}

	return maxUsedRoomIdx
}
