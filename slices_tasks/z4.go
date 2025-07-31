package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	fmt.Printf("Исходный: len = %d; cap = %d\n", len(s), cap(s))
	a := cap(s)
	for i := len(s); i <= a; i++ {
		s = append(s, i)
		fmt.Printf("После добавления:len= %d;cap= %d\n", len(s), cap(s))
	}
}
