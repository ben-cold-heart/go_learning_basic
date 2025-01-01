package main

import (
	"fmt"
)

// Struct 1 {Using for Nested Struct}
type Person struct {
	Name    string
	Age     int
	Address Address
}

// Nested Struct
type Address struct {
	Street  string
	City    string
	ZipCode string
}

// Struct 2 {Using for Struct Method}
type Book struct {
	Id   int
	Name string
}

// Anonymous fields
type Employee struct {
	string
	int
}

// Struct Method
func (b Book) statusBook() {
	fmt.Println(b.Name, "is available for rent!!")
}

func main() {
	// Person
	p := Person{Name: "Bob", Age: 20}
	fmt.Println("start :", p.Name, p.Age)

	// how to edit and add values to struct
	p.Name = "Ben"
	p.Age = 34
	fmt.Println("after edit :", p.Name, p.Age)
	fmt.Println("after edit :", p)

	// Nested Struct
	p.Address = Address{
		City:    "Bangkok",
		ZipCode: "10110",
	}
	fmt.Println("after add nested struct :", p)

	// Adding More KV on Nested Struct
	p.Address.Street = "With-U"
	fmt.Println("after edit nested struct :", p)

	// Book
	b := Book{}
	b.Id = 1
	b.Name = "big debt crisis"

	// struct Method
	b.statusBook()

	// Anonymous fields
	e := Employee{p.Name, p.Age}
	fmt.Println("employee", e.int)
	fmt.Println("employee", e.string)

}
