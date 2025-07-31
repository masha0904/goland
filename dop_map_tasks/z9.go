package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 4, 5, 4}
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v > 1 {
			fmt.Println("Дубликат:", k)
		}

	}
}
