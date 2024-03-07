package middleofthelinkedlist

import "fmt"

func GetMiddleNode(head *LinkedListNode) *LinkedListNode {
	slowPtr, fastPtr := head, head
	for fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
	}
	return slowPtr
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

func TraverseLinkedList(head *LinkedListNode) {
	current := head
	next := new(LinkedListNode)
	next = nil
	for current != nil {
		next = current.next
		current = next
	}
}

func ReverseList(slowPtr *LinkedListNode) *LinkedListNode {
	reverse := new(LinkedListNode)
	reverse = nil
	for slowPtr != nil {
		next := slowPtr.next
		slowPtr.next = reverse
		reverse = slowPtr
		slowPtr = next
	}
	return reverse
}
