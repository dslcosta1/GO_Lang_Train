package main

import (
	"fmt"
	"sort"
)

func main() {
	var slc = make([]int, 0, 3)
	var userInput int
	var end = true

	for end {
		fmt.Printf("Input a integer number. (To exit press `X`): ")
		_, err := fmt.Scanf("%d", &userInput)

		if err != nil {
			end = false
		} else {
			slc = append(slc, userInput)

			sort.Slice(slc, func(i, j int) bool {
				return slc[i] < slc[j]
			})

			fmt.Println(slc)
		}
	}
}