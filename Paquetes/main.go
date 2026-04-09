package main

import (
	"fmt"
	"os"
)

func main() {
	enVar := os.Getenv("HOME")

	if enVar == "" {
		fmt.Println("HOME enviroment variable is not set")
	} else {
		fmt.Printf("HOME enviroment variable: %s\n", enVar)
	}

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}
	defer file.Close()
}
