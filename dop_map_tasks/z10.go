package main

import "fmt"

func main() {
	m := map[int]string{1: "one", 2: "two"}
	m2 := map[string]int{}
	for k, v := range m {
		m2[v] = k
	}
	fmt.Println(m2)
}
