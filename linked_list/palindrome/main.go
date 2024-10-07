package main

import (
	"fmt"
	"palindrome/linkedList"
)

func reverse(curr *linkedList.Node) *linkedList.Node {
	var prev, front *linkedList.Node
	for curr != nil {
		front = curr.Next
		curr.Next = prev
		prev = curr
		curr = front
	}

	return prev
}

func main() {
	ll := linkedList.NewLinkedList()
	a := []int{1, 2, 3, 2, 1}

	for i := 0; i < 5; i++ {
		ll.Add(a[i])
	}

	fmt.Println("input ...")
	ll.Print()
	fmt.Println()

	slow := ll.Head
	fast := ll.Head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	new_head := reverse(slow.Next)

	first := ll.Head
	second := new_head

	for second != nil {
		if first.Data != second.Data {
			fmt.Println("not palindrome")
			return
		}
		first = first.Next
		second = second.Next
	}

	reverse(slow.Next)
	fmt.Println("palindrome")
}
