package main

import "fmt"

type Book struct {
	title     string
	author    string
	pages     int
	pagesRead int
	sequel    *Book
}

func (b *Book) read(pages int) {
	b.pagesRead += pages
}

func (b *Book) pagesLeft() int {
	return b.pages - b.pagesRead
}

func main() {
	// initialize a struct with value

	// initialize a struct without values
	harryPoterTheChamberOfSecrets := new(Book)
	harryPoterTheChamberOfSecrets.title = "Harry Potter and the Chamber of Secrets"
	harryPoterTheChamberOfSecrets.author = "J.K. Rowling"
	harryPoterTheChamberOfSecrets.pages = 341
	harryPoterTheChamberOfSecrets.sequel = &Book{
		title:     "Harry Potter and the phylosopher's stone",
		pages:     435,
		pagesRead: 0,
		sequel:    nil,
	}

	// read 100 pages
	harryPoterTheChamberOfSecrets.read(100)

	// print the sequel title
	fmt.Println(harryPoterTheChamberOfSecrets.sequel.title)

	// print the pages left
	fmt.Println(harryPoterTheChamberOfSecrets.pagesLeft())
}
