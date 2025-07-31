package main

import "fmt"

func umnozenieSlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}
func main() {
	numbers := []int{1, 2, 3}
	fmt.Println("Вход:", numbers)
	umnozenieSlice(numbers)
	fmt.Println("Выход:", numbers)
}
