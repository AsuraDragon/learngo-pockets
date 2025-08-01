package main

import (
	"testing"
)

var (
	handmaidsTale = Book{
		Author: "Margaret Atwood", Title: "The Handmaid's Tale",
	}
	oryxAndCrake = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar   = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre     = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

// equalBooks is a helper to test the equality of two lists of Books.
func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()
	if len(books) != len(target) {
		return false
	}
	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}
	return true
}

// equalBookworms is a helper to test the equality of two lists of Bookworms.
func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()
	if len(bookworms) != len(target) {
		return false
	}
	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}
		if !equalBooks(t,
			bookworms[i].Books, target[i].Books) {
			return false
		}
	}
	return true
}

func TestLoadBookworms_Success(t *testing.T) {
	tests := map[string]struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookWorms(testCase.bookwormsFile)
			if err != nil && !testCase.wantErr {
				t.Fatalf("expected an error %s, got none", err.Error())
			}
			if err == nil && testCase.wantErr {
				t.Fatalf("expected no error, got one")
			}
			if !equalBookworms(t, got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v",
					got, testCase.want)
			}
		})
	}
}

// equalBooksCount is a helper to test the equality of two maps of
// books count.
func equalBooksCount(t *testing.T, got map[Book]uint, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) { //#1
		return false
	}
	for book, targetCount := range want { //#2
		count, ok := got[book]           //#3
		if !ok || count != targetCount { //#4
			return false //#5
		}
	}

	return true //#6
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{
					handmaidsTale, theBellJar, // #1
				}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{}, //#2
		},
		"bookworm without books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: map[Book]uint{},
		}, "bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, handmaidsTale}},
				{Name: "Peggy", Books: []Book{handmaidsTale}},
			},
			want: map[Book]uint{handmaidsTale: 3},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(t, got, tc.want) { //#3
				t.Fatalf("Got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}
