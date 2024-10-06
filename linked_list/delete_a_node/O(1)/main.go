package main

import "delete_o_1/linkedList"

func main() {
	ll := linkedList.NewLinkedList()
	a := []int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		ll.Add(a[i])
	}

	n := ll.GetNode(1)
	if n != nil {
		ll.DeleteNode(n)
	}

	ll.Print()
}
