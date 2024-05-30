package main

import "fmt"

type Book struct {
	name   string
	author string
}

type Shelf struct {
	id    int
	books []Book
}

func sortByName(shelf []Book) {
	for i := 0; i < len(shelf)-1; i++ {
		smallestPosition := i
		for j := i + 1; j < len(shelf); j++ {
			if shelf[j].name < shelf[smallestPosition].name {
				smallestPosition = j
			}
		}
		shelf[i], shelf[smallestPosition] = shelf[smallestPosition], shelf[i]
	}
}

func sortByAuthor(shelf []Book) {
	for i := 0; i < len(shelf)-1; i++ {
		smallestPosition := i
		for j := i + 1; j < len(shelf); j++ {
			if shelf[j].author < shelf[smallestPosition].author {
				smallestPosition = j
			}
		}
		shelf[i], shelf[smallestPosition] = shelf[smallestPosition], shelf[i]
	}
}

func sortAllBooks(library []Shelf) {
	allBooks := make([]Book, 0, len(library))
	for i := 0; i < len(library); i++ {
		for j := 0; j < len(library[i].books); j++ {
			allBooks = append(allBooks, library[i].books[j])
		}
	}
	sortByName(allBooks)
	q := 0
	for i := 0; i < len(library); i++ {
		for j := 0; j < len(library[i].books); j++ {
			library[i].books[j] = allBooks[q]
			q += 1
		}
		fmt.Printf("%v-я полка содержит книги с %c до %c %v\n", i+1, rune(library[i].books[0].name[0]), rune(library[i].books[len(library[i].books)-1].name[0]), library[i])
	}

}

func main() {
	library := [3]Shelf{{1, []Book{{"Book Name", "NameAuthor"}, {"Shine", "Abc"}, {"Olie", "Xyz"}, {"Beta", "Def"}, {"Alpha", "Tbjorn"}}},
		{2, []Book{{"Wonder", "Land"}, {"Authorname", "Author4"}}}, {3, []Book{{"Just", "Know"}, {"Why", "Zety"}}}}
	sortByName(library[0].books)
	fmt.Println(library[0].books)
	sortByAuthor(library[1].books)
	fmt.Println(library[1].books)
	sortAllBooks(library[:])
}
