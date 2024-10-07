package main

import "fmt"

func main() {
	k := 6
	n := 4
	a := []int{3, 1, 2, 4}
	hashMap := make(map[int]int)
	hashMap[0] = 1

	preSum, count := 0, 0
	for i := 0; i < n; i++ {
		preSum = preSum + a[i]
		hashMap[preSum] = hashMap[preSum] + 1

		remove := preSum - k
		if hashMap[remove] != 0 {
			hashMap[remove] = hashMap[remove] + 1
			count = count + 1
		}
	}

	fmt.Println(count)
}
