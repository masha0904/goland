package main

import "fmt"

func main() {
	ages := map[string]int{"Alice": 25, "Bob": 30}
	_, ok := ages["Bob"]
	if ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("No exists")
	}
}
