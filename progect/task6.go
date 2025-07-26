package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 2, 3, 7}
	maximum := nums[0]
	for _, v := range nums {
		if v > maximum {
			maximum = v
		}
	}
	fmt.Println("Максимальный элемент: ", maximum)
}
