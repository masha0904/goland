package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Вход:", numbers) // [1 2 3 4 5]

	reverse(numbers)

	fmt.Println("Выход:", numbers) // [1 3 5]
}
