package main

import (
	"encoding/json"
	"os"
	"sort"
)

// A Bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookWorms reads the file and returns the list of bookworms,
// and their beloved books, found therein
func loadBookWorms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var bookworms []Bookworm
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}
	return bookworms, nil
}

// sortBooks sorts the books by Author and then Title
func sortBooks(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool { //#1
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author //#2
		}
		return books[i].Title < books[j].Title
	})
	return books
}

func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms)
	var commonBooks []Book

	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	return sortBooks(commonBooks)
}

// booksCount registers all the books and their occurrences
// from the bookworms shelves
func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint, 0)      //#1
	for _, bookworm := range bookworms { //#2
		for _, book := range bookworm.Books {
			count[book]++ //#3
		}
	}
	return count
}
