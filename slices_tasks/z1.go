package main

import "fmt"

func slice(n int) {
	srez := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		srez = append(srez, i)
	}
	fmt.Println("Изначально:", srez)
	for i := n + 1; i <= 2*n; i++ {
		srez = append(srez, i)
	}
	fmt.Println("После append:", srez)
}
func main() {
	slice(5)
}
