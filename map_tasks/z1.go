package main

import "fmt"

func countelements(items []string) map[string]int {
	m := make(map[string]int)
	for _, item := range items {
		m[item]++
	}
	return m
}
func main() {
	items := []string{"a", "b", "c", "a", "b", "c", "a", "b"}
	fmt.Println("Вход:", items)
	result := countelements(items)
	fmt.Println("Выход:", result)

}
