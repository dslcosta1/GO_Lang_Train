package main

import (
	"fmt"
)

func Swap(numbers []int, num int) {
	numbers[num], numbers[num+1] = numbers[num+1], numbers[num]
}

func BubbleSort(numbers []int) {

	for i := 0; i < len(numbers)-1; i++ {
		alreadySwapped := false
		for comp := 0; comp < len(numbers)-i-1; comp++ {
			if numbers[comp] > numbers[comp+1] {
				Swap(numbers, comp)
				alreadySwapped = true
			}

		}

		if !alreadySwapped {
			return
		}

	}

}

func main() {

	numbers := make([]int, 0, 10)
	fmt.Printf("Enter a total of 10 integers in sequence \n")
	count := 1
	for {
		var number int
		fmt.Printf("number %d: ", count)
		_, err := fmt.Scan(&number)

		if err != nil {
			fmt.Printf("Error scanning: %s", err)
		}

		numbers = append(numbers, number)
		count++

		if len(numbers) >= 10 {
			break
		}

	}

	BubbleSort(numbers)
	fmt.Printf("%d", numbers)

}