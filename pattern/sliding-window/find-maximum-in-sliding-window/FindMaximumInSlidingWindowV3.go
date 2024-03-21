package findmaximuminslidingwindow

import "container/list"

func FindMaxSlidingWindowV3(nums []int, w int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	if w > n {
		w = n
	}

	output := make([]int, 0)
	currentWindow := NewDeque()

	// Iterate over the first w elemnts
	for i := 0; i < w; i++ {
		cleanUp3(i, currentWindow, nums)
		currentWindow.PushBack(i)
	}

	// Appending the maximum element of the current window
	output = append(output, nums[currentWindow.Front()])

	// Iterate over the remaining
	for i := w; i < n; i++ {
		cleanUp3(i, currentWindow, nums)
		// Remove first index from the currentWindow if it has fallen out of the current window
		if currentWindow.Len() > 0 && currentWindow.Front() <= i-w {
			currentWindow.PopFront()
		}
		currentWindow.PushBack(i)
		output = append(output, nums[currentWindow.Front()])
	}

	return output
}

// clenaUp function cleans up the window
func cleanUp3(i int, currentWindow *Deque, nums []int) {
	for currentWindow.Len() > 0 && nums[i] >= nums[currentWindow.Back()] {
		// Remove all the indexes from currentWindow whose corresponding values are
		// smaller than or equal to the current element --> nums[i]
		currentWindow.PopBack()
	}
}

type Deque struct {
	items *list.List
}

// NewDeque is a constructor that will declare and return the Deque type object
func NewDeque() *Deque {
	return &Deque{list.New()}
}

// PushFront will push an element at the front of the dequeue
func (d *Deque) PushFront(value int) {
	d.items.PushFront(value)
}

// PushBack will push an element at the back of the dequeue
func (d *Deque) PushBack(value int) {
	d.items.PushBack(value)
}

// PopFront will pop an element from the front of the dequeue
func (d *Deque) PopFront() int {
	return d.items.Remove(d.items.Front()).(int)
}

// PopBack will pop an element from the back of the dequeue
func (d *Deque) PopBack() int {
	return d.items.Remove(d.items.Back()).(int)
}

// Front will return the element from the front of the dequeue
func (d *Deque) Front() int {
	return d.items.Front().Value.(int)
}

// Back will return the element from the back of the dequeue
func (d *Deque) Back() int {
	return d.items.Back().Value.(int)
}

// Empty will check if the dequeue is empty or not
func (d *Deque) Empty() bool {
	return d.items.Len() == 0
}

// Len will return the length of the dequeue
func (d *Deque) Len() int {
	return d.items.Len()
}
