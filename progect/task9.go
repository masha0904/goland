package main

import "fmt"

func main() {
	n := 15
	a := 0
	b := 1
	fmt.Println("Числа Фибоначи:")
	for i := 0; i < n; i++ {
		fmt.Println(a, " ")
		a, b = b, a+b
	}
}
