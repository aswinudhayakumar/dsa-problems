package main

import "fmt"

func main() {
	a := []int{4, 1, 2, 1, 2}

	xor := 0
	for i := 0; i < len(a); i++ {
		xor ^= a[i]
	}

	fmt.Println(xor)
}
