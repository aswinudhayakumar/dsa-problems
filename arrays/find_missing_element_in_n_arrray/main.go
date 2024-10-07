package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 5}
	n := 5

	sumOfArray := 0
	for i := 0; i < len(a); i++ {
		sumOfArray = sumOfArray + a[i]
	}
	sumOfN := (n * (n + 1)) / 2

	fmt.Println(sumOfN - sumOfArray)
}
