package main

import (
	"errors"
	"fmt"
)

func main() {
	err := first()
	if err != nil {
		fmt.Println("Ошибка на верхнем уровне:", err)
	}
}

func first() error {
	fmt.Println("Выполняем first()")
	if err := second(); err != nil {
		return fmt.Errorf("first: %w", err)
	}
	return nil
}

func second() error {
	fmt.Println("Выполняем second()")
	if err := third(); err != nil {
		return fmt.Errorf("second: %w", err)
	}
	return nil
}

func third() error {
	fmt.Println("Выполняем third()")
	return errors.New("что-то пошло не так в third()")
}
