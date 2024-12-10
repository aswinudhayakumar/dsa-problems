package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 15
	n := 10
	max := 0

	for i := 0; i < n; i++ {
		sum := a[i]
		for j := i + 1; j < n; j++ {
			if a[j] == 0 {
				continue
			}
			sum = sum + a[j]
			if sum == k {
				size := (j - i) + 1
				if max < size {
					max = size
				}
			}
		}
	}

	fmt.Println(max)
}
