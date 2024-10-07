package main

import "fmt"

func main() {
	a := []int{1, 1, 2, 2, 2, 3, 3, 10, 10, 10, 12, 12, 12}
	i := 0
	for j := 1; j < len(a); j++ {
		if a[i] != a[j] {
			i = i + 1
			a[i] = a[j]
		}
	}

	updatedLength := i + 1

	for j := 0; j < updatedLength; j++ {
		fmt.Println(a[j])
	}
}
