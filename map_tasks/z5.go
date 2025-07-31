package main

import "fmt"

func addGrade(db map[string]map[string]int, student string, subject string, grade int) {
	if db[student] == nil {
		db[student] = make(map[string]int)
	}
	db[student][subject] = grade
}

func main() {
	
	fmt.Println("Вход:")
	db := map[string]map[string]int{}
	fmt.Printf("db := map[string]map[string]int{}\n")
	fmt.Printf("addGrade(db, \"Alice\", \"Math\", 90)\n")

	addGrade(db, "Alice", "Math", 90)

	fmt.Println("\nВыход:")
	fmt.Printf("map[Alice:map[Math:90]]")
}
