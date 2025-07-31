package main

import "fmt"

func in(slice []int, index, value int) []int {
	slice = slice[:len(slice)+1]

	copy(slice[index+1:], slice[index:])

	slice[index] = value

	return slice
}

func main() {
	slice := make([]int, 3, 5)
	slice[0], slice[1], slice[2] = 10, 20, 30

	fmt.Printf("До: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	slice = in(slice, 1, 99)

	fmt.Printf("После: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
}
