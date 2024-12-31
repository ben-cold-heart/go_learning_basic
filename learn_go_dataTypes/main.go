package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ประกาศตัวแปรแบบ
	var name string = "Ben"
	var age int = 50.0

	var name2 = "Ben"
	var age2 = 50

	name3 := "Ben"
	age3 := 50.0

	fmt.Println(name, name2, name3)
	fmt.Println(age, age2, age3)

	// เช็คชนิดของตัวแปร
	fmt.Println("Type of name:", reflect.TypeOf(name))
	fmt.Println("Type of name2:", reflect.TypeOf(name2))
	fmt.Println("Type of name3:", reflect.TypeOf(name3))

	fmt.Println("Type of age:", reflect.TypeOf(age))
	fmt.Println("Type of age2:", reflect.TypeOf(age2))
	fmt.Println("Type of age3:", reflect.TypeOf(age3))
}
