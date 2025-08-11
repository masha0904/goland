package main

import (
	"fmt"
)

func main() {
	input := "3"
	var num int

	_, err := fmt.Sscan(input, &num)
	if err != nil {
		fmt.Println("Ошибка парсинга:", err)
		return
	}

	minVal := 0
	maxVal := 255

	if num < minVal || num > maxVal {
		fmt.Println("Число выходит за диапазон:", minVal, "-", maxVal)
		return
	}

	fmt.Println("Число успешно преобразовано:", num)
}
