package stack

type Stack struct {
	Items []int
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(k int) {
	s.Items = append(s.Items, k)
}

func (s *Stack) Pop() int {
	n := len(s.Items)
	poped := s.Items[n-1]
	s.Items = s.Items[0 : n-1]
	return poped
}

func (s *Stack) Peek() int {
	return s.Items[len(s.Items)-1]
}

func (s *Stack) Size() int {
	return len(s.Items)
}
