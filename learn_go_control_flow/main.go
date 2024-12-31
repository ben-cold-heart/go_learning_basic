package main

import (
	"fmt"
)

func main() {
	x := 20

	if x > 20 {
		fmt.Println("x is more than 20")
	} else if x == 20 {
		fmt.Println("x is equal to 20")
	} else {
		fmt.Println("x is less than 20")
	}
}
