package main

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)

func sortSlice(slc []int, c chan []int) {
	sort.Ints(slc)
	fmt.Println(slc)
	c <- slc
}

func merge(slc1 []int, slc2 []int, c chan []int) {
	var slc3 = make([]int, 0)
	i := 0
	j := 0

	continueLoop := i < len(slc1) && j < len(slc2)
	
	for continueLoop {
		if slc1[i] < slc2[j] {
			slc3 = append(slc3, slc1[i])
			i++
		} else {
			slc3 = append(slc3, slc2[j])
			j++
		}

		continueLoop = i < len(slc1) && j < len(slc2)
	}

	if i != len(slc1) {
		slc3 = append(slc3, slc1[i:]...)
	}

	if j != len(slc2) {
		slc3 = append(slc3, slc2[j:]...)
	}

	c <- slc3
}

func readUserInput() []int {
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

	return slc
}

func main () {
	c := make(chan []int, 4)
	
	
	userInput := readUserInput()
	sizeArray := len(userInput)
	mid := sizeArray/2
	quarter := mid/2
	thirdQuarter := mid + quarter


	go sortSlice(userInput[0:quarter], c)
	go sortSlice(userInput[quarter:mid], c)
	go sortSlice(userInput[mid:thirdQuarter], c)
	go sortSlice(userInput[thirdQuarter:], c)

	slc1 := <- c 
	slc2 := <- c 
	slc3 := <- c 
	slc4 := <- c 

	go merge(slc1, slc2, c)
	go merge(slc3, slc4, c)

	slc5 := <- c
	slc6 := <- c

	go merge(slc5, slc6, c)
	slcFinal := <- c 

	fmt.Printf("The final sorted string is: ")
	fmt.Println(slcFinal)
}