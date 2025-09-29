package stack

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int
}

// Size() function will return the size of the linked list.
func (s *StackLinkedList) Size() int {
	return s.size
}

/*
IsEmpty() function will return true is size of the linked list is
equal to zero or false in all other cases.
*/
func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

/*
First, the Peek() function will check if the stack is empty.
If not, it will return the peek value of stack i.e., will return the
head value of the linked list.
*/
func (s *StackLinkedList) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, true //Return 0,true if stack is empty
	}
	return s.head.value, false
}

// Push() function  will add new value at the start of the linked list.
func (s *StackLinkedList) Push(value int) {
	n := new(Node)
	n.value = value
	n.next = s.head
	s.head = n
	s.size += 1
}

/*
In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the linked list and return it.
*/
func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, true //Return 0,true if stack is empty
	}

	v := s.head.value
	s.head = s.head.next
	s.size -= 1
	return v, false
}

/* Print() function will print the elements of the linked list. */
func (s *StackLinkedList) Print() {
	for current := s.head; current.next != nil; current = current.next {
		fmt.Print(current.value, " ")
	}
	fmt.Println()
}
