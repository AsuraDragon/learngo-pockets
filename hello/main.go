package main

import "fmt"

func greet() string {
	return "Hello World"
}

func main() {
	greeting := greet()
	fmt.Println(greeting)
}
