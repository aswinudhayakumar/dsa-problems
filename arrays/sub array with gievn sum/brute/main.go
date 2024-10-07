package main

import "fmt"

func main() {
	// brute force - O(N)3
	n := 4
	a := []int{3,1,2,4}
	k := 6
	c := 0
	sum := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum = sumWithForLoop(a[i:j+1])
			if sum == k {
				fmt.Println(i, j)
				c = c + 1
			}
		}
		sum = 0
	}

	fmt.Println(c)

	// optimal
	
}

func sumWithForLoop(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
