package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	t := a[0]
	for i := 0; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a[len(a)-2] = t

	fmt.Println(a)
}
