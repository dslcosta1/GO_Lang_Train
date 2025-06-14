package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortArray(array []int, wg *sync.WaitGroup) {
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	wg.Done()
}

func main() {
	var input string
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a series of integers (space-separated).")
	input, _ = scanner.ReadString('\n')
	input = strings.TrimSpace(input)
	stringsArray := strings.Split(input, " ")

	size := len(stringsArray)
	maxElement, _ := strconv.Atoi(stringsArray[0])
	var array = make([]int, size)
	for i := 0; i < size; i++ {
		array[i], _ = strconv.Atoi(stringsArray[i])
		if array[i] > maxElement {
			maxElement = array[i]
		}
	}

	fmt.Println("Array before: ", array)

	subSizes := []int{size / 4, size / 4, size / 4, size / 4}
	for i := 0; i <= (size%4 - 1); i++ {
		subSizes[i]++
	}

	var wg sync.WaitGroup
	wg.Add(4)
	var subArrays [][]int
	currentIndex := 0
	fmt.Println()
	for i := 0; i < 4; i++ {
		subArray := array[currentIndex:(currentIndex + subSizes[i])]
		subArrays = append(subArrays, subArray)
		fmt.Println("Subarray", i+1, " before: ", subArrays[i])
		go sortArray(subArrays[i], &wg)
		currentIndex += subSizes[i]
	}
	wg.Wait()

	fmt.Println()
	for i := 0; i < 4; i++ {
		fmt.Println("Subarray", i+1, " after: ", subArrays[i])
	}

	var resArray = make([]int, size)
	subIndices := []int{0, 0, 0, 0}
	currentIndex = 0
	a := -1
	b := -1
	c := -1
	d := -1
	for currentIndex < size {
		if subIndices[0] < subSizes[0] {
			a = subArrays[0][subIndices[0]]
		} else {
			a = maxElement + 1
		}
		if subIndices[1] < subSizes[1] {
			b = subArrays[1][subIndices[1]]
		} else {
			b = maxElement + 1
		}
		if subIndices[2] < subSizes[2] {
			c = subArrays[2][subIndices[2]]
		} else {
			c = maxElement + 1
		}
		if subIndices[3] < subSizes[3] {
			d = subArrays[3][subIndices[3]]
		} else {
			d = maxElement + 1
		}

		if a <= b && a <= c && a <= d {
			resArray[currentIndex] = a
			subIndices[0]++
		} else if b <= a && b <= c && b <= d {
			resArray[currentIndex] = b
			subIndices[1]++
		} else if c <= a && c <= b && c <= d {
			resArray[currentIndex] = c
			subIndices[2]++
		} else {
			resArray[currentIndex] = d
			subIndices[3]++
		}

		currentIndex++
	}

	fmt.Println()
	fmt.Println("Array after: ", resArray)
}