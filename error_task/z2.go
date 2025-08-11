package main

import (
	"fmt"
	"math"
)

type NegativeNumberError struct {
	a float64
}

func (e NegativeNumberError) Error() string {
	return fmt.Sprintf("Число %f отрицательное", e.a)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, NegativeNumberError{a: x}
	}
	return math.Sqrt(x), nil
}

func main() {
	a := -4.0
	result, err := Sqrt(a)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Корень:", result)
}
