package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sliceInts := make([]int, 0, 3) 
	var userInput string

	fmt.Println("Enter integers to store in a sorted slice. Enter 'X' to stop.")

	for {
		fmt.Print("Enter an integer (or 'X' to stop): ")
		fmt.Scan(&userInput)

		if userInput == "X" || userInput == "x" {
			break
		}

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Invalid input! Please enter an integer or 'X' to stop.")
			continue
		}

		sliceInts = append(sliceInts, num)

		sort.Ints(sliceInts)

		fmt.Println("Sorted slice:", sliceInts)
	}

	fmt.Println("Final sorted slice:", sliceInts)
}