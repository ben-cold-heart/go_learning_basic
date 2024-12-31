package main

import (
	"fmt"
)

func main() {
	fmt.Println("------- General loop -------")

	// general syntax
	for i := 0; i < 5; i++ {
		fmt.Println("round", i)
	}

	fmt.Println("------- condition loop -------")

	// use condition loop
	x := 1
	for x < 5 {
		fmt.Println("round", x+4)
		x++
	}

	fmt.Println("------- use break for stop loop -------")

	// ใช้ break เพื่อ stop loop
	for {
		fmt.Println("round 9")
		break
	}

}
