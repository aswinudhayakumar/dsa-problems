package main

import "delete/linkedList"

func main() {
	ll := linkedList.NewLinkedList()
	a := []int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		ll.Add(a[i])
	}

	ll.DeleteElement(4)
	ll.Print()
}
