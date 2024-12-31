package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Waiting...")

	// array fixed size
	var num [5]int

	// Add Value to Array
	num[0] = 1
	num[1] = 2

	fmt.Println(num[0])

	// if index doesnt have value, return 0
	fmt.Println(num[3])

	// slice
	var num2 []int

	// add value to slice
	num2 = append(num2, 0, 1, 2, 3)
	fmt.Println(num2)

	// edit value in slice
	num2[3] = 5
	fmt.Println(num2)

	// สร้าง Slice จาก Array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]  // สร้าง slice จาก index 1 ถึง 3
	fmt.Println(slice) // [2 3 4]

	// check size and capacity of slice and array
	fmt.Println(len(num2))
	fmt.Println(cap(num2))

	// Map
	var person map[string]string = make(map[string]string)

	// Add value to Map
	person["name"] = "John"
	person["age"] = "30"

	fmt.Println(person["name"], person["age"])

	// Delete KV from Map
	delete(person, "age")
	fmt.Println(person)

	// Create Map with KV
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors["red"])

}
