package main

import (
	"fmt"
	"optimal2/linkedList"
)

func main() {
	l1 := linkedList.NewLinkedList()
	a := []int{1, 3, 1, 2, 4}

	for i := 0; i < 5; i++ {
		l1.Add(a[i])
	}

	fmt.Println("input 1...")
	l1.Print()

	l2 := linkedList.NewLinkedList()
	
	for i := 0; i < 5; i++ {
		l2.Add(a[i])
	}
	l2.Add(6)
	l2.Add(7)

	fmt.Println("input 2...")
	l2.Print()

	d1 := l1.Head
	d2 := l2.Head

	for d1 != d2 {
		if d1 != nil {
			d1 = d1.Next
		} else {
			d1 = l2.Head
		}

		if d2 != nil {
			d2 = d2.Next
		} else {
			d2 = l1.Head
		}
	}

	fmt.Println()
	if d1 != nil {
		fmt.Println("res -> ", d1.Data)
	} else {
		fmt.Println("no intersection")
	}
}
