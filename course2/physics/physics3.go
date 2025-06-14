package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GenDisplaceFn(acceleration, velocity, initialDisplacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		firstComponent := 0.5 * acceleration * time * time
		secondComponent := velocity * time
		thirdComponent := initialDisplacement
		return firstComponent + secondComponent + thirdComponent
	}
}
func main() {
	var acceleration float64
	var initialDisplacement float64
	var velocity float64
	var time float64
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Value for Acceleration")
	scanner.Scan()
	acceleration, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Println("Enter Value for Initial Displacement")
	scanner.Scan()
	initialDisplacement, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Println("Enter Value for velocity")
	scanner.Scan()
	velocity, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Println("GENERATING FUNCTION............")
	timeDependentFunc := GenDisplaceFn(acceleration, velocity, initialDisplacement)
	fmt.Println("Enter TIME Value:")
	fmt.Scan(&time)
	fmt.Println(timeDependentFunc(time))

}