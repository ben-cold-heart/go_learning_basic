package main

import (
	"fmt"
)

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Friday":
		fmt.Println("It's Friday")
	default:
		fmt.Println("It's Someday")
	}

	switch x := 15; {
	case x < 5:
		fmt.Println("น้อยกว่า 5")
	case x < 10:
		fmt.Println("น้อยกว่า 10")
		// เพื่อให้ code ถูก run ต่อไป
		fallthrough
	default:
		fmt.Println("default case")
	}

}
