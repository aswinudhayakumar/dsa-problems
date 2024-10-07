package main

import "math"

func main() {
	a := []int{10, 2, 4, 6, 1, 3}
	ss := secondSmall(a)
	sl := secondLarge(a)

	println(sl, ss)
}

func secondSmall(a []int) int {
	small, secondSmall := math.MaxInt64, math.MaxInt64

	for i := 0; i < len(a); i++ {
		if a[i] < small {
			secondSmall = small
			small = a[i]
		} else if a[i] > small && a[i] < secondSmall {
			secondSmall = a[i]
		}
	}

	return secondSmall
}

func secondLarge(a []int) int {
	large, secondLarge := math.MinInt64, math.MinInt64

	for i := 0; i < len(a); i++ {
		if a[i] > large {
			secondLarge = large
			large = a[i]
		} else if a[i] < large && a[i] > secondLarge{
			secondLarge = a[i]
		}
	}

	return secondLarge
}
