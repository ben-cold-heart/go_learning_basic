package main

import (
	"fmt"
)

// function แบบ ไม่ return ค่า
func HelloWorld() {
	fmt.Println("Hello World")
}

// function ที่รับ parameter
func Greeting(age int, name string) {
	fmt.Println("Hello", name, "and your age is", age)
}

// function ที่มีการ return ค่า
func add(a int, b int) int {
	return a + b
}

// function ที่รับหลายค่าและ return หลายค่า
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

// Name return Values
func calculator(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// Variadic Functions
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Higher Order Function
func executeFunction(f func(int, string), c func(), name string, age int) {
	fmt.Println("---------")
	fmt.Println("this is Execute from Higher Order")
	c()
	f(age, name)
}

func executeCallback(callback func(string), text string) {
	callback(text)
}

func main() {

	HelloWorld()
	Greeting(20, "Ben")

	result := add(5, 202) // Changed from 202.0 to 202
	fmt.Println("Result is", result)

	quotient, remainder := divide(11, 2)
	fmt.Println("Result2 is", quotient, "and", remainder)

	sum2, diff := calculator(5, 2)
	fmt.Println("Result3 is", sum2, "and", diff)

	fmt.Println("Result4 is", sum(1, 2, 3, 4, 5))

	func(message string) {
		fmt.Println(message)
	}("Hello Anonymous Function")

	executeFunction(Greeting, HelloWorld, "Ben", 500)

	// ส่งฟังก์ชันนิรนามเป็น callback
	executeCallback(func(msg string) {
		fmt.Println("Anonymous Callback:", msg)
	}, "This is a test!")

}
