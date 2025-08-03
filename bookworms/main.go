package main

import (
	"flag"
	"fmt"
	"os"
)

// displayBooks prints out the titles and authors of a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Printf("%+v\n", book)
	}
}

func displayRecommendations(recommendations []Bookworm) {
	for _, bookworm := range recommendations {
		fmt.Printf("\nHere are the recommendations for %s:\n", bookworm.Name)
		displayBooks(bookworm.Books)
		fmt.Println()
	}
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "filePath", "testdata/bookworms.json",
		`Introduce the file path of an existing JSON file that contains the booksworms structure information
EXAMPLE OF FILE

[
    {
        "name": "Karen Georgina Medina Gonzales",
        "books": [
            {
                "author": "Margaret Atwood",
                "title": "The Handmaid's Tale"
            },
            {
                "author": "Sylvia Plath",
                "title": "The Bell Jar"
            }
        ]
    }
]
	`)
	flag.Parse()
	bookworms, err := loadBookWorms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	commonBooks := findCommonBooks(bookworms)
	fmt.Println("Here are the books that are common:")
	displayBooks(commonBooks)

	recommendations := recommendOtherBooks(bookworms)
	displayRecommendations(recommendations)
}
