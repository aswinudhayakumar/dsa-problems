package main

import "fmt"

// most optimized
func main() {
	a := []int{1, 2, -3, 4, -5}
	n := 5
	max := 0

	sum := 0
	for i := 0; i < n; i++ {
		sum += a[i]
		if sum > max {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	fmt.Println(max)
}
