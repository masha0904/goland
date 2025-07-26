package main

import "fmt"

func main() {
	var a, b int32
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	sum := a + b
	fmt.Printf("Сумма: %d", sum)
}
