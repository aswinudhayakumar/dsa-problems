package linkedList

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Add(data int) {
	newNode := &Node{Data: data}
	if ll.Head == nil {
		ll.Head = newNode
		ll.Size = 1
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	ll.Size += 1
}

func (ll *LinkedList) DeleteElement(value int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Data == value {
		ll.Head = ll.Head.Next
		return
	}

	curr := ll.Head
	for curr.Next != nil {
		if curr.Next.Data == value {
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}
}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func (ll *LinkedList) Reverse() {
	var prev, next *Node
	current := ll.Head
	for current != nil {
		next = current.Next // store the next node
		current.Next = prev // reverse the current node's pointer
		prev = current      // move prev to current
		current = next      // move current to next
	}
	ll.Head = prev // update the head to the new front of the list
}
