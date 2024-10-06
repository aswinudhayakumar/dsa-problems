package main

import (
	"fmt"
	"optimal_recursive/linkedList"
)

func ReverseInRecursive(head *linkedList.Node) *linkedList.Node {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseInRecursive(head.Next)
	front := head.Next
	front.Next = head
	head.Next = nil

	return newHead
}

func main() {
	ll := linkedList.NewLinkedList()
	a := []int{1, 3, 2, 4}

	for i := 0; i < 4; i++ {
		ll.Add(a[i])
	}

	fmt.Println("input...")
	ll.Print()

	ll.Head = ReverseInRecursive(ll.Head)

	fmt.Println("output...")
	ll.Print()
}
