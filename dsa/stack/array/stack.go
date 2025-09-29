package stack

import "fmt"

type StackInt struct {
	s []int
}

// isEmpty() function returns true if stack is empty or false in all other cases.
func (s *StackInt) IsEmpty() bool {
	return len(s.s) == 0
}

// length() function returns the number of elements in the stack.
func (s *StackInt) Length() int {
	return len(s.s)
}

// The print function will print the elements of the array.
func (s *StackInt) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

// push() function append value to the data.
func (s *StackInt) Push(value int) {
	s.s = append(s.s, value)
}

/* In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the data array and return it. */

func (s *StackInt) Pop() int {
	if s.IsEmpty() {
		fmt.Print("Stack is empty.")
		return 0
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]

	return res
}

/*
top() function will first check that the stack is not empty
then returns the value stored in the top element
of stack (does not remove it).
*/
func (s *StackInt) Top() int {
	if s.IsEmpty() {
		fmt.Print("Stack is empty.")
		return 0
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}
