package main

import "fmt"

type Queue struct {
	Data []int
	Size int
}

func NewQueue() *Queue {
	return &Queue{
		Size: 0,
	}
}

func (q *Queue) EnQueue(k int) {
	q.Data = append(q.Data, k)
	q.Size++
}

func (q *Queue) DeQueue() int {
	r := q.Data[0]
	q.Data = q.Data[1:]
	q.Size--
	return r
}

func (q *Queue) Peek() int {
	if q.Size != 0 {
		return q.Data[0]
	}
	return -1
}

func main() {
	q := NewQueue()

	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)

	fmt.Println(q.Data)

	fmt.Println(q.Peek())
	q.DeQueue()

	fmt.Println(q.Peek())
	q.DeQueue()

	fmt.Println(q.Data)
}
