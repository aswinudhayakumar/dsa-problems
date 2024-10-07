package main

import "fmt"

type stack struct {
	items []int
}

func NewStack() stack {
	return stack{}
}

func (s *stack) push(k int) {
	s.items = append(s.items, k)
}

func (s *stack) pop() {
	n := len(s.items)
	s.items = s.items[0 : n-1]
}

func (s *stack) peek() int {
	return s.items[len(s.items)-1]
}

func (s *stack) size() int {
	return len(s.items)
}

func main() {
	s := NewStack()
	s.push(1)
	s.push(2)

	fmt.Println(s.items)
	s.pop()
	s.pop()
	fmt.Println(s.items)
}
