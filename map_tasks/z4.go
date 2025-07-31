package main

import "fmt"

func isAnagram(firstWord, secondWord string) bool {
	if len(firstWord) != len(secondWord) {
		return false
	}

	letterCounts := make(map[rune]int)

	for _, letter := range firstWord {
		letterCounts[letter]++
	}

	for _, letter := range secondWord {
		letterCounts[letter]--
		if letterCounts[letter] < 0 {
			return false
		}
	}

	return true
}

func main() {
	var word1, word2 string

	fmt.Print("Введите первое слово: ")
	fmt.Scanln(&word1)

	fmt.Print("Введите второе слово: ")
	fmt.Scanln(&word2)

	if isAnagram(word1, word2) {
		fmt.Println("Результат: true (это анаграммы)")
	} else {
		fmt.Println("Результат: false (это НЕ анаграммы)")
	}
}
