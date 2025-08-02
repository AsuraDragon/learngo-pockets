package main

import (
	"fmt"
	"os"
)

// displayBooks prints out the titles and authors of a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Printf("%v", book)
	}
}

func main() {
	bookworms, err := loadBookWorms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	commonBooks := findCommonBooks(bookworms)
	fmt.Println("Here are the books that are common:")
	displayBooks(commonBooks)
}
