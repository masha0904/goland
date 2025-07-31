package main

import "fmt"

func deletes(newslice *[]int, index int) {
	*newslice = append((*newslice)[:index], (*newslice)[index+1:]...)
}
func main() {
	slice := []int{10, 20, 30, 40}
	index := 1
	deletes(&slice, index)

	fmt.Println(slice)
}
