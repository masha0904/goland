package main

import "fmt"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
func main() {
	result, err := Divide(1.0, 0)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	fmt.Println("Результат:", result)
}
