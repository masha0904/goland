package main

import "fmt"

func main() {
	colors := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}
	for k, v := range colors {
		fmt.Printf("%s: %s\n", k, v)
	}
}
