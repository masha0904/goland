package main

import "fmt"

func plusSlices(a, b []int) []int {
	result := make([]int, len(a)+len(b))
	copy(result, a)
	copy(result[len(a):], b)
	return result
}

func main() {
	a := []int{1, 2}
	b := []int{3, 4}

	fmt.Println("Вход:", a, b)

	combined := plusSlices(a, b)
	fmt.Println("Выход:", combined)
}
