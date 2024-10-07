package main

import (
	"fmt"
	"strconv"
	"sum/linkedList"
)

func reverse(curr *linkedList.Node) *linkedList.Node {
	var prev, next *linkedList.Node
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	l1 := linkedList.NewLinkedList()
	l2 := linkedList.NewLinkedList()

	l1.Add(2)
	l1.Add(4)
	l1.Add(3)

	l2.Add(5)
	l2.Add(6)
	l2.Add(4)

	l1.Print()
	fmt.Println()
	l2.Print()
	fmt.Println()

	r := linkedList.NewLinkedList()
	c := 0

	cur1, cur2 := l1.Head, l2.Head
	for cur1 != nil || cur2 != nil {
		s := c

		if cur1 != nil {
			s = s + cur1.Data
			cur1 = cur1.Next
		}

		if cur2 != nil {
			s = s + cur2.Data
			cur2 = cur2.Next
		}

		c = s / 10
		r.Add(s % 10)
	}

	reversed := reverse(r.Head)
	curr := reversed
	resultSum := ""
	for curr != nil {
		resultSum = resultSum + strconv.Itoa(curr.Data)
		curr = curr.Next
	}

	fmt.Println()
	fmt.Println(resultSum)
}
