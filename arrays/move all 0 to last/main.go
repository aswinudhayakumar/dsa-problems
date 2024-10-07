package main

import "fmt"

func main() {
	a := []int{1, 0, 2, 3, 2, 0, 0, 4, 5, 1}
	j := -1

	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			j = i
			break
		}
	}

	if j == -1 {
		fmt.Println(a)
		return
	}

	for i := j + 1; i < len(a); i++ {
		if a[i] != 0 {
			a[i], a[j] = a[j], a[i]
			j++
		}
	}

	fmt.Println(a)
}
