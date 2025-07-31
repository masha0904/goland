package main

import "fmt"

func mapToSlices(m map[int]string) ([]int, []string) {
	keys := make([]int, 0, len(m))
	values := make([]string, 0, len(m))

	for key, value := range m {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

func main() {
	input := map[int]string{1: "a", 2: "b"}
	fmt.Println("Вход:", input)

	keySlice, valueSlice := mapToSlices(input)
	fmt.Printf("Выход: %#v, %#v\n", keySlice, valueSlice)
}
