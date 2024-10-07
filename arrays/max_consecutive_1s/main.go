package main

import "fmt"

func main() {
	a := []int{1, 1, 0, 1, 1, 1}
	max := 0
	sum := 0

	for i := 0; i < len(a); i++ {
		if a[i] == 1 {
			sum = sum + 1
		} else {
			sum = 0
		}
		max = sum
	}

	fmt.Println(max)
}
