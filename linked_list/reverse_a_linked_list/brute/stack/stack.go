package stack

type Stack struct {
	data []int
	size int
}

func NewStack() *Stack {
	return &Stack{
		size: 0,
	}
}

func (s *Stack) Push(element int) {
	s.data = append(s.data, element)
	s.size++
}

func (s *Stack) Peek() int {
	return s.data[s.size-1]
}

func (s *Stack) Pop() int {
	r := s.Peek()
	s.data = s.data[:s.size-1]
	s.size--
	return r
}
