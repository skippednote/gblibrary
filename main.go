package main

import (
	"log"
)

func init() {
	setup()
}

func main() {
	books, err := getBooksList()
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan book, len(books))
	for _, book := range books {
		getBook(book, c)
	}

	for i := 0; i < len(books); i++ {
		save(<-c)
	}
}
