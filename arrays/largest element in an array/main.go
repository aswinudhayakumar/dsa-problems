package main

import "fmt"

func main() {
	a := []int{6, 5, 3, 7, 1, 3}
	max := a[0]

	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	fmt.Println(max)
}
