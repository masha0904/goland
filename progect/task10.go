package main

import "fmt"

func UniqueUserIDs(userIDs []int64) []int64 {
	seen := make(map[int64]bool)
	result := make([]int64, 0, len(userIDs))

	for _, id := range userIDs {
		if !seen[id] {
			seen[id] = true
			result = append(result, id)
		}
	}

	return result
}
func main() {
	fmt.Println(UniqueUserIDs([]int64{1, 2, 3, 5, 3, 3}))
}
