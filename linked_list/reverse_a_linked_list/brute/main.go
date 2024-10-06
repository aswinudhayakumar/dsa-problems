package main

import (
	"fmt"
	"reverse_brute/linkedList"
	"reverse_brute/stack"
)

func main() {
	ll := linkedList.NewLinkedList()
	a := []int{1, 3, 2, 4}
	st := stack.NewStack()

	for i := 0; i < 4; i++ {
		ll.Add(a[i])
	}

	fmt.Println("input...")
	ll.Print()

	current := ll.Head
	for current != nil {
		st.Push(current.Data)
		current = current.Next
	}

	current = ll.Head
	for current != nil {
		current.Data = st.Pop()
		current = current.Next
	}

	fmt.Println()
	fmt.Println("Reversed...")
	ll.Print()
}
