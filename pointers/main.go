package main

import "fmt"

func main() {

	a := 10
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T \n", b)
	// Read value
	fmt.Println(" value of a using pointer", *b)
	fmt.Println(" value of a using pointer", *&a)

	// Change value
	*b = 11
	c := *b + 1
	fmt.Println("Value after changes ", a, c)
}
