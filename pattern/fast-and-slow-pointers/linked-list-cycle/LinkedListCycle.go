package linkedlistcycle

import (
	"fmt"
)

func DetectCycle(head *LinkedListNode) bool {
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
