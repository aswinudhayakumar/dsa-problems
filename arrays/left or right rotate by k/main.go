package main

import "fmt"

func main() {
	// left rotation
	a := []int{1, 2, 3, 4, 5, 6, 7}
	k := 2
	n := 7

	// input
	fmt.Println(a)

	reverse(a[0:k])
	reverse(a[k:n])

	// output
	fmt.Println("left rotation")
	fmt.Println(reverse(a))

	fmt.Println()
	// --------------------------------

	// right rotation
	a = []int{1, 2, 3, 4, 5, 6, 7}

	// input
	fmt.Println(a)

	reverse(a[n-k : n])
	reverse(a[0 : n-k])

	// output
	fmt.Println("right rotation")
	fmt.Println(reverse(a))
}

func reverse(a []int) []int {
	left := 0
	right := len(a) - 1

	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}

	return a
}
