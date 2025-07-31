package main

import "fmt"

func main() {
	scores := map[string]int{"math": 90, "physics": 85, "chemistry": 78}
	a := len(scores)
	fmt.Println("Количество элементов:", a)
}
