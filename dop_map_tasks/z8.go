package main

import "fmt"

func main() {
	text := "hello"
	m1 := map[rune]int{}
	for _, c := range text {
		m1[(c)]++
	}
	for k, v := range m1 {
		fmt.Printf("%c: %d\n", k, v)
	}
}
