package main

import "fmt"

func mergeMaps(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)

	for k, v := range m1 {
		result[k] = v
	}

	for k, v := range m2 {
		result[k] += v
	}

	return result
}

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}

	fmt.Println("Вход:")
	fmt.Println("m1 :=", m1)
	fmt.Println("m2 :=", m2)

	merged := mergeMaps(m1, m2)

	fmt.Println("Выход:", merged)
}
