package main

import "fmt"

func filterMap(m map[string]int, n int) map[string]int {
	result := make(map[string]int)
	for key, value := range m {
		if value > n {
			result[key] = value
		}
	}
	return result
}

func main() {
	inputMap := map[string]int{"a": 10, "b": 5, "c": 20}
	filterValue := 15

	filtered := filterMap(inputMap, filterValue)

	fmt.Println(filtered)
}
