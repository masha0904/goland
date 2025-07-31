package main

import "fmt"

func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0)

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

func main() {
	input := []string{"a", "b", "a", "c"}

	fmt.Println("Вход:", input)

	output := removeDuplicates(input)

	fmt.Println("Выход:", output)
}
