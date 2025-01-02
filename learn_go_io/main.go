package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create File
	file, err := os.Create("./example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello IO\nThis's from Ben")
	fmt.Println("complete create and write file")

	// Open and Write file
	file, err = os.OpenFile("./example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	file.WriteString("\nThis's from Ben But second times!")
	fmt.Println("complete open and write file")

	// only Read file
	data, err := os.ReadFile("./example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Start Reading file")
	fmt.Println(string(data))

	// Read line by line
	file, err = os.Open("./example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("Start Reding line by line")
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Err:", err)
	}

	// Delete file
	errDel := os.Remove("./example.txt")

	if errDel != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File is deleted")
}
