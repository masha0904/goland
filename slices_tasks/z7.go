package main

import "fmt"

func deletechet(slice *[]int) {
	s := *slice
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j]%2 != 0 {
			s[i] = s[j]
			i++
		}
	}
	*slice = s[:i]
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("До фильтрации:", numbers) // [1 2 3 4 5]

	deletechet(&numbers)

	fmt.Println("После фильтрации:", numbers) // [1 3 5]
}
