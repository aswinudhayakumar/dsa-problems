package stack

type stack struct {
	Items []string
}

func NewStack() stack {
	return stack{}
}

func (s *stack) Push(k string) {
	s.Items = append(s.Items, k)
}

func (s *stack) Pop() string {
	n := len(s.Items)
	poped := s.Items[n-1]
	s.Items = s.Items[0 : n-1]
	return poped
}

func (s *stack) Peek() string {
	return s.Items[len(s.Items)-1]
}

func (s *stack) Size() int {
	return len(s.Items)
}
