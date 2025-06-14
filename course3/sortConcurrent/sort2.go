package main

import (
	"fmt"
	"sort"
	"sync"
)

// merge combines two sorted slices into one sorted slice
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// calculateIndices calculates the start indices for dividing the array into n parts
func calculateIndices(size, n int) []int {
	q, r := size/n, size%n
	indices := make([]int, n+1)
	for i := 0; i <= n; i++ {
		indices[i] = i*q + min(i, r)
	}
	return indices
}

// divideArray splits the array into n contiguous subarrays
func divideArray(arr []int, n int) [][]int {
	size := len(arr)
	indices := calculateIndices(size, n)

	parts := make([][]int, n)
	for i := 0; i < n; i++ {
		parts[i] = arr[indices[i]:indices[i+1]]
	}
	return parts
}

func main() {
	var n, numParts int
	numParts = 4

	fmt.Print("Enter the number of integers: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter the integers:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Divide array into n subarrays
	parts := divideArray(arr, numParts)

	var wg sync.WaitGroup
	wg.Add(numParts)

	// Sort each part concurrently
	for i := 0; i < numParts; i++ {
		go func(i int) {
			defer wg.Done()
			// fmt.Printf("Goroutine %d sorting: %v\n", i+1, parts[i])
			sort.Ints(parts[i])
			fmt.Printf("Goroutine %d sorted: %v\n", i+1, parts[i])
		}(i)
	}

	wg.Wait()

	// Merge sorted subarrays
	merged := parts[0]
	for i := 1; i < numParts; i++ {
		merged = merge(merged, parts[i])
	}

	fmt.Println("Sorted array:", merged)
}