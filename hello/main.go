package main

import "fmt"

type language string //#1

func greet(l language) string {
	switch l { //# 2
	case "en":
		return "Hello World"
	case "fr":
		return "Bonjour le monde"
	default:
		return ":=)"
	}
}

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}
