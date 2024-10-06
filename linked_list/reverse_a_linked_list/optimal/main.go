package main

import (
	"fmt"
	"reverse_optimal/linkedList"
)

func main() {
	ll := linkedList.NewLinkedList()
	a := []int{1, 3, 2, 4}

	for i := 0; i < 4; i++ {
		ll.Add(a[i])
	}

	fmt.Println("input...")
	ll.Print()

	var pre, next *linkedList.Node
	temp := ll.Head

	for temp != nil {
		next = temp.Next
		temp.Next = pre
		pre = temp
		temp = next
	}

	ll.Head = pre
	ll.Print()
}
