package main

import (
	"fmt"
	"strconv"
)

/*
Structures can have properties and methods.
Two type of methods
1. pointer receiver
2. value receiver
*/

// Define struct book
type Book struct {
	bookName string
	author   string
	isbn     int
}

// func (of type struct) fnName() returnType // Value receiver method
func (b Book) getBookDetails() string {
	return "Book : " + b.bookName + " Author : " + b.author + " isbn : " + strconv.Itoa(b.isbn)
}

// Pointer receiver method
func (b *Book) updateISBNtoZero() {
	b.isbn = 00000000
}

func (b *Book) setISBN(isb int) {
	b.isbn = isb
}

func main() {

	// Init Book structure
	b1 := Book{bookName: "Totochan", author: "Testuko K", isbn: 123123}
	b2 := Book{"Adayalangal", "sethu", 10111211}

	fmt.Println(b1)
	fmt.Println(b2)

	// Modify struct values
	fmt.Println(b1.bookName)
	b2.isbn = 4567891011
	fmt.Println(b2)

	// using methods
	fmt.Println(b1.getBookDetails())

	b1.updateISBNtoZero()
	fmt.Println(b1.getBookDetails())

	b1.setISBN(123456789)
	fmt.Println(b1.getBookDetails())

	//Empty 
	b3 := Book{}
	b3.bookName = "NewBook"
	fmt.Println(b3)
}
