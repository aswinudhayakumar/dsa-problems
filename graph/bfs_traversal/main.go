package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
}

func (g *Graph) BFSTraversal(start int) []int {
	res := []int{}
	visited := make(map[int]bool)
	visited[0] = true

	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		element := queue.Front()
		current := element.Value.(int)
		queue.Remove(element)

		res = append(res, current)

		for _, data := range g.adjList[current] {
			if !visited[data] {
				queue.PushBack(data)
				visited[data] = true
			}
		}
	}

	return res
}

func main() {
	g := NewGraph()

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	// output 0, 1, 2, 3, 4
	fmt.Println(g.BFSTraversal(0))
}
