package main

import (
	"detect/linkedList"
	"fmt"
)

func detectLoop(head *linkedList.Node) bool {
	slow := head.Next
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	ll := linkedList.NewLinkedList()
	a := []int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		ll.Add(a[i])
	}

	var t *linkedList.Node
	current := ll.Head
	for current.Next != nil {
		if current.Data == 3 {
			t = current
		}
		current = current.Next
	}
	current.Next = t

	if detectLoop(ll.Head) {
		fmt.Println("loop detected....")
		return
	}
	fmt.Println("no loop detected....")
}
