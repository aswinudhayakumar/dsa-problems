package main

import (
	"fmt"
	"optimal1/linkedList"
)

func getLengthOfLinkedList(curr *linkedList.Node) int {
	len := 0
	for curr != nil {
		len++
		curr = curr.Next
	}
	return len
}

func main() {
	l1 := linkedList.NewLinkedList()
	a := []int{1, 3, 1, 2, 4}

	for i := 0; i < 5; i++ {
		l1.Add(a[i])
	}

	fmt.Println("input 1...")
	l1.Print()

	l2 := linkedList.NewLinkedList()
	t := l1.Head.Next.Next.Next
	l2.Head = &linkedList.Node{
		Data: 3,
		Next: t,
	}

	fmt.Println("input 2...")
	l2.Print()

	curr := l1.Head
	l1s := getLengthOfLinkedList(curr)

	curr = l2.Head
	l2s := getLengthOfLinkedList(curr)

	// find difference and move the higher sized list pointer to the size of lower sized list pointer
	d1 := l1.Head
	d2 := l2.Head
	if l1s != l2s {
		if l1s > l2s {
			diff := l1s - l2s
			for i := 0; i < diff; i++ {
				d1 = d1.Next
			}
		}
		if l2s > l1s {
			diff := l2s - l1s
			for i := 0; i < diff; i++ {
				d2 = d2.Next
			}
		}
	}

	// intersection is present if d1 and d2 matches
	for d1 != nil {
		if d1 == d2 {
			fmt.Println("intersection - ", d1.Data)
			return
		}
		d1 = d1.Next
		d2 = d2.Next
	}
}
