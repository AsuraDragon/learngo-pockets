package main

import (
	"sort"
)

type bookCollection map[Book]struct{}

type bookRecommendations map[Book]bookCollection

func newCollection() bookCollection {
	return make(bookCollection)
}

func listOtherBooksOnShelves(bookIndexToRemove int, myBooks []Book) []Book {
	otherBooksOnShelves := make([]Book, bookIndexToRemove, len(myBooks)-1)
	copy(otherBooksOnShelves, myBooks[:bookIndexToRemove])
	otherBooksOnShelves = append(otherBooksOnShelves, myBooks[bookIndexToRemove+1:]...)
	return otherBooksOnShelves
}

func registerBookRecommendations(recommendations bookRecommendations, currentBook Book, otherBooksOnShelves []Book) {
	for _, book := range otherBooksOnShelves {
		collection, ok := recommendations[currentBook]
		if !ok {
			collection = newCollection()
			recommendations[currentBook] = collection
		}
		collection[book] = struct{}{}
	}
}

func bookCollectionToListOfBooks(bc bookCollection) []Book {
	bookList := make([]Book, 0, len(bc))
	for book := range bc {
		bookList = append(bookList, book)
	}
	sort.Slice(bookList, func(i, j int) bool {
		if bookList[i].Author != bookList[j].Author {
			return bookList[i].Author < bookList[j].Author
		}
		return bookList[i].Title < bookList[j].Title
	})
	return bookList
}

func recommendBooks(recommendations bookRecommendations, myBooks []Book) []Book {
	collectionOfBooks := make(bookCollection)

	myShelf := make(map[Book]bool)
	for _, myBook := range myBooks {
		myShelf[myBook] = true
	}

	for _, myBook := range myBooks {
		for recomendations := range recommendations[myBook] {
			if myShelf[recomendations] {
				continue
			}
			collectionOfBooks[recomendations] = struct{}{}
		}
	}
	recommendationsForABook := bookCollectionToListOfBooks(collectionOfBooks)
	return recommendationsForABook
}

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookRecommendations)
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookRecommendations(sb, book, otherBooksOnShelves)
		}
	}

	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworms := range bookworms {
		recommendations[i] = Bookworm{
			Name:  bookworms.Name,
			Books: recommendBooks(sb, bookworms.Books),
		}
	}
	return recommendations
}
