package main

import (
	"fmt"
	"removekdigits/stack"
)

// reverseString function reverses a given string
func reverseString(s string) string {
	// Convert string to a slice of runes to handle Unicode characters
	runes := []rune(s)

	// Reverse the slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string
	return string(runes)
}

func main() {
	a := "123456"
	k := 3
	s := stack.NewStack()
	r := ""

	if k == len(a) {
		fmt.Println(0)
		return
	}

	for _, rv := range a {
		v := int(rv)
		if s.Size() == 0 {
			s.Push(v)
		} else if v >= s.Peek() {
			s.Push(v)
		} else {
			if k > 0 {
				s.Pop()
				k = k - 1
				s.Push(v)
			}
		}
	}

	for k > 0 {
		s.Pop()
		k = k - 1
	}

	for s.Size() != 0 {
		r = r + string(s.Pop())
	}

	fmt.Println(reverseString(r))
}
