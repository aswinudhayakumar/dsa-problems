package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	k := 3

	for i := 0; i < len(a); i++ {
		if a[i] == k {
			fmt.Println(i)
			return
		}
	}

	fmt.Println("-1")
}
