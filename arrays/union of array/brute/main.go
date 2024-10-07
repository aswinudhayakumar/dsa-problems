package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{2, 3, 4, 5, 11, 12}
	freq := make(map[int]int)

	for i := 0; i < len(a); i++ {
		freq[a[i]] = freq[a[i]] + 1
	}

	for i := 0; i < len(b); i++ {
		freq[b[i]] = freq[b[i]] + 1
	}

	u := []int{}
	for k, _ := range freq {
		u = append(u, k)
	}

	sort.Ints(u)
	fmt.Println(u)
}
