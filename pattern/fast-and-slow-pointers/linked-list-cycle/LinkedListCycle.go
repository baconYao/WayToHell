package linkedlistcycle

import (
	"fmt"
	"strings"
)

func DetectCycle(head *LinkedListNode) bool {
	slowPtr, fastPtr := head, head.next
	for fastPtr != nil && fastPtr.next != nil {
		if slowPtr == fastPtr {
			return true
		}
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
	}
	return false
}

type LinkedListNode struct {
	data int
	next *LinkedListNode
}

func NewLinkedListNode(data int, next *LinkedListNode) *LinkedListNode {
	node := new(LinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *LinkedListNode {
	node := new(LinkedListNode)
	node.data = data
	node.next = nil
	return node
}

type LinkedList struct {
	head *LinkedListNode
}

/*
InsertNodeAtHead method will insert a LinkedListNode at head of a linked list.
*/
func (l *LinkedList) InsertNodeAtHead(node *LinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

/*
CreateLinkedList method will create the linked list using
the given integer array with the help of InsertAthead method.
*/
func (l *LinkedList) CreateLinkedList(lst []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		newNode := InitLinkedListNode(lst[i])
		l.InsertNodeAtHead(newNode)
	}
}

// GetLength returns the number of nodes in the linked list
func (ll *LinkedList) GetLength(head *LinkedListNode) int {
	temp := head
	length := 0
	for temp != nil {
		length++
		temp = temp.next
	}
	return length
}

// GetNode returns the node at the specified position (index) of the linked list
func (ll *LinkedList) GetNode(head *LinkedListNode, pos int) *LinkedListNode {
	if pos != -1 {
		p := 0
		ptr := head
		for p < pos {
			ptr = ptr.next
			p++
		}
		return ptr
	}
	return nil
}

// DisplayLinkedList method will display the elements of linked list.
func (l *LinkedList) DisplayLinkedList() {
	temp := l.head
	fmt.Print("[")
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}

/*
	DisplayLinkedListWithForwardArrow method will display the linked list

not in the form of an array, but rather a list with arrows pointing to
the next element
*/
func DisplayLinkedListWithForwardArrow(l *LinkedListNode) {
	temp := l
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(" → ")
		} else {
			fmt.Print(" → null")
		}
	}
}
func DisplayLinkedListWithForwardArrowLoop(l *LinkedListNode) {
	temp := l
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(" → ")
		}
	}
}

func show() {
	inputs := [][]int{
		{2, 4, 6, 8, 10, 12},
		{1, 3, 5, 7, 9, 11},
		{0, 1, 2, 3, 4, 6},
		{3, 4, 7, 9, 11, 17},
		{5, 1, 4, 9, 2, 3},
		{4, 4, 4, 4, 4, 4},
	}
	pos := []int{0, -1, 1, -1, 2, -1}
	j := 1

	for i := range inputs {
		inputLinkedList := &LinkedList{}
		inputLinkedList.CreateLinkedList(inputs[i])
		fmt.Printf("%d.\tInput: ", j)
		if pos[i] == -1 {
			DisplayLinkedListWithForwardArrow(inputLinkedList.head)
		} else {
			DisplayLinkedListWithForwardArrowLoop(inputLinkedList.head)
		}
		fmt.Println("\n\tpos:", pos[i])
		if pos[i] != -1 {
			length := inputLinkedList.GetLength(inputLinkedList.head)
			lastNode := inputLinkedList.GetNode(inputLinkedList.head, length-1)
			lastNode.next = inputLinkedList.GetNode(inputLinkedList.head, pos[i])
		}

		fmt.Printf("\n\tDetected cycle = %t\n", DetectCycle(inputLinkedList.head))
		j++
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
