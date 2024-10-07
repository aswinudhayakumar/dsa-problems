package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 6, 4, 5}

	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			fmt.Println("false")
			return
		}
	}

	fmt.Println("true")
}
