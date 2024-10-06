package main

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

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func main() {
	head := NewLinkedList()

	for i := 1; i <= 10; i++ {
		head.Add(i)
	}

	fmt.Println("size --> ", head.Size)
	head.Print()
}
