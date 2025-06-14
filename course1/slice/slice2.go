package main

import (
	"fmt"
	"strings"
	"slices"
	"strconv"
)

func main() {

/*
	Write a program which 
		1 prompts the user to enter integers and stores the integers in a sorted slice. 
		2 The program should be written as a loop. 
		3 Before entering the loop, the program should create an empty integer slice of size (length) 3. 
		4 During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. 
		5 The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. 
		6 The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/
	var userInp string //get inp as string, convert to int if it is not "X" or "x"
	sliceInt := make([]int, 3)
	var counter = 0


	for !strings.EqualFold(userInp, "x") {

		fmt.Printf("Enter an integer (or enter 'X' to stop):") //
		fmt.Scan(&userInp)

		if (!strings.EqualFold(userInp, "x")) {
			if userInt, err := strconv.Atoi(userInp); err == nil {

				//TODO add logic in order to replace initial 3 (zero values)
				//then append
				if counter < 3 {
					idxToReplace := slices.Index(sliceInt, 0)
					sliceInt = slices.Replace(sliceInt, idxToReplace, idxToReplace + 1, userInt)
					counter += 1

				 	fmt.Printf("TODO, not sure how to handle this case")
				} else {
					sliceInt = append(sliceInt, userInt)
				}

		 		slices.Sort(sliceInt)
			 	fmt.Printf("The sorted slice is: %v\n", sliceInt)
			}
		} else {
			fmt.Printf("Bye!\n")
			break
		}
	}
}