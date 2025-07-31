package main

import "fmt"

func add(slicePtr *[]string) {
	*slicePtr = append(*slicePtr, "modified")
}
func main() {
	slice := []string{"a", "b"}
	fmt.Println("Вход:", slice)

	add(&slice)

	fmt.Println("Выход:", slice)
}
