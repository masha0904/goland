package main

import "fmt"

func main() {
	readers := map[string]map[string][]string{}

	readers["Гракова М. Д."] = map[string][]string{
		"Книги": {"Война и мир", "Преступление и наказание"},
	}

	readers["Тарун Г. Г."] = map[string][]string{
		"Книги":   {"Богатый папа, бедный папа"},
		"Журналы": {"Незнайка"},
	}

	fmt.Println("Список читателей с изданиями:")
	for reader, publications := range readers {
		fmt.Printf("- %s:\n", reader)

		count := 0

		for pubType, items := range publications {
			fmt.Printf("  %s: %v\n", pubType, items)
			count += len(items)
		}

		fmt.Printf("Количество изданий на руках: %d\n", count)
	}

	fmt.Printf("\nКоличество читателей: %d\n", len(readers))
}
