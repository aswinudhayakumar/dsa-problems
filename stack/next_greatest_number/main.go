package main

import (
	"fmt"
	"nge/stack"
)

func main() {
	a := []int{4, 12, 5, 3, 1, 2, 5, 3, 1, 2, 4, 6}
	n := 12
	r := [12]int{}
	s := stack.NewStack()

	for j := ((2 * n) - 1); j >= 0; j-- {
		i := j % n
		if s.Size() == 0 {
			r[i] = -1
			s.Push(a[i])
		} else if a[i] < s.Peek() {
			r[i] = s.Peek()
			s.Push(a[i])
		} else {
			s.Pop()
			for s.Size() != 0 {
				if a[i] >= s.Peek() {
					s.Pop()
				} else {
					r[i] = s.Peek()
					s.Push(a[i])
					break
				}
			}
			if s.Size() == 0 {
				r[i] = -1
				s.Push(a[i])
			}
		}
	}

	fmt.Println(a)
	fmt.Println(r)
}
