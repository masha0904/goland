package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 10
	m["cherry"] = 15
	fmt.Println(m)
}
