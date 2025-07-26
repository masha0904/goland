package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 2, 3, 7}
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	fmt.Println("Сумма элементов: ", sum)
}
