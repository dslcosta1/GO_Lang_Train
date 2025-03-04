package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput string

	fmt.Printf("Enter the string: ")
	fmt.Scan(&userInput)

	var lowerUserInput = strings.ToLower(userInput)

	if strings.HasPrefix(lowerUserInput, "i") && 
	strings.Contains(lowerUserInput, "a") && 
	strings.HasSuffix(lowerUserInput, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}