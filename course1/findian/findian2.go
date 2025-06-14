package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		startsWith = "i"
		endsWith   = "n"
		contains   = "a"

		input string
	)

	fmt.Println("Please enter a string:")

	// Create a reader for standard input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	// Trim spaces and convert to lowercase
	input = strings.TrimSpace(strings.ToLower(input))
	fmt.Println("Scanned:", input)

	if strings.HasPrefix(input, startsWith) &&
		strings.HasSuffix(input, endsWith) &&
		strings.Contains(input, contains) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}