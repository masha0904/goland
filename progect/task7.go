package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
func isitpolindrom(s string) string {
	if reverseString(s) == s {
		return ("Строка является полиндромом")
	}
	return ("Строка не является полиндромом")
}
func main() {
	s := "Доброе утро!"
	fmt.Println(isitpolindrom(s))
}
