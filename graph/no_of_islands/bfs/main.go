package main

import (
	"container/list"
	"fmt"
)

type QueueData struct {
	row int
	col int
}

func bfs(n, m, row, col int, visited [][]bool, input [][]int) {
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	visited[row][col] = true
	queue := list.New()

	queue.PushBack(QueueData{row, col})

	for queue.Len() > 0 {
		element := queue.Front()
		current := element.Value.(QueueData)
		queue.Remove(element)

		for _, direction := range directions {
			r := current.row + direction[0]
			c := current.col + direction[1]

			if r >= 0 && r < n && c < m && c >= 0 && input[r][c] == 1 && !visited[r][c] {
				visited[r][c] = true
				queue.PushBack(QueueData{r, c})
			}
		}
	}
}

func findIslands(input [][]int) int {
	n := len(input)
	m := len(input[0])
	count := 0

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, len(input[0]))
	}

	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if !visited[row][col] && input[row][col] == 1 {
				bfs(n, m, row, col, visited, input)
				count++
			}
		}
	}

	return count
}

func main() {

	matrix := [][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{1, 1, 0, 1},
	}

	fmt.Println(findIslands(matrix))
}
