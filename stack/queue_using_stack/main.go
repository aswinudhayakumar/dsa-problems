package main

import (
	"fmt"
	"queueUsingStack/stack"
)

type Queue struct {
	Input  stack.Stack
	Output stack.Stack
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) push(k int) {
	q.Input.Push(k)
}

func (q *Queue) peek() int {
	if q.Output.Size() == 0 {
		for i := 0; q.Input.Size()-1 >= 0; i++ {
			q.Output.Push(q.Input.Pop())
		}
	}

	return q.Output.Peek()
}

func (q *Queue) pop() int {
	if q.Output.Size() == 0 {
		for i := 0; q.Input.Size()-1 >= 0; i++ {
			q.Output.Push(q.Input.Pop())
		}
	}

	return q.Output.Pop()
}

func main() {
	q := NewQueue()

	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)

	fmt.Println(q)
	fmt.Println(q.peek())

	q.pop()
	q.pop()
	q.pop()
	fmt.Println(q.peek(), q)

	q.push(5)
	q.push(6)
	q.push(7)
	fmt.Println(q.peek(), q)

	q.pop()
	q.pop()
	q.pop()
	fmt.Println(q.peek(), q)
}
