package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var a int
	fmt.Print("Введите число: ")
	fmt.Scan(&a)
	fmt.Printf("Факториал: %d\n", factorial(a))
}
