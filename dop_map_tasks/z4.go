package main

import "fmt"

func main() {
	fruits := map[string]int{"apple": 5, "banana": 10, "cherry": 15}
	delete(fruits, "banana")
	fmt.Println(fruits)
}
