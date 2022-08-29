package implementqueueusingstacksgo

type MyQueue struct {
	pOrder Stack // Positive order
	rOrder Stack // Reverse order
}

func Constructor() MyQueue {
	return *new(MyQueue)
}

func (this *MyQueue) Push(x int) {
	this.pOrder.Push(x)
}

func (this *MyQueue) Pop() int {
	// Put element from pOrder to rOrder
	for {
		pv, ok := this.pOrder.Pop()
		if !ok {
			break
		}
		this.rOrder.Push(pv)
	}
	value, _ := this.rOrder.Pop()
	// Put element back to pOrder from rOrder
	for {
		rv, ok := this.rOrder.Pop()
		if !ok {
			break
		}
		this.pOrder.Push(rv)
	}
	return value
}

func (this *MyQueue) Peek() int {
	// Put element from pOrder to rOrder
	for {
		pv, ok := this.pOrder.Pop()
		if !ok {
			break
		}
		this.rOrder.Push(pv)
	}
	value, _ := this.rOrder.Peek()
	// Put element back to pOrder from rOrder
	for {
		rv, ok := this.rOrder.Pop()
		if !ok {
			break
		}
		this.pOrder.Push(rv)
	}
	return value
}

func (this *MyQueue) Empty() bool {
	return this.pOrder.IsEmpty()
}

type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:index]
	return elem, true
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	elem := (*s)[index]
	return elem, true
}

// ====================================================
// 非 Stack 的解法
type MyQueue2 struct {
	Queue []int
}

func Constructor2() MyQueue2 {
	return MyQueue2{}
}

func (this *MyQueue2) Push(x int) {
	this.Queue = append(this.Queue, x)
}

func (this *MyQueue2) Pop() int {
	res := this.Queue[0]
	this.Queue = this.Queue[1:]
	return res
}

func (this *MyQueue2) Peek() int {
	return this.Queue[0]
}

func (this *MyQueue2) Empty() bool {
	return len(this.Queue) == 0
}
