package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{2, 3, 4, 5, 11, 12}
	u := []int{}

	n := len(a)
	m := len(b)
	i := 0
	j := 0

	for i < n && j <= m {
		if a[i] < b[j] {
			if len(u) == 0 || u[len(u)-1] != a[i] {
				u = append(u, a[i])
			}
			i++
		} else {
			if len(u) == 0 || u[len(u)-1] != b[j] {
				u = append(u, b[j])
			}
			j++
		}
	}

	for i < n {
		if u[len(u)-1] != a[i] {
			u = append(u, a[i])
		}
		i++
	}

	for j < m {
		if u[len(u)-1] != b[j] {
			u = append(u, b[j])
		}
		j++
	}

	fmt.Println(u)
}
