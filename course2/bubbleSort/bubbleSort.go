package main

import (
	"fmt"
	"strings"
	"strconv"
)

func Swap(slc []int) {
	aux := slc[0]
	slc[0] = slc[1]
	slc[1] = aux

}


func BubbleSort(slc []int) {
	if (len(slc) == 1) {
		return
	}

	for index, value := range slc {
		if index + 1 == len(slc) {
			break
		}

		if value > slc[index + 1] {
			Swap(slc[index:index + 2])
		}
	}

	BubbleSort(slc[0:len(slc)-1])
}

func main() {
	slc := []int{}
	var inputUser string

	for true {
		fmt.Printf("Input a integer number: (When you are done input `X`): ")
		fmt.Scan(&inputUser)
		inputUser = strings.ToLower(inputUser)
		if inputUser == "x" {
			if (len(slc) < 10) {
				fmt.Println("Please input at least 10 numbers")
			} else {
				break
			}
		}

		intInput, err := strconv.Atoi(inputUser)

		if err != nil {
			fmt.Println("Please input a valid int number or the value `X`")
		} else {
			slc = append(slc, intInput)
		}
	}

	BubbleSort(slc)
	fmt.Println(slc)
}