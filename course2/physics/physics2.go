package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var acceleration float64
	var initialVelocity float64
	var initialDisplacement float64
	var time float64

	scanner := bufio.NewScanner(os.Stdin)
	waitForInput(scanner, &acceleration, "Enter acceleration: ")
	waitForInput(scanner, &initialVelocity, "Enter initial velocity: ")
	waitForInput(scanner, &initialDisplacement, "Enter initial displacement: ")
	waitForInput(scanner, &time, "Enter time: ")

	fmt.Print("\n") // Empty line just to separate input and output

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Printf("Displacement after %f seconds is %f", time, fn(time))
}

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		// formula: s = 1/2*a*t^2 + v0*t + s0
		return 0.5*acceleration*math.Pow(time, 2) + initialVelocity*time + initialDisplacement
	}
}

func waitForInput(scanner *bufio.Scanner, variable *float64, prompt string) {
	fmt.Print(prompt)
	scanner.Scan()
	panicIf(scanner.Err())

	var err error
	*variable, err = strconv.ParseFloat(scanner.Text(), 64)
	panicIf(err)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}